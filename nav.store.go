package main

import "github.com/murlokswarm/flux"

const (
	navDetailShow    = "nav-detail-show"
	navDetailHide    = "nav-detail-hide"
	navFullScreenON  = "nav-fullscreen-on"
	navFullScreenOFF = "nav-fullscreen-off"
)

type NavStore struct {
	flux.Store
}

func (s *NavStore) OnDispatch(a flux.Action) error {
	switch a.Name {
	case navDetailHide, navDetailShow, navFullScreenOFF, navFullScreenON:
		s.Emit(flux.Event{
			Name: a.Name,
		})
	}
	return nil
}
