package main

import (
	"github.com/murlokswarm/app"
	"github.com/murlokswarm/flux"
)

var (
	navDetailShow = "nav-detail-show"
	navDetailHide = "nav-detail-hide"
)

type NavView struct {
	DetailVisibilty string
}

func (v *NavView) OnMount() {
	navStore.Register(v)
}

func (v *NavView) OnDismount() {
	navStore.Unregister(v)
}

func (v *NavView) OnStoreEvent(e flux.Event) {
	switch e.Name {
	case navDetailShow:
		v.ShowDetail()

	case navDetailHide:
		v.HideDetail()
	}
}

func (v *NavView) ShowDetail() {
	v.DetailVisibilty = "Nav-Detail-Show"
	app.Render(v)
}

func (v *NavView) HideDetail() {
	v.DetailVisibilty = "Nav-Detail-Hidden"
	app.Render(v)
}

func (v *NavView) Render() string {
	return `
<div class="Nav">
    <div class="Nav-Home">
        <HomeView />
    </div>
    <div class="Nav-Detail {{.DetailVisibilty}}">
		<DetailView />
	</div>
</div>
    `
}

func init() {
	app.RegisterComponent(&NavView{})
}
