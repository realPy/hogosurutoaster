package hogosurutoaster

import (
	"strings"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/animationevent"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/customevent"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/htmldivelement"
	"github.com/realPy/hogosuru/htmllinkelement"
	"github.com/realPy/hogosuru/htmlstyleelement"
	"github.com/realPy/hogosuru/node"
	"github.com/realPy/hogosuru/object"
	"github.com/realPy/hogosuru/promise"
)

const (
	ToasterError  = "error"
	ToasterWarn   = "warn"
	ToasterOk     = "ok"
	ToasterInfo   = "info"
	ToasterCustom = ""
)

const (
	FROMTOPRIGHT = iota
	FROMTOPLEFT
	FROMBOTTOMRIGHT
	FROMBOTTOMLEFT
)

//go:generate go run cmd/csscompact.go hogosurutoaster

type Toaster struct {
	parentNode      node.Node
	container       htmldivelement.HtmlDivElement
	ToasterPosition int
}

func (t *Toaster) OnLoad(d document.Document, n node.Node, route string) (*promise.Promise, []hogosuru.Rendering) {
	var err error
	t.parentNode = n

	if t.container.Empty() {

		customevent.GetInterface()
		object.GetInterface()

		if t.container, err = htmldivelement.New(d); hogosuru.AssertErr(err) {
			t.container.SetID("hogosuru-toasters")

			switch t.ToasterPosition {
			case FROMTOPRIGHT:
				t.container.SetClassName("hogosuru-toasters-top-right")
			case FROMTOPLEFT:
				t.container.SetClassName("hogosuru-toasters-top-left")
			case FROMBOTTOMRIGHT:
				t.container.SetClassName("hogosuru-toasters-bottom-right")
			case FROMBOTTOMLEFT:
				t.container.SetClassName("hogosuru-toasters-bottom-left")
			default:
				t.container.SetClassName("hogosuru-toasters-top-right")
			}

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

				if d, err := document.New(); hogosuru.AssertErr(err) {

					d.AddEventListener("hogosurutoaster-notify", func(e event.Event) {

						if obj, err := baseobject.Discover(e.JSObject()); err == nil {
							if c, ok := obj.(customevent.CustomEventFrom); ok {
								if detail, err := c.CustomEvent().Detail(); hogosuru.AssertErr(err) {
									if objdetail, ok := detail.(object.Object); ok {

										if mapObject, err := objdetail.Map(); hogosuru.AssertErr(err) {

											if mapObject.Has_("type") {
												t.AddMessage(mapObject.Get_("message").(string), mapObject.Get_("type").(string))
											}
										}

									}

								}

							}
						}

					})
					d.AddEventListener("hogosurutoaster-customnotify", func(e event.Event) {

						if obj, err := baseobject.Discover(e.JSObject()); err == nil {
							if c, ok := obj.(customevent.CustomEventFrom); ok {
								if detail, err := c.CustomEvent().Detail(); hogosuru.AssertErr(err) {
									if objdetail, ok := detail.(object.Object); ok {

										if mapObject, err := objdetail.Map(); hogosuru.AssertErr(err) {

											if mapObject.Has_("fontColor") && mapObject.Has_("backgroundColor") && mapObject.Has_("borderColor") && mapObject.Has_("materialDesignIcon") && mapObject.Has_("materialDesignIconColor") {
												t.CustomMessage(mapObject.Get_("message").(string), mapObject.Get_("fontColor").(string), mapObject.Get_("backgroundColor").(string), mapObject.Get_("borderColor").(string), mapObject.Get_("materialDesignIcon").(string), mapObject.Get_("materialDesignIconColor").(string))
											}
										}

									}

								}

							}
						}

					})

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

			switch t.ToasterPosition {
			case FROMTOPRIGHT, FROMBOTTOMRIGHT:
				classToaster.WriteString("hogosurutoaster hogosurutoastershow-right")
			case FROMTOPLEFT, FROMBOTTOMLEFT:
				classToaster.WriteString("hogosurutoaster hogosurutoastershow-left")
			default:
				classToaster.WriteString("hogosurutoaster hogosurutoastershow-right")
			}

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

func (t *Toaster) OnEndChildsRendering() {
	t.parentNode.AppendChild(t.container.Node)

}

func (t *Toaster) Node(r hogosuru.Rendering) node.Node {

	return t.container.Node
}

func (t *Toaster) OnUnload() {

	p, _ := t.container.ParentNode()

	p.RemoveChild(t.container.Node)

}
