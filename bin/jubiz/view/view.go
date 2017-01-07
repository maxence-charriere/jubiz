package view

import "github.com/murlokswarm/app"

func init() {
	app.RegisterComponent(&MenuBar{})
	app.RegisterComponent(&WindowMenu{})
	app.RegisterComponent(&NavView{})
	app.RegisterComponent(&HomeView{})
	app.RegisterComponent(&DetailView{})
	app.RegisterComponent(&TileView{})
}
