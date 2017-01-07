package store

import "github.com/murlokswarm/flux"

const (
	NavShowDetail       = "nav-show-detail"
	NavHideDetail       = "nav-hide-detail"
	NavToggleFullScreen = "nav-toggle-fullscreen"
)

type navStore struct {
	flux.Store
}

func (s *navStore) OnDispatch(a flux.Action) error {
	switch a.Name {
	case NavShowDetail, NavHideDetail, NavToggleFullScreen:
		s.Emit(flux.Event{
			Name:    a.Name,
			Payload: a.Payload,
		})
	}
	return nil
}
