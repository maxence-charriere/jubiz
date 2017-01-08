package view

import (
	"github.com/maxence-charriere/jubiz"
	"github.com/maxence-charriere/jubiz/bin/jubiz/store"
	"github.com/murlokswarm/app"
	"github.com/murlokswarm/flux"
)

type HomeView struct {
	Articles    jubiz.ArticleList
	Downloading bool
	Min         bool
}

func (v *HomeView) OnMount() {
	store.Articles.Register(v)
}

func (v *HomeView) OnDismount() {
	store.Articles.Unregister(v)
}

func (v *HomeView) OnStoreEvent(e flux.Event) {
	switch e.Name {

	case store.LocalArticlesLoaded:
		v.Articles, _ = e.Payload.(jubiz.ArticleList)
		app.Render(v)

	case store.ArticlesDownloading:
		v.Downloading = true
		app.Render(v)

	case store.ArticlesDownloaded:
		if e.Error == nil {
			v.Articles, _ = e.Payload.(jubiz.ArticleList)
		}
		v.Downloading = false
		app.Render(v)
	}
}

func (v *HomeView) Render() string {
	v.Min = len(v.Articles) != 0

	return `
<div class="Home">
	{{range .Articles}}
		<div class="Home-Tile">
			<TileView Article="{{json .}}" />
		</div>
		
	{{end}}

	<div>
	{{if .Downloading}}
	<div class="Home-Download{{if .Min}}-Min{{end}}">
		<svg class="Home-Download{{if .Min}}-Min{{end}}-Icon" viewBox="0 0 50 50">
			<path d="M 25 2.5 C 21.974279 2.5 19.5 4.9742784 19.5 8 C 19.5 11.025721 21.974279 13.5 25 13.5 C 28.025721 13.5 30.5 11.025721 30.5 8 C 30.5 4.9742784 28.025721 2.5 25 2.5 z M 25 4.5 C 26.944841 4.5 28.5 6.0551588 28.5 8 C 28.5 9.9448407 26.944841 11.5 25 11.5 C 23.055159 11.5 21.5 9.9448407 21.5 8 C 21.5 6.0551588 23.055159 4.5 25 4.5 z M 13 8 C 10.250421 8 8 10.250421 8 13 C 8 15.749579 10.250421 18 13 18 C 15.749579 18 18 15.749579 18 13 C 18 10.250421 15.749579 8 13 8 z M 13 10 C 14.668699 10 16 11.331301 16 13 C 16 14.668699 14.668699 16 13 16 C 11.331301 16 10 14.668699 10 13 C 10 11.331301 11.331301 10 13 10 z M 37 11 C 35.907275 11 35 11.907275 35 13 C 35 14.092725 35.907275 15 37 15 C 38.092725 15 39 14.092725 39 13 C 39 11.907275 38.092725 11 37 11 z M 8 20.5 C 5.5265634 20.5 3.5 22.526563 3.5 25 C 3.5 27.473437 5.5265634 29.5 8 29.5 C 10.473437 29.5 12.5 27.473437 12.5 25 C 12.5 22.526563 10.473437 20.5 8 20.5 z M 8 22.5 C 9.3925566 22.5 10.5 23.607443 10.5 25 C 10.5 26.392557 9.3925566 27.5 8 27.5 C 6.6074434 27.5 5.5 26.392557 5.5 25 C 5.5 23.607443 6.6074434 22.5 8 22.5 z M 42 22.5 C 40.631133 22.5 39.5 23.631133 39.5 25 C 39.5 26.368867 40.631133 27.5 42 27.5 C 43.368867 27.5 44.5 26.368867 44.5 25 C 44.5 23.631133 43.368867 22.5 42 22.5 z M 42 24.5 C 42.287987 24.5 42.5 24.712013 42.5 25 C 42.5 25.287987 42.287987 25.5 42 25.5 C 41.712013 25.5 41.5 25.287987 41.5 25 C 41.5 24.712013 41.712013 24.5 42 24.5 z M 13 33 C 10.802706 33 9 34.802706 9 37 C 9 39.197294 10.802706 41 13 41 C 15.197294 41 17 39.197294 17 37 C 17 34.802706 15.197294 33 13 33 z M 37 34 C 35.35499 34 34 35.35499 34 37 C 34 38.64501 35.35499 40 37 40 C 38.64501 40 40 38.64501 40 37 C 40 35.35499 38.64501 34 37 34 z M 13 35 C 14.116414 35 15 35.883586 15 37 C 15 38.116414 14.116414 39 13 39 C 11.883586 39 11 38.116414 11 37 C 11 35.883586 11.883586 35 13 35 z M 37 36 C 37.564128 36 38 36.435872 38 37 C 38 37.564128 37.564128 38 37 38 C 36.435872 38 36 37.564128 36 37 C 36 36.435872 36.435872 36 37 36 z M 25 38.5 C 23.078848 38.5 21.5 40.078848 21.5 42 C 21.5 43.921152 23.078848 45.5 25 45.5 C 26.921152 45.5 28.5 43.921152 28.5 42 C 28.5 40.078848 26.921152 38.5 25 38.5 z M 25 40.5 C 25.840272 40.5 26.5 41.159728 26.5 42 C 26.5 42.840272 25.840272 43.5 25 43.5 C 24.159728 43.5 23.5 42.840272 23.5 42 C 23.5 41.159728 24.159728 40.5 25 40.5 z" />
		</svg>
		<h1 class="Home-Download{{if .Min}}-Min{{end}}-Title">downloading</h1>
	</div>
	{{end}}
	</div>
</div>
    `
}
