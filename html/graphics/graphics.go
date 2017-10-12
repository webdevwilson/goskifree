package graphics

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/webdevwilson/goskifree/core/graphics"
	"github.com/webdevwilson/goskifree/html/webgl"
	"honnef.co/go/js/dom"
)

type WebGLCanvas struct {
	document dom.Document
	console  *js.Object
	context  *webgl.Context
	messages *dom.HTMLSpanElement
}

var _ graphics.Canvas = &WebGLCanvas{}

func NewWebGLCanvas() graphics.Canvas {
	// initialize the body by creating the following structure:
	// <div class="container">
	//   <canvas id="canvas"></canvas>
	//	 <div id="overlay"><span></span></div>
	// </div>
	document := dom.GetWindow().Document()
	body := document.GetElementsByTagName("body")[0].(*dom.HTMLBodyElement)
	body.SetAttribute("style", "margin: 0; padding: 0")
	container := document.CreateElement("div").(*dom.HTMLDivElement)
	container.SetClass("container")
	body.AppendChild(container)

	// create the canvas element
	canvas := document.CreateElement("canvas").(*dom.HTMLCanvasElement)
	canvas.SetID("canvas")
	canvas.SetAttribute("style", "width: 100vw; height: 100vh; display: block;")
	container.AppendChild(canvas)

	// create our text overlay
	overlay := document.CreateElement("div")
	overlay.SetID("overlay")
	overlay.SetAttribute("style", "position: absolute; height: 1.6em; width: 100vw; bottom: 0; color: white")
	container.AppendChild(overlay)

	text := document.CreateElement("span").(*dom.HTMLSpanElement)
	overlay.AppendChild(text)

	// initialize the webgl context
	attrs := webgl.DefaultAttributes()
	attrs.Alpha = false

	gl, err := webgl.NewContext(canvas.Underlying(), attrs)
	if err != nil {
		js.Global.Call("alert", "Error: "+err.Error())
	}

	return &WebGLCanvas{
		document: document,
		console:  js.Global.Get("console"),
		context:  gl,
		messages: text,
	}
}

func (g *WebGLCanvas) Write(text string) {
	g.console.Call("log", text)
	g.messages.SetInnerHTML(text)
}

func (g *WebGLCanvas) Draw(sprite graphics.Sprite, point graphics.Point) {

}

func (g *WebGLCanvas) Destroy() {
	// cleanup here
}
