package main

var (

  waveMark string = "<!--\nThis Document is generated using Wave.\nWave: https://www.github.com/KILLinefficiency/Wave\n-->\n\n"

  htmlTopBodyTemplate string = "<!DOCTYPE html>\n<html>\n\t<head>\n\t\t<title>%s</title>\n\t</head>\n"

  htmlCSSTemplate string = "\t<body style = 'background-color: %s; background-image: %s; color: %s; text-align: %s; margin: %spx; border-style: %s;'>\n"

  cssTemplate string = "style = 'font-family: %s; color: %s; background-color: %s; font-size: %spx; text-align: %s; margin: %spx; border-style: %s; list-style-type: %s;'"

  textTemplate string = "\t\t<p %s>%s</p>\n"

  linkTemplate string = "\t\t<a href = '%s' %s>%s</a>\n"

  mailTemplate string = "\t\t<a href = 'mailto:%s' %s>%s</a>\n"

  pointsTemplate string = "\t\t\t<li>%s</li>\n"

  pointsBodyTemplate string = "\t\t<%s %s>\n%s\t\t</%s>\n"

  tableBorder string = "style = 'border: 2px solid black;'"

  tableValuesTemplate string = "\t\t\t\t\t<td %s>%s</td>\n"

  tableBodyTemplate string = "\t\t\t\t<tr %s>\n%s\t\t\t\t</tr>\n"

  tableCompleteTemplate string = "\t\t<div %s>\n\t\t\t<table %s>\n%s\t\t\t</table>\n\t\t</div>\n"

  checkboxTemplate string = "\t\t\t<input type = 'checkbox'>%s<br>\n"

  checkboxBodyTemplate string = "\t\t<div %s>\n%s\t\t</div>\n"

  quoteTemplate string = "\t\t<br><b><i>\"%s\"</b></i><br>\n"

  imageTemplate string = "\t\t<div style = 'text-align: %s; margin: %spx; border-style: %s;'>\n\t\t\t<img width = '%s' height = '%s' src = %s>\n\t\t</div>\n"

  htmlEnd string = "\t</body>\n</html>"

)
