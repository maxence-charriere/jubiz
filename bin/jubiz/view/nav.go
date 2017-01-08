package view

import (
	"github.com/maxence-charriere/jubiz/bin/jubiz/store"
	"github.com/murlokswarm/app"
	"github.com/murlokswarm/flux"
)

const (
	NavDefaultState = iota
	NavDetailState
)

type NavView struct {
	State           int
	DetailVisibilty string
}

func (v *NavView) OnMount() {
	store.Nav.Register(v)
}

func (v *NavView) OnDismount() {
	store.Nav.Unregister(v)
}

func (v *NavView) OnStoreEvent(e flux.Event) {
	switch e.Name {
	case store.NavShowDetail:
		v.State = NavDetailState
		v.DetailVisibilty = "Nav-Detail-Show"
		app.Render(v)

	case store.NavHideDetail:
		v.State = NavDefaultState
		v.DetailVisibilty = "Nav-Detail-Hidden"
		app.Render(v)

	case store.NavClose:
		if v.State == NavDefaultState {
			win := app.Context(v).(app.Windower)
			win.Close()
		}
	}
}

func (v *NavView) Render() string {
	return `
<div class="Nav">
    <div class="Nav-Home">
        <HomeView />
    </div>
	<div class="Nav-HomeTitleBar" />
    <div class="Nav-Detail {{.DetailVisibilty}}">
		<DetailView />
	</div>
</div>
    `
}
