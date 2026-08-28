package persistence

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"sync"

	"go.etcd.io/bbolt"
	"jellyfield/model"
)

const (
	EntityJellyfish     = "Jellyfish"
	EntitySceneState    = "SceneState"
	EntityGestureRecord = "GestureRecord"
	EntityControlEvent  = "ControlEvent"
	bucketScenes        = "scenes"
	bucketGestures      = "gestures"
	bucketControls      = "controls"
)

type Store struct {
	db   *bbolt.DB
	path string
	mu   sync.RWMutex
}

func Open(path string) (*Store, error) {
	if path == "" {
		return nil, fmt.Errorf("database path is required")
	}
	db, err := bbolt.Open(filepath.Clean(path), 0600, nil)
	if err != nil {
		return nil, err
	}
	s := &Store{db: db, path: path}
	if err := s.initialize(); err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}
func (s *Store) initialize() error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		for _, name := range []string{bucketScenes, bucketGestures, bucketControls} {
			if _, err := tx.CreateBucketIfNotExists([]byte(name)); err != nil {
				return err
			}
		}
		return nil
	})
}
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		return nil
	}
	err := s.db.Close()
	s.db = nil
	return err
}
func (s *Store) Path() string { return s.path }

func marshal(value any) ([]byte, error) { return json.Marshal(value) }
func (s *Store) put(bucket, key string, value any) error {
	data, err := marshal(value)
	if err != nil {
		return err
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return fmt.Errorf("store is closed")
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(bucket)).Put([]byte(key), data) })
}
func (s *Store) get(bucket, key string, target any) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return fmt.Errorf("store is closed")
	}
	return s.db.View(func(tx *bbolt.Tx) error {
		data := tx.Bucket([]byte(bucket)).Get([]byte(key))
		if data == nil {
			return bbolt.ErrBucketNotFound
		}
		return json.Unmarshal(append([]byte(nil), data...), target)
	})
}

func (s *Store) SaveScene(scene model.SceneState) error {
	if err := scene.Validate(); err != nil {
		return err
	}
	return s.put(bucketScenes, scene.ID, scene)
}
func (s *Store) LoadScene(id string) (model.SceneState, error) {
	var scene model.SceneState
	err := s.get(bucketScenes, id, &scene)
	return scene, err
}
func (s *Store) DeleteScene(id string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return fmt.Errorf("store is closed")
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(bucketScenes)).Delete([]byte(id)) })
}
func (s *Store) AppendGesture(record model.GestureRecord) error {
	if record.ID == "" {
		return fmt.Errorf("gesture id required")
	}
	return s.put(bucketGestures, record.ID, record)
}
func (s *Store) LoadGesture(id string) (model.GestureRecord, error) {
	var record model.GestureRecord
	err := s.get(bucketGestures, id, &record)
	return record, err
}
func (s *Store) AppendControl(event model.ControlEvent) error {
	if event.ID == "" {
		return fmt.Errorf("control event id required")
	}
	return s.put(bucketControls, event.ID, event)
}
func (s *Store) LoadControl(id string) (model.ControlEvent, error) {
	var event model.ControlEvent
	err := s.get(bucketControls, id, &event)
	return event, err
}
