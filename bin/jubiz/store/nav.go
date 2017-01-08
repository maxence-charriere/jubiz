package store

import (
	"github.com/murlokswarm/app"
	"github.com/murlokswarm/flux"
)

const (
	NavShowDetail       = "nav-show-detail"
	NavHideDetail       = "nav-hide-detail"
	NavToggleError      = "nav-toggle-error"
	NavToggleFullScreen = "nav-toggle-fullscreen"
	NavClose            = "nav-close"
)

type navStore struct {
	flux.Store
}

func (s *navStore) OnDispatch(a flux.Action) error {
	switch a.Name {
	case NavShowDetail, NavHideDetail, NavToggleFullScreen, NavToggleError, NavClose:
		s.Emit(flux.Event{
			Name:    a.Name,
			Payload: a.Payload,
		})

		app.Dock().SetBadge(nil)
	}
	return nil
}
