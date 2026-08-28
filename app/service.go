package app

import (
	"fmt"
	"jellyfield/control"
	"jellyfield/gesture"
	"jellyfield/model"
	"jellyfield/persistence"
	"jellyfield/presentation"
	"jellyfield/simulation"
)

type Service struct {
	scene      model.SceneState
	engine     *simulation.Engine
	store      *persistence.Store
	recognizer gesture.Recognizer
	translator *gesture.Translator
	panel      *control.Panel
	renderer   presentation.Renderer
}

func New(scene model.SceneState, store *persistence.Store) (*Service, error) {
	engine, err := simulation.NewEngine(scene)
	if err != nil {
		return nil, err
	}
	return &Service{scene: scene, engine: engine, store: store, recognizer: gesture.NewRecognizer(scene.Width, scene.Height), translator: gesture.NewTranslator(), panel: control.NewPanel(engine), renderer: presentation.NewRenderer(false)}, nil
}
func (s *Service) Scene() model.SceneState { return s.scene }
func (s *Service) RegisterJellyfish(j model.Jellyfish) error {
	if err := j.Validate(); err != nil {
		return err
	}
	for i, current := range s.scene.Jellyfish {
		if current.ID == j.ID {
			s.scene.Jellyfish[i] = j
			return s.engine.ApplyJellyfish(j)
		}
	}
	s.scene.Jellyfish = append(s.scene.Jellyfish, j)
	return s.engine.ApplyJellyfish(j)
}
func (s *Service) Save() error {
	if s.store == nil {
		return fmt.Errorf("service has no store")
	}
	return s.store.SaveScene(s.scene)
}
func (s *Service) Restore() error {
	if s.store == nil {
		return fmt.Errorf("service has no store")
	}
	scene, err := s.store.LoadScene(s.scene.ID)
	if err != nil {
		return err
	}
	engine, err := simulation.NewEngine(scene)
	if err != nil {
		return err
	}
	s.scene = scene
	s.engine = engine
	s.panel = control.NewPanel(engine)
	return nil
}
func (s *Service) HandleGesture(g model.Gesture) (model.ControlEvent, error) {
	command, record, err := s.translator.Translate(g)
	if err != nil {
		return model.ControlEvent{}, err
	}
	event, err := s.panel.Apply(command)
	if err != nil {
		return model.ControlEvent{}, err
	}
	if g.Kind == model.GestureSelect {
		s.scene.SelectedID = g.TargetID
	}
	if s.store != nil {
		if err := s.store.AppendGesture(record); err != nil {
			return model.ControlEvent{}, err
		}
		if err := s.store.AppendControl(event); err != nil {
			return model.ControlEvent{}, err
		}
	}
	return event, nil
}
func (s *Service) Step() model.RenderFrame {
	frame := s.engine.Step()
	frame.SelectedID = s.scene.SelectedID
	for i := range frame.Jellyfish {
		for j := range s.scene.Jellyfish {
			if frame.Jellyfish[i].ID == s.scene.Jellyfish[j].ID {
				frame.Jellyfish[i].Name = s.scene.Jellyfish[j].Name
				frame.Jellyfish[i].Position = s.scene.Jellyfish[j].Position
			}
		}
	}
	return frame
}
func (s *Service) Render() (string, error) {
	frame := s.engine.Render(s.scene.SelectedID)
	return s.renderer.JSON(frame)
}
func (s *Service) Summary() string { return s.renderer.Summary(s.engine.Render(s.scene.SelectedID)) }
func (s *Service) SelectAt(point model.Vector) error {
	g, err := s.recognizer.Select(s.scene, point)
	if err != nil {
		return err
	}
	_, err = s.HandleGesture(g)
	return err
}
func (s *Service) SetSpeedAt(point model.Vector, amount float64) error {
	g, err := s.recognizer.Speed(s.scene, point, amount)
	if err != nil {
		return err
	}
	_, err = s.HandleGesture(g)
	return err
}
func (s *Service) SetColorAt(point model.Vector, color model.Color) error {
	g, err := s.recognizer.Color(s.scene, point, color)
	if err != nil {
		return err
	}
	_, err = s.HandleGesture(g)
	return err
}
func (s *Service) Query(query model.Query) []model.Jellyfish {
	result := []model.Jellyfish{}
	for _, j := range s.scene.Jellyfish {
		if query.Matches(j) {
			result = append(result, j)
		}
	}
	return result
}
func (s *Service) Metrics() simulation.Metrics { return s.engine.Measure() }

func (s *Service) SetRenderer(renderer presentation.Renderer) { s.renderer = renderer }

func (s *Service) SelectedJellyfish() (model.Jellyfish, bool) {
	if s.scene.SelectedID == "" {
		return model.Jellyfish{}, false
	}
	return s.scene.Find(s.scene.SelectedID)
}

func (s *Service) ApplyCommand(command model.ControlCommand) error {
	event, err := s.panel.Apply(command)
	if err != nil {
		return err
	}
	if command.Select {
		s.scene.SelectedID = command.TargetID
	}
	if s.store != nil {
		return s.store.AppendControl(event)
	}
	return nil
}
