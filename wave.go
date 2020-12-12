package main

import (
  "os"
  "fmt"
  "strings"
  "strconv"
  "io/ioutil"
)

var htmlBody string
var pTitle string = "Wave Document"
var pBGcolor string = "white"
var pBGimage string = "none"
var pAlign string = "left"
var pBox string = "0"
var pBoxStyle string = "hidden"

var cTab int = 4
var cFont string = "Arial"
var cSize string = "17"
var cColor string = "black"
var cBox string = "0"
var cBoxStyle string = "hidden"
var cAlign string
var cBGcolor string
var cPointsType string = "ul"
var cPointsStyle string = "disc"
var cWidth string = "none"
var cHeight string = "none"
var cDelimiter string = ";"
var cLink string = "https://www.github.com/KILLinefficiency/Wave"
var cTitle string = "Wave"

func strMultiply(strText string, times int) string {
  var strFinal string
  for loop := 0; loop < times; loop = loop + 1 {
    strFinal = strFinal + strText
  }
  return strFinal
}

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
      case "!default":
        cTab, cFont = 4, "Arial"
        cSize, cColor = "17", "black"
        cBox, cBoxStyle = "0", "hidden"
        cBoxStyle = "hidden"
        cAlign, cBGcolor = "", ""
        cPointsType, cPointsStyle = "ul", "disc"
        cWidth, cHeight = "none", "none"
        cLink, cTitle = "https://www.github.com/KILLinefficiency/Wave", "Wave"
        cDelimiter = ";"
    }

    switch tokens[0] {
      case "$text":
        htmlBody += fmt.Sprintf("\t\t<p style = 'font-family: %s; color: %s; background-color: %s; font-size: %spx; text-align: %s; margin: %spx; border-style: %s;'>%s</p>\n", cFont, cColor, cBGcolor, cSize, cAlign, cBox, cBoxStyle, property)

      case "$file":
        textFile, _ := ioutil.ReadFile(property)
        var fileStr string = string(textFile)
        fileStr = strings.Replace(fileStr, "\n", "<br>", -1)
        fileStr = strings.Replace(fileStr, " ", "&nbsp;", -1)
        fileStr = strings.Replace(fileStr, "\t", strMultiply("&nbsp;", cTab), -1)
        htmlBody += fmt.Sprintf("\t\t<p style = 'font-family: %s; color: %s; background-color: %s; font-size: %spx; text-align: %s; margin: %spx; border-style: %s;'>%s</p>\n", cFont, cColor, cBGcolor, cSize, cAlign, cBox, cBoxStyle,fileStr)

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
          cTitle = strings.TrimSpace(linkTitle[1])
        }
        htmlBody += fmt.Sprintf("\t\t<a href = %s style = 'font-family: %s; color: %s; background-color: %s; font-size: %spx; text-align: %s; margin: %spx; border-style: %s;'>%s</a>\n", cLink, cFont, cColor, cBGcolor, cSize, cAlign, cBox, cBoxStyle, cTitle)

      case "$points":
        listPoints := strings.Split(property, cDelimiter)
        var allPoints string
        for _, point := range listPoints {
          allPoints += fmt.Sprintf("\t\t\t<li>%s</li>\n", strings.TrimSpace(point))
        }
        var pointsBody string = fmt.Sprintf("\t\t<%s style = 'font-family: %s; color: %s; background-color: %s; font-size: %spx; text-align: %s; margin: %spx; border-style: %s; list-style-type: %s;'>\n%s\t\t</%s>\n", cPointsType, cFont, cColor,cBGcolor, cSize, cAlign, cBox, cBoxStyle, cPointsStyle, allPoints, cPointsType)
        htmlBody += pointsBody

      case "$quote":
        htmlBody += fmt.Sprintf("\t\t<br><b><i>\"%s\"</b></i><br>\n", property)

      case "$pic":
        htmlBody += fmt.Sprintf("\t\t<div style = 'text-align: %s; margin: %spx; border-style: %s;'>\n\t\t\t<img width = '%s' height = '%s' src = %s>\n\t\t</div>\n", cAlign, cBox, cBoxStyle,cWidth, cHeight, property)

      case "$html":
        htmlBody += fmt.Sprintf("\t\t%s\n", property)
    }

  }

  var htmlTop string = fmt.Sprintf(`
<!--
This Document is generated using Wave.
Wave: https://www.github.com/KILLinefficiency/Wave
-->

<!DOCTYPE html>
<html>
    <head>
        <title>%s</title>
        <style>
          body {
              background-color: %s;
              background-image: %s;
              text-align: %s;
              margin: %spx;
              border-style: %s;
          }
        </style>
      </head>
    `, pTitle, pBGcolor, pBGimage, pAlign, pBox, pBoxStyle)

  var htmlComplete string = htmlTop + "\n\t\t<body>\n" + htmlBody + "\t\t</body>\n\n</html>\n"

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
