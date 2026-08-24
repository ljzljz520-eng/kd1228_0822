package persistence

import (
	"fmt"
	"go.etcd.io/bbolt"
	"jellyfield/model"
)

func (s *Store) SceneIDs() ([]string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return nil, fmt.Errorf("store is closed")
	}
	ids := []string{}
	err := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(bucketScenes)).ForEach(func(k, v []byte) error {
			if v != nil {
				ids = append(ids, string(k))
			}
			return nil
		})
	})
	return ids, err
}
func (s *Store) GestureIDs() ([]string, error) { return s.ids(bucketGestures) }
func (s *Store) ControlIDs() ([]string, error) { return s.ids(bucketControls) }
func (s *Store) ids(bucket string) ([]string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return nil, fmt.Errorf("store is closed")
	}
	result := []string{}
	err := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(bucket)).ForEach(func(k, v []byte) error {
			if v != nil {
				result = append(result, string(k))
			}
			return nil
		})
	})
	return result, err
}
func (s *Store) FindScenes(query model.Query) ([]model.SceneState, error) {
	ids, err := s.SceneIDs()
	if err != nil {
		return nil, err
	}
	result := []model.SceneState{}
	for _, id := range ids {
		scene, loadErr := s.LoadScene(id)
		if loadErr != nil {
			continue
		}
		if query.SceneID != "" && query.SceneID != scene.ID {
			continue
		}
		matched := false
		for _, j := range scene.Jellyfish {
			if query.Matches(j) {
				matched = true
				break
			}
		}
		if matched {
			result = append(result, scene)
		}
	}
	return result, nil
}
