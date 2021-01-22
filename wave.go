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

var (
  pTitle string = "Wave Document"
  pBGcolor string = "white"
  pBGimage string = "none"
  pAlign string = "left"
  pBox string = "0"
  pBoxStyle string = "hidden"
)

var (
  cTab int = 4
  cFont string = "Arial"
  cSize string = "17"
  cColor string = "black"
  cBox string = "0"
  cBoxStyle string = "hidden"
  cAlign string
  cBGcolor string
  cPointsType string = "ul"
  cPointsStyle string = "disc"
  cWidth string = "none"
  cHeight string = "none"
  cDelimiter string = ";"
  cTableDelimiter = "|"
  cLink string = "https://www.github.com/KILLinefficiency/Wave"
  cMailTitle, cMailAddress string
  cLinkTitle string = "Wave"
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
        pTitle = property
      case "~bg":
        pBGcolor = property
        cBGcolor = pBGcolor
      case "~pic":
        pBGimage = property
      case "~align":
        pAlign = property
        cAlign = pAlign
      case "~box":
        pBox = property
      case "~box-style":
        pBoxStyle = property
    }

    switch tokens[0] {
      case "!tab":
        tabNumber, _ := strconv.Atoi(property)
        cTab = tabNumber
      case "!font":
        cFont = property
      case "!size":
        cSize = property
      case "!color":
        cColor = property
      case "!box":
        cBox = property
      case "!box-style":
        cBoxStyle = property
      case "!align":
        cAlign = property
      case "!bg":
        cBGcolor = property
      case "!points-type":
        if property == "ordered" {
          cPointsType = "ol"
        }
        if property == "unordered" {
          cPointsType = "ul"
        }
      case "!points-style":
        cPointsStyle = property
      case "!dim":
        widthHeight := strings.Split(property, "x")
        cWidth = strings.TrimSpace(widthHeight[0])
        cHeight = strings.TrimSpace(widthHeight[1])
      case "!sep":
        cDelimiter = property
      case "!colsep":
        cTableDelimiter = property
      case "!default":
        cTab, cFont = 4, "Arial"
        cSize, cColor = "17", "black"
        cBox, cBoxStyle = "0", "hidden"
        cBoxStyle = "hidden"
        cAlign, cBGcolor = "", ""
        cPointsType, cPointsStyle = "ul", "disc"
        cWidth, cHeight = "none", "none"
        cLink, cLinkTitle = "https://www.github.com/KILLinefficiency/Wave", "Wave"
        cMailAddress, cMailTitle = "", ""
        cDelimiter, cTableDelimiter = ";", "|"
    }

    cssBody = fmt.Sprintf("style = 'font-family: %s; color: %s; background-color: %s; font-size: %spx; text-align: %s; margin: %spx; border-style: %s; list-style-type: %s;'", cFont, cColor, cBGcolor, cSize, cAlign, cBox, cBoxStyle, cPointsStyle)

    switch tokens[0] {
      case "$text":
        htmlBody += fmt.Sprintf("\t\t<p %s>%s</p>\n", cssBody, property)

      case "$file":
        textFile, _ := ioutil.ReadFile(property)
        var fileStr string = string(textFile)
        fileStr = strings.Replace(fileStr, "\n", "<br>", -1)
        fileStr = strings.Replace(fileStr, " ", "&nbsp;", -1)
        fileStr = strings.Replace(fileStr, "\t", strMultiply("&nbsp;", cTab), -1)
        htmlBody += fmt.Sprintf("\t\t<p %s>%s</p>\n", cssBody, fileStr)

      case "$nl":
        if len(tokens) == 1 {
          property = "1"
        }
        times, _ := strconv.Atoi(property)
        htmlBody += fmt.Sprintf("\t\t%s\n", strMultiply("<br>", times))

      case "$link":
        linkTitle := strings.Split(property, cDelimiter)
        if len(linkTitle) > 1 {
          cLink = strings.TrimSpace(linkTitle[0])
          cLinkTitle = strings.TrimSpace(linkTitle[1])
        }
        htmlBody += fmt.Sprintf("\t\t<a href = '%s' %s>%s</a>\n", cLink, cssBody, cLinkTitle)

      case "$mail":
        mailTitle := strings.Split(property, cDelimiter)
        if len(mailTitle) > 1 {
          cMailAddress = strings.TrimSpace(mailTitle[0])
          cMailTitle = strings.TrimSpace(mailTitle[1])
        }
        htmlBody += fmt.Sprintf("\t\t<a href = 'mailto:%s' %s>%s</a>\n", cMailAddress, cssBody, cMailTitle)

      case "$points":
        listPoints := strings.Split(property, cDelimiter)
        var allPoints string
        for _, point := range listPoints {
          allPoints += fmt.Sprintf("\t\t\t<li>%s</li>\n", strings.TrimSpace(point))
        }
        var pointsBody string = fmt.Sprintf("\t\t<%s %s>\n%s\t\t</%s>\n", cPointsType, cssBody, allPoints, cPointsType)
        htmlBody += pointsBody

      case "$table":
        var tableBody string
        var tableBorder string = "style = 'border: 2px solid black;'"
        tableRows := strings.Split(property, cTableDelimiter)
        for _, rowValues := range tableRows {
          values := strings.Split(rowValues, cDelimiter)
          var rowBody string
          for _, addValues := range values {
            rowBody += fmt.Sprintf("\t\t\t\t\t<td %s>%s</td>\n", tableBorder, strings.TrimSpace(addValues))
          }
          tableBody += fmt.Sprintf("\t\t\t\t<tr %s>\n%s\t\t\t\t</tr>\n", tableBorder, rowBody)
        }
        htmlBody += fmt.Sprintf("\t\t<div %s>\n\t\t\t<table %s>\n%s\t\t\t</table>\n\t\t</div>\n", cssBody, tableBorder, tableBody)

      case "$check":
        checkPoints := strings.Split(property, cDelimiter)
        var checkPointsBody string
        for _, points := range checkPoints {
          checkPointsBody += fmt.Sprintf("\t\t\t<input type = 'checkbox'>%s<br>\n", strMultiply("&nbsp;", 2) + strings.TrimSpace(points))
        }
        htmlBody += fmt.Sprintf("\t\t<div %s>\n%s\t\t</div>\n", cssBody, checkPointsBody)

      case "$quote":
        htmlBody += fmt.Sprintf("\t\t<br><b><i>\"%s\"</b></i><br>\n", property)

      case "$pic":
        htmlBody += fmt.Sprintf("\t\t<div style = 'text-align: %s; margin: %spx; border-style: %s;'>\n\t\t\t<img width = '%s' height = '%s' src = %s>\n\t\t</div>\n", cAlign, cBox, cBoxStyle,cWidth, cHeight, property)

      case "$html":
        htmlBody += fmt.Sprintf("\t\t%s\n", property)
    }

  }

  var waveMark string = "\n<!--\nThis Document is generated using Wave.\nWave: https://www.github.com/KILLinefficiency/Wave\n-->\n\n"
  var htmlTopBody string = fmt.Sprintf("<!DOCTYPE html>\n<html>\n\t<head>\n\t\t<title>%s</title>\n\t</head>\n", pTitle)
  var htmlCSS string = fmt.Sprintf("\t<body style = 'background-color: %s; background-image: %s; text-align: %s; margin: %spx; border-style: %s;'>\n", pBGcolor, pBGimage, pAlign, pBox, pBoxStyle)
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
