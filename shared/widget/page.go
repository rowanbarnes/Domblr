package widget

import (
	"Domblr/res"
	"Domblr/util"
	"bytes"
	_ "embed"
	"text/template"
)

type Config struct {
	Title string
	CSS   string
	HTML  string
}

// Page widget
type Page struct {
	Config Config
	Body   Widget
	// Style contains Variables for setting the look of widgets
	// Nullable after Setup
	Style map[int]string
}

func (p *Page) Setup(style map[int]string) {
	// Initialize the Node
	// TODO consider making Page a Widget and make use of given default style
	util.Panic(p.Body.Setup(nil, 0))
}

func (p *Page) Render(buffer *bytes.Buffer) {
	// Render the css and html code
	var css, html bytes.Buffer
	p.Body.Render(&css, &html)
	p.Config.CSS = css.String()
	p.Config.HTML = html.String()

	// Write the boilerplate document and the rendered css/html
	tmpl, err := template.New("boilerplate").Parse(res.BoilerplateHTML)
	util.Panic(err)
	err = tmpl.Execute(buffer, p.Config)
	util.Panic(err)
}
