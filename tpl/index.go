package tpl

import (
	"bytes"
	"github.com/dgodd/raindropgear_go/model"
	"github.com/dgodd/raindropgear_go/tpl/layout"
	"github.com/sipin/gorazor/gorazor"
)

func Index(belts []model.Belt) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n\n<div class=\"container\">\n<div class=\"row\">\n<img class=\"img-responsive\" src=\"https://raindropgear.com/assets/beltwithoils-c2029d88912ee1e48fe05b514c56771f70abbb244b88aa12214eaa3251651dd2.webp\" />\n</div>\n<div class=\"hero-unit\">\n<h1>Welcome to Raindrop Gear.</h1>\n<br />\n<p>We specialise in making belts to wear while you are raindropping, this allows for the ease of not having to move away from your client when you are changing oils.</p>\n<p>The belts are created to hold 15ml size bottles. They are 14.5 inches wide and 5.5 inches long. They are made with durable fabric and high grade elastic that will withstand constant use, and is not damaged by the oils breaking down the elastic.  The  strap is  made with adjustable belting. They are also hand washable and drip dry.</p>\n</div>\n\n\n\n<div class=\"items\">")
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

		_buffer.WriteString("<title>RaindropGear</title>")

		return _buffer.String()
	}
	_buffer.WriteString("\n\n<div class=\"row\">\n<div class=\"col-3\">\n<a href=\"/belts/silver\">\n<img class=\"img-responsive\" src=\"https://www.filepicker.io/api/file/vVbUNSwOQL21B2ESGWx5/convert?w=300&amp;h=300&amp;fit=crop\" />\n</a>\n<a href=\"/belts/silver\">Silver</a>\n<div class=\"price\">$55.00</div>\n</div>\n<div class=\"col-3\">\n<a href=\"/belts/purple\">\n<img class=\"img-responsive\" src=\"https://www.filepicker.io/api/file/fQicCpgQNuA2bpkLp2Td/convert?w=300&amp;h=300&amp;fit=crop\" />\n</a>\n<a href=\"/belts/purple\">Purple</a>\n<div class=\"price\">$55.00</div>\n</div>\n<div class=\"col-3\">\n<a href=\"/belts/red-belt\">\n<img class=\"img-responsive\" src=\"https://www.filepicker.io/api/file/bATTEKoqSy2cRgIEWdb2/convert?w=300&amp;h=300&amp;fit=crop\" />\n</a>\n<a href=\"/belts/red-belt\">Red Belt </a>\n<div class=\"price\">$55.00</div>\n</div>\n<div class=\"col-3\">\n<a href=\"/belts/yellow-filigree\">\n<img class=\"img-responsive\" src=\"https://www.filepicker.io/api/file/FkQTF7TDSoWshVItFKNz/convert?w=300&amp;h=300&amp;fit=crop\" />\n</a>\n<a href=\"/belts/yellow-filigree\">Yellow Filigree</a>\n<div class=\"price\">$55.00</div>\n</div>\n<div class=\"col-3\">\n<a href=\"/belts/sage-belt\">\n<img class=\"img-responsive\" src=\"https://www.filepicker.io/api/file/LshrJQCsRYGp0Jc59tSP/convert?w=300&amp;h=300&amp;fit=crop\" />\n</a>\n<a href=\"/belts/sage-belt\">Sage Belt</a>\n<div class=\"price\">$55.00</div>\n</div>\n<div class=\"col-3\">\n<a href=\"/belts/grape-mauve\">\n<img class=\"img-responsive\" src=\"https://www.filepicker.io/api/file/ezv59G44T6STfTPBF0Dq/convert?w=300&amp;h=300&amp;fit=crop\" />\n</a>\n<a href=\"/belts/grape-mauve\">Grape &amp; Mauve </a>\n<div class=\"price\">$55.00</div>\n</div>\n<div class=\"col-3\">\n<a href=\"/belts/metalic-mint\">\n<img class=\"img-responsive\" src=\"https://www.filepicker.io/api/file/UgscDUgySmO3OPKcuQ2G/convert?w=300&amp;h=300&amp;fit=crop\" />\n</a>\n<a href=\"/belts/metalic-mint\">Metalic Mint</a>\n<div class=\"price\">$55.00</div>\n</div>\n<div class=\"col-3\">\n<a href=\"/belts/black\">\n<img class=\"img-responsive\" src=\"https://www.filepicker.io/api/file/7dRZb9TkyQ5BPR2YsBcw/convert?w=300&amp;h=300&amp;fit=crop\" />\n</a>\n<a href=\"/belts/black\">Black</a>\n<div class=\"price\">$55.00</div>\n</div>\n<div class=\"col-3\">\n<a href=\"/belts/mauve\">\n<img class=\"img-responsive\" src=\"https://www.filepicker.io/api/file/P7nPJ9aRTK1i4lAQ45FA/convert?w=300&amp;h=300&amp;fit=crop\" />\n</a>\n<a href=\"/belts/mauve\">Mauve</a>\n<div class=\"price\">$55.00</div>\n</div>\n</div>\n</div>")

	return layout.Base(_buffer.String(), title())
}
