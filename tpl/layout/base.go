package layout

import (
	"bytes"
)

func Base(body string, title string) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n<!DOCTYPE html>\n<html>\n<head>\n  <meta charset=\"utf-8\">\n  <meta name=\"viewport\" content=\"width=device-width,minimum-scale=1,initial-scale=1\">\n  <meta content=\"IE=Edge\" http-equiv=\"X-UA-Compatible\">\n  <meta property=\"og:description\" content=\"We specialise in making belts to wear while you are raindropping, this allows for the ease of not having to move away from your client when you are changing oils.\">\n  <meta name=\"description\" content=\"We specialise in making belts to wear while you are raindropping, this allows for the ease of not having to move away from your client when you are changing oils.\">\n  ")
	_buffer.WriteString((title))
	_buffer.WriteString("\n  <link rel=\"shortcut icon\" href=\"/favicon.ico\">\n  <link rel=\"canonical\" href=\"http://www.raindropgear.com/\">\n  <link href=\"https://fonts.googleapis.com/css?family=Roboto:100,200,300,400,500,700\" rel=\"stylesheet\" type=\"text/css\">\n</head>\n<body>\n<amp-sidebar id='sidebar' side='right' layout='nodisplay'>\n  <button class=\"tab hamburger\" id=\"menu-button\" on='tap:sidebar.toggle'></button>\n  <nav>\n    <ul>\n      <li>\n        <ul>\n          <li><a href=\"/\">Raindrop Gear</a></li>\n          <li class=\"is-active\"><a href=\"/\">Home</a></li>\n          <li><a href=\"/rdg-howto.pdf\">How to</a></li>\n        </ul>\n      </li>\n    </ul>\n  </nav>\n</amp-sidebar>\n\n<header class=\"header fixed\">\n  <div class=\"wrap\">\n    <div class=\"container\">\n      <div class=\"nav-container\">\n         <div class=\"left-nav\">\n            <a href=\"/\" class=\"tab header-title\">\n              <span>Raindrop Gear</span>\n            </a>\n        </div>\n        <div class=\"right alt\">\n          <a class=\"tab desktop is-active\" href=\"/\">\n            Home\n          </a>\n          <a class=\"tab desktop\" href=\"/rdg-howto.pdf\">\n            How to\n          </a>\n          <button class=\"tab hamburger\" id=\"menu-button\" on='tap:sidebar.toggle'></button>\n        </div>\n      </div>\n    </div>\n  </div>\n</header>\n<div class=\"wrap\">\n  <div class=\"content\">\n    <div class=\"container bgw\">\n      ")
	_buffer.WriteString((body))
	_buffer.WriteString("\n    </div>\n    <footer style=\"margin-top:2em;\">\n      <div class=\"container\">\n        &copy; LYPC 2016 All rights reserved.\n      </div>\n    </footer>\n  </div>\n</div>\n</body>\n</html>")

	return _buffer.String()
}
