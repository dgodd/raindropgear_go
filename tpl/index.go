package tpl

import (
	"bytes"
	"github.com/dgodd/raindropgear_go/model"
	"github.com/dgodd/raindropgear_go/tpl/layout"
	"github.com/sipin/gorazor/gorazor"
)

func Index(belts []model.Belt) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n\n<ul>")
	for _, b := range belts {

		_buffer.WriteString("<li>")
		_buffer.WriteString(gorazor.HTMLEscape(b.Title))
		_buffer.WriteString("</li>")

	}
	_buffer.WriteString("\n</ul>")

	title := func() string {
		var _buffer bytes.Buffer

		_buffer.WriteString("<title>Raindropgear</title>")

		return _buffer.String()
	}

	return layout.Base(_buffer.String(), title())
}
