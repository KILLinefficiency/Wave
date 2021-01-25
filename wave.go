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

  var contentDefaults = make(map[string]string)
  copyMap(contentProp, contentDefaults)

  for _, line := range scriptLines {
    tokens := strings.Split(strings.TrimSpace(line), " ")
    property := strings.Join(tokens[1:], " ")

    switch tokens[0] {
      case "~theme":
        pageProp["pTheme"] = property
      case "~title":
        pageProp["pTitle"] = property
      case "~bg":
        pageProp["pBGcolor"] = property
        contentProp["cBGcolor"] = property
      case "~pic":
        pageProp["pBGimage"] = property
      case "~color":
        property = setTheme(property, pageProp["pTheme"])
        pageProp["pColor"] = property
        contentProp["cColor"] = property
      case "~align":
        pageProp["pAlign"] = property
        contentProp["cAlign"] = property
      case "~box":
        pageProp["pBox"] = property
      case "~box-style":
        pageProp["pBoxStyle"] = property
    }

    switch tokens[0] {
      case "!tab":
        contentProp["cTab"] = property
      case "!font":
        contentProp["cFont"] = property
      case "!size":
        contentProp["cSize"] = property
      case "!color":
        property = setTheme(property, pageProp["pTheme"])
        contentProp["cColor"] = property
      case "!box":
        contentProp["cBox"] = property
      case "!box-style":
        contentProp["cBoxStyle"] = property
      case "!align":
        contentProp["cAlign"] = property
      case "!bg":
        contentProp["cBGcolor"] = property
      case "!points-type":
        if property == "ordered" {
          contentProp["cPointsType"] = "ol"
        }
        if property == "unordered" {
          contentProp["cPointsType"] = "ul"
        }
      case "!points-style":
        contentProp["cPointsStyle"] = property
      case "!dim":
        widthHeight := strings.Split(property, "x")
        contentProp["cWidth"] = strings.TrimSpace(widthHeight[0])
        contentProp["cHeight"] = strings.TrimSpace(widthHeight[1])
      case "!sep":
        contentProp["cDelimiter"] = property
      case "!colsep":
        contentProp["cTableDelimiter"] = property
      case "!default":
        copyMap(contentDefaults, contentProp)
    }

    cssBody = fmt.Sprintf(cssTemplate, contentProp["cFont"], contentProp["cColor"], contentProp["cBGcolor"], contentProp["cSize"], contentProp["cAlign"], contentProp["cBox"], contentProp["cBoxStyle"], contentProp["cPointsStyle"])

    switch tokens[0] {
      case "$text":
        htmlBody += fmt.Sprintf(textTemplate, cssBody, property)

      case "$file":
        textFile, _ := ioutil.ReadFile(property)
        tabNumber, _ := strconv.Atoi(contentProp["cTab"])
        var fileStr string = string(textFile)
        fileStr = strings.Replace(fileStr, "\n", "<br>", -1)
        fileStr = strings.Replace(fileStr, " ", "&nbsp;", -1)
        fileStr = strings.Replace(fileStr, "\t", strMultiply("&nbsp;", tabNumber), -1)
        htmlBody += fmt.Sprintf(textTemplate, cssBody, fileStr)

      case "$nl":
        if len(tokens) == 1 {
          property = "1"
        }
        times, _ := strconv.Atoi(property)
        htmlBody += fmt.Sprintf("\t\t%s\n", strMultiply("<br>", times))

      case "$link":
        linkTitle := strings.Split(property, contentProp["cDelimiter"])
        if len(linkTitle) > 1 {
          contentProp["cLink"] = strings.TrimSpace(linkTitle[0])
          contentProp["cLinkTitle"] = strings.TrimSpace(linkTitle[1])
        }
        htmlBody += fmt.Sprintf(linkTemplate, contentProp["cLink"], cssBody, contentProp["cLinkTitle"])

      case "$mail":
        mailTitle := strings.Split(property, contentProp["cDelimiter"])
        if len(mailTitle) > 1 {
          contentProp["cMailAddress"] = strings.TrimSpace(mailTitle[0])
          contentProp["cMailTitle"] = strings.TrimSpace(mailTitle[1])
        }
        htmlBody += fmt.Sprintf(mailTemplate, contentProp["cMailAddress"], cssBody, contentProp["cMailTitle"])

      case "$points":
        listPoints := strings.Split(property, contentProp["cDelimiter"])
        var allPoints string
        for _, point := range listPoints {
          allPoints += fmt.Sprintf(pointsTemplate, strings.TrimSpace(point))
        }
        var pointsBody string = fmt.Sprintf(pointsBodyTemplate, contentProp["cPointsType"], cssBody, allPoints, contentProp["cPointsType"])
        htmlBody += pointsBody

      case "$table":
        var tableBody string
        tableRows := strings.Split(property, contentProp["cTableDelimiter"])
        for _, rowValues := range tableRows {
          values := strings.Split(rowValues, contentProp["cDelimiter"])
          var rowBody string
          for _, addValues := range values {
            rowBody += fmt.Sprintf(tableValuesTemplate, tableBorder, strings.TrimSpace(addValues))
          }
          tableBody += fmt.Sprintf(tableBodyTemplate, tableBorder, rowBody)
        }
        htmlBody += fmt.Sprintf(tableCompleteTemplate, cssBody, tableBorder, tableBody)

      case "$check":
        checkPoints := strings.Split(property, contentProp["cDelimiter"])
        var checkPointsBody string
        for _, points := range checkPoints {
          checkPointsBody += fmt.Sprintf(checkboxTemplate, strMultiply("&nbsp;", 2) + strings.TrimSpace(points))
        }
        htmlBody += fmt.Sprintf(checkboxBodyTemplate, cssBody, checkPointsBody)

      case "$quote":
        htmlBody += fmt.Sprintf(quoteTemplate, property)

      case "$pic":
        htmlBody += fmt.Sprintf(imageTemplate, contentProp["cAlign"], contentProp["cBox"], contentProp["cBoxStyle"],contentProp["cWidth"], contentProp["cHeight"], property)

      case "$html":
        htmlBody += fmt.Sprintf("\t\t%s\n", property)
    }

  }

  var htmlTopBody string = fmt.Sprintf(htmlTopBodyTemplate, pageProp["pTitle"])
  var htmlCSS string = fmt.Sprintf(htmlCSSTemplate, pageProp["pBGcolor"], pageProp["pBGimage"], pageProp["pColor"], pageProp["pAlign"], pageProp["pBox"], pageProp["pBoxStyle"])
  var htmlComplete string = waveMark + htmlTopBody + htmlCSS + htmlBody + htmlEnd

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
