package hogosurutoaster

import (
	"strings"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/animationevent"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/htmldivelement"
	"github.com/realPy/hogosuru/htmllinkelement"
	"github.com/realPy/hogosuru/htmlstyleelement"
	"github.com/realPy/hogosuru/node"
	"github.com/realPy/hogosuru/promise"
)

const (
	ToasterError  = "error"
	ToasterWarn   = "warn"
	ToasterOk     = "ok"
	ToasterInfo   = "info"
	ToasterCustom = ""
)

//go:generate go run cmd/csscompact.go hogosurutoaster

type Toaster struct {
	parentNode node.Node
	container  htmldivelement.HtmlDivElement
}

func (t *Toaster) OnLoad(d document.Document, n node.Node, route string) (*promise.Promise, []hogosuru.Rendering) {
	var err error
	t.parentNode = n

	if t.container.Empty() {

		if t.container, err = htmldivelement.New(d); hogosuru.AssertErr(err) {
			t.container.SetID("hogosuru-toasters")
			if head, err := d.Head(); hogosuru.AssertErr(err) {
				if link, err := htmllinkelement.New(d); hogosuru.AssertErr(err) {

					link.SetRel("stylesheet")
					link.SetHref("https://fonts.googleapis.com/icon?family=Material+Icons")
					head.AppendChild(link.Node)

				}
				if style, err := htmlstyleelement.New(d); hogosuru.AssertErr(err) {

					style.SetTextContent(css)
					head.AppendChild(style.Node)

				}

			}

		}
	}

	return nil, nil
}

func (t Toaster) message(messageToaster string, fontColor, backgroundColor, borderColor, materialDesignIcon, materialDesignIconColor, toasterClass string) {
	if d, err := document.New(); hogosuru.AssertErr(err) {

		if div, err := htmldivelement.New(d); hogosuru.AssertErr(err) {

			var classToaster strings.Builder

			if i, err := d.CreateHTMLElement("i"); hogosuru.AssertErr(err) {
				i.SetClassName("material-icons , hogosurutoaster-icon")
				i.Style_().SetProperty("color", materialDesignIconColor)
				i.SetTextContent(materialDesignIcon)
				div.AppendChild(i.Node)

			}

			if divtext, err := htmldivelement.New(d); hogosuru.AssertErr(err) {
				divtext.SetClassName("hogosurutoaster-text")
				divtext.Style_().SetProperty("color", fontColor)
				div.AppendChild(divtext.Node)
				divtext.SetTextContent(messageToaster)
			}

			classToaster.WriteString("hogosurutoaster hogosurutoastershow")

			if toasterClass != "" {
				classToaster.WriteString(" hogosurutoaster-")
				classToaster.WriteString(toasterClass)
			} else {

				div.Style_().SetProperty("color", borderColor)
				div.Style_().SetProperty("background-color", backgroundColor)
			}

			div.SetClassName(classToaster.String())
			t.container.AppendChild(div.Node)

			div.OnAnimationEnd(func(a animationevent.AnimationEvent) {

				if name, err := a.AnimationName(); hogosuru.AssertErr(err) && name == "fadeout" {
					t.container.RemoveChild(div.Node)
				}

			})
		}
	}

}

func (t Toaster) AddMessage(messageToaster string, typeToaster string) {

	switch typeToaster {

	case ToasterError:
		t.message(messageToaster, "black", "", "", "error", "red", typeToaster)
	case ToasterWarn:
		t.message(messageToaster, "black", "", "", "warning", "orange", typeToaster)
	case ToasterOk:
		t.message(messageToaster, "black", "", "", "done", "green", typeToaster)
	case ToasterInfo:
		t.message(messageToaster, "black", "", "", "info", "blue", typeToaster)
	}
}

func (t Toaster) CustomMessage(messageToaster string, fontColor, backgroundColor, borderColor, materialDesignIcon, materialDesignIconColor string) {

	t.message(messageToaster, fontColor, backgroundColor, borderColor, materialDesignIcon, materialDesignIconColor, "")
}

func (t Toaster) SetText(message string) {
	t.container.SetTextContent(message)
}

func (t *Toaster) OnEndChildRendering(r hogosuru.Rendering) {

}

func (t *Toaster) OnEndChildsRendering(tree node.Node) {
	t.parentNode.AppendChild(t.container.Node)
}

func (t *Toaster) Node() node.Node {

	return t.container.Node
}

func (t *Toaster) OnUnload() {

	p, _ := t.container.ParentNode()

	p.RemoveChild(t.container.Node)

}
