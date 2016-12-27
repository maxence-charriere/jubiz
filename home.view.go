package main

import (
	"github.com/murlokswarm/app"
	"github.com/murlokswarm/flux"
)

type HomeView struct {
	Articles articleList
}

func (v *HomeView) OnMount() {
	articlesStore.Register(v)
}

func (v *HomeView) OnDismount() {
	articlesStore.Unregister(v)
}

func (v *HomeView) OnStoreEvent(e flux.Event) {
	switch e.Name {
	case articlesRead, articlesGet:
		if articles, ok := e.Payload.(articleList); ok {
			v.Articles = articles
			app.Render(v)
		}
	}
}

func (v *HomeView) Render() string {
	return `
<div class="Home">
	{{range .Articles}}
		<div class="Home-Tile">
			<TileView Article="{{json .}}" />
		</div>
	{{end}}
</div>
    `
}

func init() {
	app.RegisterComponent(&HomeView{})
}
