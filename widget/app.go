package widget

import (
	"Domblr/res"
	"Domblr/util"
	_ "embed"
	"text/template"
)

type Config struct {
	Title string
	CSS   string
	HTML  string
}

// App widget
type App struct {
	Config Config
	Body   Widget
	// Style contains Variables for setting the look of widgets
	// Nullable after Build
	Style map[int]string
}

func (p *App) Build(ctx *BuildContext) {
	// Initialize the Node
	// TODO consider making App a Widget and make use of given default style
	util.Panic(p.Body.Build(ctx))
}

func (p *App) Render(ctx *RenderContext) {
	// Render the CSS and HTML code
	p.Body.Render(ctx)
	p.Config.CSS = ctx.CSS.String()
	p.Config.HTML = ctx.HTML.String()

	// Write the boilerplate document and the rendered CSS/HTML
	tmpl, err := template.New("boilerplate").Parse(res.BoilerplateHTML)
	util.Panic(err)
	err = tmpl.Execute(ctx.Buffer, p.Config)
	util.Panic(err)
}
