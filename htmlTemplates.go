package main

var templates = map[string]string {

  "waveMark": "<!--\nThis Document is generated using Wave.\nWave: https://www.github.com/KILLinefficiency/Wave\n-->\n\n",

  "htmlTopBody": "<!DOCTYPE html>\n<html>\n\t<head>\n\t\t<title>%s</title>\n\t</head>\n",

  "htmlCSS": "\t<body style = 'background-color: %s; background-image: %s; color: %s; text-align: %s; margin: %spx; border-style: %s;'>\n",

  "css": "style = 'font-family: %s; color: %s; background-color: %s; font-size: %spx; text-align: %s; margin: %spx; border-style: %s; list-style-type: %s;'",

  "text": "\t\t<p %s>%s</p>\n",

  "link": "\t\t<a href = '%s' target = '_blank' %s>%s</a>\n",

  "mail": "\t\t<a href = 'mailto:%s' %s>%s</a>\n",

  "points": "\t\t\t<li>%s</li>\n",

  "pointsBody": "\t\t<%s %s>\n%s\t\t</%s>\n",

  "tableBorder": "style = 'border: 2px solid black;'",

  "tableValues": "\t\t\t\t\t<td %s>%s</td>\n",

  "tableBody": "\t\t\t\t<tr %s>\n%s\t\t\t\t</tr>\n",

  "tableComplete": "\t\t<div %s>\n\t\t\t<table %s>\n%s\t\t\t</table>\n\t\t</div>\n",

  "checkbox": "\t\t\t<input type = 'checkbox'>%s<br>\n",

  "checkboxBody": "\t\t<div %s>\n%s\t\t</div>\n",

  "quote": "\t\t<br><b><i>\"%s\"</b></i><br>\n",

  "image": "\t\t<div style = 'text-align: %s; margin: %spx; border-style: %s;'>\n\t\t\t<img width = '%s' height = '%s' src = %s>\n\t\t</div>\n",

  "htmlEnd": "\t</body>\n</html>",

}
