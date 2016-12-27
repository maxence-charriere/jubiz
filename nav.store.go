package main

import "github.com/murlokswarm/flux"

type NavStore struct {
	flux.Store
}

func (s *NavStore) OnDispatch(a flux.Action) {
	switch a.Name {
	case "nav-detail-show":
		s.Emit(flux.Event{
			Name: "nav-detail-show",
		})

	case "nav-detail-hide":
		s.Emit(flux.Event{
			Name: "nav-detail-hide",
		})
	}
}
