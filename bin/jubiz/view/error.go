package view

import (
	"github.com/maxence-charriere/jubiz/bin/jubiz/store"
	"github.com/murlokswarm/app"
	"github.com/murlokswarm/flux"
)

type Error struct {
	Message string
	Action  *flux.Action
}

func (v *Error) OnMount() {
	store.Articles.Register(v)
}

func (v *Error) OnDismount() {
	store.Articles.Unregister(v)
}

func (v *Error) OnStoreEvent(e flux.Event) {
	if e.Error == nil {
		return
	}

	switch e.Name {
	case store.ArticlesDownloaded:
		v.Message = "Oops, internet is gone."
		v.Action = &flux.Action{
			Name: store.DownloadArticles,
		}
		app.Render(v)

		flux.Dispatch(flux.Action{
			Name:    store.NavToggleError,
			Payload: true,
		})

	case store.LocalArticlesLoaded, store.ArticlesSaved:
		v.Message = e.Error.Error()
		app.Render(v)

		flux.Dispatch(flux.Action{
			Name:    store.NavToggleError,
			Payload: true,
		})
	}
}

func (v *Error) Render() string {
	return `
<div class="Error">
    <div>{{html .Message}}</div>
    <div>
        {{if .Action}}
        <button title="Retry" onmousedown="OnRetryClicked">
            <svg class="Error-Icon" viewBox="0 0 50 50">
                <path d="M 39.90625 5.96875 A 1.0001 1.0001 0 0 0 39.78125 6 A 1.0001 1.0001 0 0 0 39 7 L 39 13.6875 C 38.79507 13.43207 38.592962 13.179445 38.375 12.9375 C 35.081942 9.2866616 30.299299 7 25 7 C 15.069709 7 7 15.069709 7 25 A 1.0001 1.0001 0 1 0 9 25 C 9 16.150291 16.150291 9 25 9 C 29.722701 9 33.946058 11.034088 36.875 14.28125 C 37.082376 14.511444 37.273219 14.754803 37.46875 15 L 31 15 A 1.0001 1.0001 0 1 0 31 17 L 40 17 L 41 17 L 41 16 L 41 7 A 1.0001 1.0001 0 0 0 39.90625 5.96875 z M 41.90625 23.96875 A 1.0001 1.0001 0 0 0 41.78125 24 A 1.0001 1.0001 0 0 0 41 25 C 41 33.849709 33.849709 41 25 41 C 20.277432 41 16.054012 38.965048 13.125 35.71875 C 12.917373 35.487941 12.726795 35.245209 12.53125 35 L 19 35 A 1.0001 1.0001 0 1 0 19 33 L 10 33 L 9 33 L 9 34 L 9 43 A 1.0001 1.0001 0 1 0 11 43 L 11 36.3125 C 11.204906 36.567893 11.407273 36.820463 11.625 37.0625 C 14.917988 40.712202 19.700568 43 25 43 C 34.930291 43 43 34.930291 43 25 A 1.0001 1.0001 0 0 0 41.90625 23.96875 z" />
            </svg>
        </button>
        {{end}}
        
        <button title="Hide" onmousedown="OnCloseClicked">
            <svg class="Error-Icon" viewBox="0 0 50 50">
                <path d="M 7.71875 6.28125 L 6.28125 7.71875 L 23.5625 25 L 6.28125 42.28125 L 7.71875 43.71875 L 25 26.4375 L 42.28125 43.71875 L 43.71875 42.28125 L 26.4375 25 L 43.71875 7.71875 L 42.28125 6.28125 L 25 23.5625 L 7.71875 6.28125 z" />
            </svg>
        </button>
    </div>
</div>
    `
}

func (v *Error) OnRetryClicked() {
	flux.Dispatch(
		flux.Action{
			Name:    store.NavToggleError,
			Payload: false,
		},
		*v.Action,
	)
}

func (v *Error) OnCloseClicked() {
	flux.Dispatch(flux.Action{
		Name:    store.NavToggleError,
		Payload: false,
	})
}
