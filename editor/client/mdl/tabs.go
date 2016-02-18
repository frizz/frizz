package mdl

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"

	"honnef.co/go/js/dom"
)

type TabsStruct struct {
	*dom.HTMLDivElement
	header *tabsHeader
	panels []*dom.HTMLDivElement
}

func NewTabs() *TabsStruct {
	ts := &TabsStruct{HTMLDivElement: get("div").(*dom.HTMLDivElement)}
	ts.Class().SetString("mdl-tabs mdl-js-tabs mdl-js-ripple-effect")
	ts.header = &tabsHeader{HTMLDivElement: get("div").(*dom.HTMLDivElement)}
	ts.header.Class().SetString("mdl-tabs__tab-bar")
	ts.HTMLDivElement.AppendChild(ts.header)
	return ts
}

func (t *TabsStruct) AddTab(title string, content dom.Element) {
	index := len(t.header.HTMLDivElement.ChildNodes())
	id := fmt.Sprintf("panel%d", index)
	tab := get("a").(*dom.HTMLAnchorElement)
	tab.Href = fmt.Sprintf("#%s", id)
	tab.Class().Add("mdl-tabs__tab")
	if index == 0 {
		tab.Class().Add("is-active")
	}
	tab.SetTextContent(title)
	t.header.AppendChild(tab)

	content.Class().Add("mdl-tabs__panel")
	if index == 0 {
		content.Class().Add("is-active")
	}
	content.SetID(id)
	t.HTMLDivElement.AppendChild(content)
}

type tabsHeader struct {
	*dom.HTMLDivElement
}

func (t *TabsStruct) Upgrade() {
	js.Global.Get("componentHandler").Call("upgradeElement", t)
}

/**
<div class="mdl-tabs mdl-js-tabs mdl-js-ripple-effect">
  <div class="mdl-tabs__tab-bar">
      <a href="#panel1" class="mdl-tabs__tab is-active">Panel one</a>
      <a href="#panel2" class="mdl-tabs__tab">Panel two</a>
      <a href="#panel3" class="mdl-tabs__tab">Panel three</a>
  </div>

  <div class="mdl-tabs__panel is-active" id="panel1">
    panel1
  </div>
  <div class="mdl-tabs__panel" id="panel2">
    panel2
  </div>
  <div class="mdl-tabs__panel" id="panel3">
    panel3
  </div>
</div>
*/
