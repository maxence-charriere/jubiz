package view

import (
	"github.com/maxence-charriere/jubiz"
	"github.com/maxence-charriere/jubiz/bin/jubiz/store"
	"github.com/murlokswarm/app"
	"github.com/murlokswarm/flux"
)

type HomeView struct {
	Articles jubiz.ArticleList
}

func (v *HomeView) OnMount() {
	store.Articles.Register(v)
}

func (v *HomeView) OnDismount() {
	store.Articles.Unregister(v)
}

func (v *HomeView) OnStoreEvent(e flux.Event) {
	switch e.Name {
	case store.LocalArticlesLoaded, store.ArticlesDownloaded:
		if articles, ok := e.Payload.(jubiz.ArticleList); ok {
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
