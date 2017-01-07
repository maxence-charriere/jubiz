package view

type MenuBar struct {
}

func (m *MenuBar) Render() string {
	return `
<menu>
    <menu label="app">
        <menuitem label="Quit" shortcut="meta+q" selector="terminate:" />        
    </menu>
    <WindowMenu />
</menu>
    `
}

type WindowMenu struct {
}

func (m *WindowMenu) Render() string {
	return `
<menu label="Window">
    <menuitem label="Close" selector="performClose:" shortcut="meta+w" />
</menu>
    `
}
