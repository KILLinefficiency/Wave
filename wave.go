package main

import (
  "os"
  "fmt"
  "strings"
  "strconv"
  "io/ioutil"
)

var (
  htmlBody string
  cssBody string
)

func main() {
  if len(os.Args) == 1 {
    fmt.Printf("No Wave Script passed.\nPass in a Wave Script as a command-line argument.\nLike:\n\twave <scriptName>\n")
    os.Exit(1)
  }

  var sourceName string = os.Args[1]

  byteStream, err := ioutil.ReadFile(sourceName)
  if err != nil {
    fmt.Printf("Invalid file address: %s\n", sourceName)
    os.Exit(2)
  }

  var script string = string(byteStream)
  scriptLines := strings.Split(strings.TrimSpace(script), "\n")
  for _, line := range scriptLines {
    tokens := strings.Split(strings.TrimSpace(line), " ")
    property := strings.Join(tokens[1:], " ")

    switch tokens[0] {
      case "~title":
        pageDefaults["pTitle"] = property
      case "~bg":
        pageDefaults["pBGcolor"] = property
        contentDefaults["cBGcolor"] = pageDefaults["pBGcolor"]
      case "~pic":
        pageDefaults["pBGimage"] = property
      case "~align":
        pageDefaults["pAlign"] = property
        contentDefaults["cAlign"] = pageDefaults["pAlign"]
      case "~box":
        pageDefaults["pBox"] = property
      case "~box-style":
        pageDefaults["pBoxStyle"] = property
    }

    switch tokens[0] {
      case "!tab":
        contentDefaults["cTab"] = property
      case "!font":
        contentDefaults["cFont"] = property
      case "!size":
        contentDefaults["cSize"] = property
      case "!color":
        contentDefaults["cColor"] = property
      case "!box":
        contentDefaults["cBox"] = property
      case "!box-style":
        contentDefaults["cBoxStyle"] = property
      case "!align":
        contentDefaults["cAlign"] = property
      case "!bg":
        contentDefaults["cBGcolor"] = property
      case "!points-type":
        if property == "ordered" {
          contentDefaults["cPointsType"] = "ol"
        }
        if property == "unordered" {
          contentDefaults["cPointsType"] = "ul"
        }
      case "!points-style":
        contentDefaults["cPointsStyle"] = property
      case "!dim":
        widthHeight := strings.Split(property, "x")
        contentDefaults["cWidth"] = strings.TrimSpace(widthHeight[0])
        contentDefaults["cHeight"] = strings.TrimSpace(widthHeight[1])
      case "!sep":
        contentDefaults["cDelimiter"] = property
      case "!colsep":
        contentDefaults["cTableDelimiter"] = property
      case "!default":
        fmt.Println("will do this later brrr...")
    }

    cssBody = fmt.Sprintf("style = 'font-family: %s; color: %s; background-color: %s; font-size: %spx; text-align: %s; margin: %spx; border-style: %s; list-style-type: %s;'", contentDefaults["cFont"], contentDefaults["cColor"], contentDefaults["cBGcolor"], contentDefaults["cSize"], contentDefaults["cAlign"], contentDefaults["cBox"], contentDefaults["cBoxStyle"], contentDefaults["cPointsStyle"])

    switch tokens[0] {
      case "$text":
        htmlBody += fmt.Sprintf("\t\t<p %s>%s</p>\n", cssBody, property)

      case "$file":
        textFile, _ := ioutil.ReadFile(property)
        tabNumber, _ := strconv.Atoi(contentDefaults["cTab"])
        var fileStr string = string(textFile)
        fileStr = strings.Replace(fileStr, "\n", "<br>", -1)
        fileStr = strings.Replace(fileStr, " ", "&nbsp;", -1)
        fileStr = strings.Replace(fileStr, "\t", strMultiply("&nbsp;", tabNumber), -1)
        htmlBody += fmt.Sprintf("\t\t<p %s>%s</p>\n", cssBody, fileStr)

      case "$nl":
        if len(tokens) == 1 {
          property = "1"
        }
        times, _ := strconv.Atoi(property)
        htmlBody += fmt.Sprintf("\t\t%s\n", strMultiply("<br>", times))

      case "$link":
        linkTitle := strings.Split(property, contentDefaults["cDelimiter"])
        if len(linkTitle) > 1 {
          contentDefaults["cLink"] = strings.TrimSpace(linkTitle[0])
          contentDefaults["cLinkTitle"] = strings.TrimSpace(linkTitle[1])
        }
        htmlBody += fmt.Sprintf("\t\t<a href = '%s' %s>%s</a>\n", contentDefaults["cLink"], cssBody, contentDefaults["cLinkTitle"])

      case "$mail":
        mailTitle := strings.Split(property, contentDefaults["cDelimiter"])
        if len(mailTitle) > 1 {
          contentDefaults["cMailAddress"] = strings.TrimSpace(mailTitle[0])
          contentDefaults["cMailTitle"] = strings.TrimSpace(mailTitle[1])
        }
        htmlBody += fmt.Sprintf("\t\t<a href = 'mailto:%s' %s>%s</a>\n", contentDefaults["cMailAddress"], cssBody, contentDefaults["cMailTitle"])

      case "$points":
        listPoints := strings.Split(property, contentDefaults["cDelimiter"])
        var allPoints string
        for _, point := range listPoints {
          allPoints += fmt.Sprintf("\t\t\t<li>%s</li>\n", strings.TrimSpace(point))
        }
        var pointsBody string = fmt.Sprintf("\t\t<%s %s>\n%s\t\t</%s>\n", contentDefaults["cPointsType"], cssBody, allPoints, contentDefaults["cPointsType"])
        htmlBody += pointsBody

      case "$table":
        var tableBody string
        var tableBorder string = "style = 'border: 2px solid black;'"
        tableRows := strings.Split(property, contentDefaults["cTableDelimiter"])
        for _, rowValues := range tableRows {
          values := strings.Split(rowValues, contentDefaults["cDelimiter"])
          var rowBody string
          for _, addValues := range values {
            rowBody += fmt.Sprintf("\t\t\t\t\t<td %s>%s</td>\n", tableBorder, strings.TrimSpace(addValues))
          }
          tableBody += fmt.Sprintf("\t\t\t\t<tr %s>\n%s\t\t\t\t</tr>\n", tableBorder, rowBody)
        }
        htmlBody += fmt.Sprintf("\t\t<div %s>\n\t\t\t<table %s>\n%s\t\t\t</table>\n\t\t</div>\n", cssBody, tableBorder, tableBody)

      case "$check":
        checkPoints := strings.Split(property, contentDefaults["cDelimiter"])
        var checkPointsBody string
        for _, points := range checkPoints {
          checkPointsBody += fmt.Sprintf("\t\t\t<input type = 'checkbox'>%s<br>\n", strMultiply("&nbsp;", 2) + strings.TrimSpace(points))
        }
        htmlBody += fmt.Sprintf("\t\t<div %s>\n%s\t\t</div>\n", cssBody, checkPointsBody)

      case "$quote":
        htmlBody += fmt.Sprintf("\t\t<br><b><i>\"%s\"</b></i><br>\n", property)

      case "$pic":
        htmlBody += fmt.Sprintf("\t\t<div style = 'text-align: %s; margin: %spx; border-style: %s;'>\n\t\t\t<img width = '%s' height = '%s' src = %s>\n\t\t</div>\n", contentDefaults["cAlign"], contentDefaults["cBox"], contentDefaults["cBoxStyle"],contentDefaults["cWidth"], contentDefaults["cHeight"], property)

      case "$html":
        htmlBody += fmt.Sprintf("\t\t%s\n", property)
    }

  }

  var waveMark string = "\n<!--\nThis Document is generated using Wave.\nWave: https://www.github.com/KILLinefficiency/Wave\n-->\n\n"
  var htmlTopBody string = fmt.Sprintf("<!DOCTYPE html>\n<html>\n\t<head>\n\t\t<title>%s</title>\n\t</head>\n", pageDefaults["pTitle"])
  var htmlCSS string = fmt.Sprintf("\t<body style = 'background-color: %s; background-image: %s; text-align: %s; margin: %spx; border-style: %s;'>\n", pageDefaults["pBGcolor"], pageDefaults["pBGimage"], pageDefaults["pAlign"], pageDefaults["pBox"], pageDefaults["pBoxStyle"])
  var htmlComplete string = waveMark + htmlTopBody + htmlCSS + htmlBody + "\t</body>\n</html>\n"

  fileName := strings.Split(sourceName, ".")
  if len(fileName) == 1 {
    fileName = append(fileName, ".html")
  } else {
    fileName[len(fileName) - 1] = ".html"
  }

  htmlFileName := strings.Join(fileName, "")
  htmlFile, err := os.Create(htmlFileName)
  if err != nil {
    fmt.Printf("Unable to create file: %s\n\nSource Code for the Document:\n\n%s\n", htmlFileName, htmlComplete)
    os.Exit(3)
  }
  htmlFile.WriteString(htmlComplete)
  htmlFile.Close()
}
