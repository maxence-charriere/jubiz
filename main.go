package main

import (
	"github.com/murlokswarm/app"
	"github.com/murlokswarm/flux"
	_ "github.com/murlokswarm/mac"
)

var (
	navStore      = &NavStore{}
	articlesStore = &ArticlesStore{}

	mainWindow app.Windower
)

func init() {
	flux.Register(navStore)
	flux.Register(articlesStore)
}

func main() {
	app.OnLaunch = func() {
		menuBar := &MenuBar{}
		app.MenuBar().Mount(menuBar)

		mainWindow = newMainWindow()
		mainWindow.Mount(&NavView{})

		flux.Dispatch(
			flux.Action{Name: "articles-read"},
			flux.Action{Name: "articles-get"},
			flux.Action{Name: "articles-save"},
		)
	}
	app.OnReopen = func(hasVisibleWindow bool) {
		if mainWindow != nil {
			return
		}

		mainWindow = newMainWindow()
		mainWindow.Mount(&NavView{})

		flux.Dispatch(
			flux.Action{Name: "articles-read"},
			flux.Action{Name: "articles-get"},
			flux.Action{Name: "articles-save"},
		)
	}
	app.Run()
}

func newMainWindow() app.Windower {
	return app.NewWindow(app.Window{
		Title:           "Jubiz",
		TitlebarHidden:  true,
		Width:           1280,
		Height:          720,
		BackgroundColor: "#1e1e1e",
		OnClose: func() bool {
			mainWindow = nil
			return true
		},
	})
}
