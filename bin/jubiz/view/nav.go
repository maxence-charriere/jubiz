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
	State            int
	DetailVisibility string
	ErrorVisibility  string
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
		v.DetailVisibility = "Nav-Detail-Show"
		app.Render(v)

	case store.NavHideDetail:
		v.State = NavDefaultState
		v.DetailVisibility = "Nav-Detail-Hidden"
		app.Render(v)

	case store.NavToggleError:
		if showErr, ok := e.Payload.(bool); ok && showErr {
			v.ErrorVisibility = "Nav-Error-Show"
		} else {
			v.ErrorVisibility = "Nav-Error-Hidden"
		}
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
    <div class="Nav-Detail {{.DetailVisibility}}">
		<DetailView />
	</div>
	<div class="Nav-Error {{.ErrorVisibility}}">
		<Error />
	</div>
</div>
    `
}
