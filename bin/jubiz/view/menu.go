package view

import "github.com/murlokswarm/flux"
import "github.com/maxence-charriere/jubiz/bin/jubiz/store"

type MenuBar struct {
}

func (m *MenuBar) Render() string {
	return `
<menu>
    <menu label="app">
        <menuitem label="About" selector="orderFrontStandardAboutPanel:" separator="true" />        
        <menuitem label="Quit" shortcut="meta+q" selector="terminate:" />        
    </menu>
    <EditMenu />
    <WindowMenu />
</menu>
    `
}

type WindowMenu struct{}

func (m *WindowMenu) Render() string {
	return `
<menu label="Window">
    <menuitem label="Close" _onclick="OnCloseClick" shortcut="meta+w" />
</menu>
    `
}

func (m *WindowMenu) OnCloseClick() {
	flux.Dispatch(flux.Action{
		Name: store.NavClose,
	})
}

type EditMenu struct{}

func (m *EditMenu) Render() string {
	return `
<menu label="Edit">
    <menuitem label="Cut" selector="cut:" shortcut="meta+x" />
    <menuitem label="Copy" selector="copy:" shortcut="meta+c" />
    <menuitem label="Paste" selector="paste:" shortcut="meta+v" separator="true" />
    <menuitem label="Select All" selector="selectAll:" shortcut="meta+a" />
</menu>
    `
}
