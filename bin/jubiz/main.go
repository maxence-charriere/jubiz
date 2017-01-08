package main

// Driver import
import (
	_ "github.com/murlokswarm/mac"
)

import (
	"time"

	"github.com/maxence-charriere/jubiz/bin/jubiz/store"
	"github.com/maxence-charriere/jubiz/bin/jubiz/view"
	"github.com/murlokswarm/app"
	"github.com/murlokswarm/flux"
)

var (
	mainWindow app.Windower
)

func main() {
	app.OnLaunch = func() {
		menuBar := &view.MenuBar{}
		app.MenuBar().Mount(menuBar)
		mainWindow = newMainWindow()

		go startBackgroundDownload(time.Minute * 20)
	}
	app.OnReopen = func(hasVisibleWindow bool) {
		if mainWindow != nil {
			return
		}
		mainWindow = newMainWindow()
	}
	app.Run()
}

func newMainWindow() app.Windower {
	win := app.NewWindow(app.Window{
		Title:           "Jubiz",
		TitlebarHidden:  true,
		Width:           1280,
		Height:          720,
		BackgroundColor: "#1e1e1e",
		OnClose: func() bool {
			mainWindow = nil
			return true
		},
		OnFullScreen: func() {
			flux.Dispatch(flux.Action{
				Name:    store.NavToggleFullScreen,
				Payload: true,
			})
		},
		OnExitFullScreen: func() {
			flux.Dispatch(flux.Action{
				Name:    store.NavToggleFullScreen,
				Payload: false,
			})
		},
	})
	win.Mount(&view.NavView{})

	flux.Dispatch(
		flux.Action{Name: store.LoadLocalArticles},
		flux.Action{Name: store.DownloadArticles},
		flux.Action{Name: store.SaveArticles},
	)
	return win
}

func startBackgroundDownload(interval time.Duration) {
	t := time.Tick(interval)

	for range t {
		flux.Dispatch(
			flux.Action{Name: store.DownloadArticles},
			flux.Action{Name: store.SaveArticles},
		)
	}
}
