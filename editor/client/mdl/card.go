package mdl

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

type CardStruct struct {
	*dom.HTMLDivElement
	Content *dom.HTMLDivElement
}

func Card(titleText string, buttonText string) *CardStruct {
	c := &CardStruct{}
	c.HTMLDivElement = get("div").(*dom.HTMLDivElement)
	c.Class().SetString("demo-card-wide mdl-card mdl-shadow--2dp")
	c.Style().Set("width", "512px")
	c.Style().Set("margin", "auto")
	c.Style().Set("top", "10%")

	//<div class="mdl-card__title">
	titleDiv := get("div").(*dom.HTMLDivElement)
	titleDiv.Class().SetString("mdl-card__title")

	//<h2 class="mdl-card__title-text">Welcome</h2>
	title := get("h2").(*dom.HTMLHeadingElement)
	title.Class().SetString("mdl-card__title-text")
	title.SetTextContent(titleText)

	titleDiv.AppendChild(title)

	//<div class="mdl-card__supporting-text">
	c.Content = get("div").(*dom.HTMLDivElement)
	c.Content.Class().SetString("mdl-card__supporting-text")

	//<div class="mdl-card__actions mdl-card--border">
	buttonDiv := get("div").(*dom.HTMLDivElement)
	buttonDiv.Class().SetString("mdl-card__actions mdl-card--border")

	//<a class="mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect">
	button := get("a").(*dom.HTMLAnchorElement)
	button.Class().SetString("mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect")
	button.SetTextContent(buttonText)

	buttonDiv.AppendChild(button)

	c.AppendChild(titleDiv)
	c.AppendChild(c.Content)
	c.AppendChild(buttonDiv)

	dom.GetWindow().Document().GetElementsByTagName("body")[0].AppendChild(c)

	js.Global.Get("componentHandler").Call("upgradeElement", c)

	return c
}

/**
<!-- Wide card with share menu button -->
<style>
.demo-card-wide.mdl-card {
  width: 512px;
}
.demo-card-wide > .mdl-card__title {
  color: #fff;
  height: 176px;
  background: url('../assets/demos/welcome_card.jpg') center / cover;
}
.demo-card-wide > .mdl-card__menu {
  color: #fff;
}
</style>

<div class="demo-card-wide mdl-card mdl-shadow--2dp">
  <div class="mdl-card__title">
    <h2 class="mdl-card__title-text">Welcome</h2>
  </div>
  <div class="mdl-card__supporting-text">
    Lorem ipsum dolor sit amet, consectetur adipiscing elit.
    Mauris sagittis pellentesque lacus eleifend lacinia...
  </div>
  <div class="mdl-card__actions mdl-card--border">
    <a class="mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect">
      Get Started
    </a>
  </div>
  <div class="mdl-card__menu">
    <button class="mdl-button mdl-button--icon mdl-js-button mdl-js-ripple-effect">
      <i class="material-icons">share</i>
    </button>
  </div>
</div>/
*/
