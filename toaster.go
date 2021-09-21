package hogosurutoaster

import (
	"strings"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/animationevent"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/htmldivelement"
	"github.com/realPy/hogosuru/htmlstyleelement"
	"github.com/realPy/hogosuru/node"
	"github.com/realPy/hogosuru/promise"
)

const (
	ToasterError = "error"
	ToasterWarn  = "warn"
	ToasterOk    = "ok"
	ToasterInfo  = "info"
	ToasterDefault
)

var toasteranimate string = `

.hogosurutoastershow {
	visibility: visible;
	-webkit-animation: slide 0.5s, fadeout 0.5s 3.5s;
	animation: slide 0.5s, fadeout 0.5s 3.5s;
  }
  

  @-webkit-keyframes slide {
	from {right: -300px;} 
	to {right: 0px;}
  }

  @keyframes slide {
	from {right: -300px;} 
	to {right: 0px;}
  }
  @keyframes fadeout {
	from { opacity: 1;}
	to { opacity: 0;}
  }

  @-webkit-keyframes fadeout {
	from {opacity: 1;} 
	to {opacity: 0;}
  }
  
  @keyframes fadeout {
	from { opacity: 1;}
	to { opacity: 0;}
  }

.hogosurutoaster {
	min-width: 250px;
	border: solid 1px;
	text-align: center;
	box-shadow: 2px 2px 4px gray;
	border-radius: 4px;
	padding: 20px;
	font-size: 15px;
	margin: 2px;
	position: relative;
	right: -300px,
}

.hogosurutoaster-error {
	  background-color: #ffaeae;
	  color: #ff0707;
  }
  
.hogosurutoaster-warn {
	background-color: #ffb100;
	color: #fff;
}

.hogosurutoaster-ok {
	background-color: #ffb100;
	color: #fff;
}

.hogosurutoaster-info {
    background-color: #d2daff;
    color: #2317c5;
}

.hogosurutoaster-ok {
    background-color: #dfffe7;
    color: #127d3e;
}


.hogosurutoaster-warn {
    background-color: #ffe6ae;
    color: #e68e1f;
}`

type Toaster struct {
	parentNode      node.Node
	container       htmldivelement.HtmlDivElement
	BackgroundColor string
}

func (t *Toaster) OnLoad(d document.Document, n node.Node, route string) (*promise.Promise, []hogosuru.Rendering) {
	var err error
	t.parentNode = n

	if t.container.Empty() {
		if t.container, err = htmldivelement.New(d); hogosuru.AssertErr(err) {
			t.container.SetID("hogosuru-toasters")
			t.container.Style_().SetProperty("position", "fixed")
			t.container.Style_().SetProperty("right", "0px")
			t.container.Style_().SetProperty("top", "0px")
			t.container.Style_().SetProperty("display", "flex")
			t.container.Style_().SetProperty("z-index", "1")
			t.container.Style_().SetProperty("flex-direction", "column-reverse")
			t.container.Style_().SetProperty("padding", "4px")

			if style, err := htmlstyleelement.New(d); hogosuru.AssertErr(err) {

				if head, err := d.Head(); hogosuru.AssertErr(err) {
					style.SetTextContent(toasteranimate)
					head.AppendChild(style.Node)

				}

			}
		}
	}

	return nil, nil
}

func (t Toaster) AddMessage(messageToaster string, typeToaster string) {

	if d, err := document.New(); hogosuru.AssertErr(err) {

		if div, err := htmldivelement.New(d); hogosuru.AssertErr(err) {

			var classToaster strings.Builder

			switch typeToaster {
			case ToasterError, ToasterWarn, ToasterOk, ToasterInfo:
				classToaster.WriteString("hogosurutoaster hogosurutoastershow hogosurutoaster-")
				classToaster.WriteString(typeToaster)
			default:
				classToaster.WriteString("hogosurutoaster hogosurutoasters")

			}

			div.SetTextContent(messageToaster)

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
