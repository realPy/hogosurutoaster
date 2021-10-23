package main

import (
	"github.com/realPy/hogosurutoaster"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/htmlbuttonelement"
	"github.com/realPy/hogosuru/htmldivelement"
	"github.com/realPy/hogosuru/htmlinputelement"
	"github.com/realPy/hogosuru/htmloptionelement"
	"github.com/realPy/hogosuru/htmlselectelement"
	"github.com/realPy/hogosuru/node"
	"github.com/realPy/hogosuru/promise"
)

type GlobalContainer struct {
	parentNode node.Node
	node       node.Node
	toaster    hogosurutoaster.Toaster
}

func (w *GlobalContainer) OnLoad(d document.Document, n node.Node, route string) (*promise.Promise, []hogosuru.Rendering) {

	w.parentNode = n

	if global, err := htmldivelement.New(d); err == nil {

		global.SetID("global-container")
		w.node = global.Node
		testButton, _ := htmlbuttonelement.New(d)
		testButton.SetTextContent("Toast!")
		input, _ := htmlinputelement.New(d)
		input.SetType("text")

		s, _ := htmlselectelement.New(d)
		s.SetID("Select")

		o, _ := htmloptionelement.New(d)
		o.SetTextContent("Error")
		o.SetValue("error")
		s.AppendChild(o.Node)

		o, _ = htmloptionelement.New(d)
		o.SetTextContent("Warn")
		o.SetValue("warn")
		s.AppendChild(o.Node)

		o, _ = htmloptionelement.New(d)
		o.SetTextContent("Ok")
		o.SetValue("ok")
		s.AppendChild(o.Node)

		o, _ = htmloptionelement.New(d)
		o.SetTextContent("Info")
		o.SetValue("info")
		s.AppendChild(o.Node)

		o, _ = htmloptionelement.New(d)
		o.SetTextContent("custom")
		o.SetValue("home")
		s.AppendChild(o.Node)

		input.SetType("text")
		w.toaster.ToasterPosition = hogosurutoaster.FROMTOPRIGHT
		testButton.OnClick(func(e event.Event) {

			iValue, _ := input.Value()
			sValue, _ := s.Value()

			if sValue == "home" {
				w.toaster.CustomMessage(iValue, "green", "white", "red", "home", "blue")

			} else {
				w.toaster.AddMessage(iValue, sValue)
			}

		})

		global.AppendChild(testButton.Node)

		global.AppendChild(input.Node)
		global.AppendChild(s.Node)

	}

	//if no promise return we dont wait all childs to append

	return nil, []hogosuru.Rendering{&w.toaster}
}

func (w *GlobalContainer) Node(r hogosuru.Rendering) node.Node {

	return w.node
}

func (w *GlobalContainer) OnEndChildRendering(r hogosuru.Rendering) {

}

func (w *GlobalContainer) OnEndChildsRendering() {

	w.parentNode.AppendChild(w.node)
}

func (w *GlobalContainer) OnUnload() {

}

func main() {
	hogosuru.Init()
	hogosuru.Router().DefaultRendering(&GlobalContainer{})
	hogosuru.Router().Start(hogosuru.HASHROUTE)
	ch := make(chan struct{})
	<-ch

}
