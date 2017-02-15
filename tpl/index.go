package tpl

import (
	"bytes"
	"github.com/dgodd/raindropgear_go/model"
	"github.com/dgodd/raindropgear_go/tpl/layout"
	"github.com/sipin/gorazor/gorazor"
)

func Index(belts []model.Belt) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n\n<amp-img src=\"/assets/beltwithoils-c2029d88912ee1e48fe05b514c56771f70abbb244b88aa12214eaa3251651dd2.webp\" width=\"1200\" height=\"450\" layout=\"responsive\" class=\"doc_header_bg i-amphtml-element i-amphtml-layout-responsive i-amphtml-layout-size-defined\"><i-amphtml-sizer style=\"display: block; padding-top: 37.5%;\"></i-amphtml-sizer><div class=\"i-amphtml-loading-container i-amphtml-fill-content amp-hidden\"><div class=\"i-amphtml-loader\"><div class=\"i-amphtml-loader-dot\"></div><div class=\"i-amphtml-loader-dot\"></div><div class=\"i-amphtml-loader-dot\"></div></div></div></amp-img>\n<h1>Welcome to Raindrop Gear.</h1>\n<p>We specialise in making belts to wear while you are raindropping, this allows for the ease of not having to move away from your client when you are changing oils.</p>\n<p>The belts are created to hold 15ml size bottles. They are 14.5 inches wide and 5.5 inches long. They are made with durable fabric and high grade elastic that will withstand constant use, and is not damaged by the oils breaking down the elastic.  The  strap is  made with adjustable belting. They are also hand washable and drip dry.</p>\n\n<div class=\"items\">")
	for _, b := range belts {

		_buffer.WriteString("<a class=\"tile item\" href=\"/belts/")
		_buffer.WriteString(gorazor.HTMLEscape(b.Slug))
		_buffer.WriteString("\">\n<img src=\"")
		_buffer.WriteString(gorazor.HTMLEscape(b.Front))
		_buffer.WriteString("/convert?w=300&amp;h=300&amp;fit=crop\" width=\"300\" height=\"300\">\n<p class=\"name\">")
		_buffer.WriteString(gorazor.HTMLEscape(b.Title))
		_buffer.WriteString("</p>\n<p class=\"price\">$55.00</p>\n</a>")

	}
	_buffer.WriteString("\n</div>")

	title := func() string {
		var _buffer bytes.Buffer

		_buffer.WriteString("<title>Raindropgear</title>")

		return _buffer.String()
	}

	return layout.Base(_buffer.String(), title())
}
