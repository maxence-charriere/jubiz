package view

import (
	"github.com/maxence-charriere/jubiz/bin/jubiz/store"
	"github.com/murlokswarm/app"
	"github.com/murlokswarm/flux"
)

type NavView struct {
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
		v.DetailVisibilty = "Nav-Detail-Show"

	case store.NavHideDetail:
		v.DetailVisibilty = "Nav-Detail-Hidden"
	}
	app.Render(v)
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
