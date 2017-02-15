package layout

import (
	"bytes"
)

func Base(body string, title string) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n<!DOCTYPE html>\n<html>\n<head>")
	_buffer.WriteString((title))
	_buffer.WriteString("\n<meta content=\"width=device-width, initial-scale=1.0\" name=\"viewport\" />\n<link href=\"//netdna.bootstrapcdn.com/bootstrap/3.0.0-rc1/css/bootstrap.min.css\" rel=\"stylesheet\" />\n<link rel=\"stylesheet\" media=\"all\" href=\"/assets/application-b3afbeca3fa53a9a3c1491445fa92aac7c0ad0e20cb142c9ac056e72c723e656.css\" />\n<link href=\"/favicon.ico\" rel=\"shortcut icon\" />\n<link href=\"/manifest.json\" rel=\"manifest\" />\n<meta content=\"#563d7c\" name=\"theme-color\" />\n<link href=\"https://www.raindropgear.com/\" rel=\"canonical\" />\n</head>\n<body class=\"root\" id=\"pages\">\n<header class=\"navbar navbar-inverse navbar-fixed-top\">\n<div class=\"container\">\n<a class=\"navbar-brand\" href=\"/\">Raindrop Gear</a>\n<button class=\"navbar-toggle\" data-target=\".bs-navbar-collapse\" data-toggle=\"collapse\" type=\"button\">\n<span class=\"icon-bar\">\n</span>\n<span class=\"icon-bar\">\n</span>\n<span class=\"icon-bar\">\n</span>\n</button>\n<div class=\"nav-collapse collapse bs-navbar-collapse\">\n<ul class=\"nav navbar-nav\">\n<li class=\"active\">\n<a href=\"/\">Home</a>\n</li>\n<li>\n<a target=\"_blank\" href=\"/rdg-howto.pdf\">How to</a>\n</li>\n</ul>\n</div>\n</div>\n</header>")
	_buffer.WriteString((body))
	_buffer.WriteString("\n\n<div class=\"push\">\n</div>\n<footer class=\"navbar\">\n<p>&copy; LYPC 2013 All rights reserved.</p>\n</footer>\n</body>\n</html>")

	return _buffer.String()
}
