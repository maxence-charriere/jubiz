package view

import (
	"math/rand"
	"time"

	"github.com/maxence-charriere/jubiz"
	"github.com/maxence-charriere/jubiz/bin/jubiz/store"
	"github.com/murlokswarm/app"
	"github.com/murlokswarm/flux"
)

type DetailView struct {
	Article        jubiz.Article
	HeroBackground jubiz.Image
	TimeSpan       string
	FullScreen     bool
}

func (v *DetailView) OnMount() {
	store.Nav.Register(v)
	store.Articles.Register(v)
}

func (v *DetailView) OnDismount() {
	store.Nav.Unregister(v)
	store.Articles.Unregister(v)
}

func (v *DetailView) OnStoreEvent(e flux.Event) {
	switch e.Name {
	case store.NavShowDetail:
		v.Article = e.Payload.(jubiz.Article)

	case store.NavHideDetail:
		v.Article = jubiz.Article{}

	case store.NavToggleFullScreen:
		v.FullScreen = e.Payload.(bool)
	}
	app.Render(v)
}

func (v *DetailView) Render() string {
	v.TimeSpan = timeSpanFromNow(v.Article.PubDate)

	if l := len(v.Article.Images); l > 0 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		idx := r.Intn(l)
		v.HeroBackground = v.Article.Images[idx]
	}

	return `
<div class="Detail">
    <button class="Detail-Close{{if .FullScreen}} Detail-CloseFullScreen{{end}}" title="Retour" _onclick="OnCloseClicked">
		<svg viewBox="0 0 50 50">
			<path d="M 19.8125 13.09375 A 1.0001 1.0001 0 0 0 19.25 13.40625 L 8.34375 24.28125 L 7.65625 25 L 8.34375 25.71875 L 19.25 36.59375 A 1.0001 1.0001 0 1 0 20.65625 35.1875 L 11.46875 26 L 41 26 A 1.0001 1.0001 0 1 0 41 24 L 11.46875 24 L 20.65625 14.8125 A 1.0001 1.0001 0 0 0 19.8125 13.09375 z" />
		</svg>
	</button>
	<div class="Detail-Hero" style="background-image:url('{{.HeroBackground.URL}}')">
		<div class="Detail-HeroBox">
			<div class="Detail-HeroContent">
				{{if ge (len .Article.Categories) 1}}
					<span class="Detail-Category">{{index .Article.Categories 0}}</span>
					<div class="Detail-ContentSep"></div>
				{{end}}
        		<h1 class="Detail-Title">{{html .Article.Title}}</h1>
				<span class="Detail-Time">Il y a {{html .TimeSpan}}</span>
        		<div class="Detail-Text">{{.Article.Text}}</div>
				<p>Written by <a href="{{html .Article.Author.TwitterURL}}">@{{html .Article.Author.Name}}</a></p>
				<div class="Detail-Actions">
					<button title="fubiz.net">
						<a href="{{html .Article.URL}}">
							<svg class="Detail-Icon" viewBox="0 0 50 50">
								<path d="M24.6,29.6c0.9,2.5,0.3,5.3-1.6,7.3l-6,6c-1.3,1.3-3.1,2-4.9,2c-1.9,0-3.6-0.7-4.9-2c-2.7-2.7-2.7-7.2,0-9.9l6-6	c1.3-1.3,3.1-2,4.9-2c0.8,0,1.6,0.1,2.3,0.4l1.5-1.5C20.7,23.3,19.3,23,18,23c-2.3,0-4.6,0.9-6.4,2.6l-6,6c-3.5,3.5-3.5,9.2,0,12.7 C7.4,46.1,9.7,47,12,47s4.6-0.9,6.4-2.6l6-6c2.8-2.8,3.3-6.9,1.7-10.2L24.6,29.6z M44.4,5.6C42.6,3.9,40.3,3,38,3s-4.6,0.9-6.4,2.6	l-6,6c-2.8,2.8-3.3,6.9-1.7,10.2l1.5-1.5C24.5,17.8,25.1,15,27,13l6-6c1.3-1.3,3.1-2,4.9-2c1.9,0,3.6,0.7,4.9,2 c2.7,2.7,2.7,7.2,0,9.9l-6,6c-1.3,1.3-3.1,2-4.9,2c-0.8,0-1.6-0.1-2.3-0.4L28.1,26c1.2,0.6,2.5,0.9,3.9,0.9c2.3,0,4.6-0.9,6.4-2.6	l6-6C47.9,14.9,47.9,9.1,44.4,5.6z M32.1,17.9c-0.2-0.2-0.4-0.3-0.7-0.3c-0.2,0-0.5,0.1-0.7,0.3L18,30.6c-0.4,0.4-0.4,1,0,1.4 c0.2,0.2,0.5,0.3,0.7,0.3s0.5-0.1,0.7-0.3l12.7-12.7C32.5,18.9,32.5,18.3,32.1,17.9z" />
							</svg>
						</a>
					</button>
					<button title="Partager" _onmousedown="OnShareClicked">
						<svg class="Detail-Icon" viewBox="0 0 50 50">
							<path d="M 25 0.59375 L 24.28125 1.28125 L 16.28125 9.28125 A 1.016466 1.016466 0 1 0 17.71875 10.71875 L 24 4.4375 L 24 32 A 1.0001 1.0001 0 1 0 26 32 L 26 4.4375 L 32.28125 10.71875 A 1.016466 1.016466 0 1 0 33.71875 9.28125 L 25.71875 1.28125 L 25 0.59375 z M 7 16 L 7 17 L 7 49 L 7 50 L 8 50 L 42 50 L 43 50 L 43 49 L 43 17 L 43 16 L 42 16 L 33 16 A 1.0001 1.0001 0 1 0 33 18 L 41 18 L 41 48 L 9 48 L 9 18 L 17 18 A 1.0001 1.0001 0 1 0 17 16 L 8 16 L 7 16 z" />
						</svg>
					</button>
				</div>
			</div>
		</div>
	</div>

	{{range .Article.Images}}
		<img class="Detail-Image" src="{{.URL}}"  alt="{{.URL}}"/>
	{{end}}

	{{range  .Article.Videos}}
		<iframe class="Detail-Video" src="{{.URL}}"></iframe>
	{{end}}
</div>
    `
}

func (v *DetailView) OnCloseClicked() {
	flux.Dispatch(flux.Action{
		Name: store.NavHideDetail,
	})
}

func (v *DetailView) OnShareClicked() {
	flux.Dispatch(flux.Action{
		Name:    store.ShareArticle,
		Payload: v.Article,
	})
}
