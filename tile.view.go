package main

import (
	"github.com/murlokswarm/app"
	"github.com/murlokswarm/flux"
)

type TileView struct {
	Article  article
	TimeSpan string
}

func (v *TileView) OnMount() {
	articlesStore.Register(v)
}

func (v *TileView) OnDismount() {
	articlesStore.Unregister(v)
}

func (v *TileView) OnStoreEvent(e flux.Event) {
	switch e.Name {
	case articleSet:
		if a := e.Payload.(article); a.URL == v.Article.URL {
			v.Article = a
			app.Render(v)
		}
	}
}

func (v *TileView) Render() string {
	v.TimeSpan = timeSpanFromNow(v.Article.PubDate)

	return `
<div class="Tile" _onclick="OnClick">
	{{if ge (len .Article.Images) 1 }}
		<div class="Tile-Background" style="background-image:url('{{(index .Article.Images 0).URL}}')"></div>
		{{if eq (len .Article.Images) 1 }}
			<div class="Tile-Background Tile-BackgroundHover" style="background-image:url('{{(index .Article.Images 0).URL}}')"></div>
		{{else}}
			<div class=" Tile-Background Tile-BackgroundHover" style="background-image:url('{{(index .Article.Images 1).URL}}')"></div>
		{{end}}
	{{end}}
	<div class="Tile-Content">
		{{if ge (len .Article.Categories) 1}}
			<span class="Tile-Category">{{index .Article.Categories 0}}</span>
			<div class="Tile-ContentSep"></div>
		{{end}}
    	<h1 class="Tile-Title">{{html .Article.Title}}</h1>
		<span class="Tile-Time">
			<div class="{{if not .Article.Read}}Tile-NotRead{{end}}"></div>
			Il y a {{html .TimeSpan}}
		</span>
	</div>
</div>
    `
}

func (v *TileView) OnClick() {
	v.Article.Read = true

	flux.Dispatch(
		flux.Action{
			Name:    articleDisplay,
			Payload: v.Article,
		},
		flux.Action{
			Name: navDetailShow,
		},
		flux.Action{
			Name:    articleSet,
			Payload: v.Article,
		},
		flux.Action{
			Name: articlesSave,
		},
	)
}

func init() {
	app.RegisterComponent(&TileView{})
}
