package main

import (
  "os"
  "fmt"
  "strings"
  "strconv"
  "io/ioutil"
)

func main() {
  if len(os.Args) == 1 {
    fmt.Printf("No Wave Script passed.\nPass in a Wave Script as a command-line argument.\nLike:\n\twave <scriptName>\n")
    os.Exit(1)
  }

  var htmlBody string
  var cssBody string
  var sourceName string = os.Args[1]

  scriptLines := splitFileText(sourceName)

  var contentDefaults = make(map[string]string)
  contentDefaults = copyMap(contentProp)

  for _, line := range scriptLines {
    tokens := strings.Split(strings.TrimSpace(line), " ")
    var property string = strings.Join(tokens[1:], " ")

    for name, value := range variables {
      property = strings.Replace(property, "%" + name, value, -1)
    }

    if strings.HasPrefix(tokens[0], "~") && tokens[0] != "~set" {
      pageProp[tokens[0]] = property
    } else if tokens[0] == "~set" {
      var varValue string = strings.Join(tokens[2:], " ")
      variables[tokens[1]] = varValue
    }

    switch tokens[0] {
      case "!tab":
        contentProp["cTab"] = property
      case "!font":
        contentProp["cFont"] = property
      case "!size":
        contentProp["cSize"] = property
      case "!color":
        property = setTheme(property, pageProp["~theme"])
        contentProp["cColor"] = property
      case "!box":
        contentProp["cBox"] = property
      case "!box-style":
        contentProp["cBoxStyle"] = property
      case "!align":
        contentProp["cAlign"] = property
      case "!bg":
        property = setTheme(property, pageProp["~theme"])
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
        contentProp = copyMap(contentDefaults)
    }

    cssBody = fmt.Sprintf(templates["css"], contentProp["cFont"], contentProp["cColor"], contentProp["cBGcolor"], contentProp["cSize"], contentProp["cAlign"], contentProp["cBox"], contentProp["cBoxStyle"], contentProp["cPointsStyle"])

    switch tokens[0] {
      case "$text":
        htmlBody += fmt.Sprintf(templates["text"], cssBody, property)

      case "$file":
        textFile, _ := ioutil.ReadFile(property)
        tabNumber, _ := strconv.Atoi(contentProp["cTab"])
        var fileStr string = string(textFile)
        fileStr = strings.Replace(fileStr, "\n", "<br>", -1)
        fileStr = strings.Replace(fileStr, " ", "&nbsp;", -1)
        fileStr = strings.Replace(fileStr, "\t", strMultiply("&nbsp;", tabNumber), -1)
        htmlBody += fmt.Sprintf(templates["text"], cssBody, fileStr)

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
        htmlBody += fmt.Sprintf(templates["link"], contentProp["cLink"], cssBody, contentProp["cLinkTitle"])

      case "$mail":
        mailTitle := strings.Split(property, contentProp["cDelimiter"])
        if len(mailTitle) > 1 {
          contentProp["cMailAddress"] = strings.TrimSpace(mailTitle[0])
          contentProp["cMailTitle"] = strings.TrimSpace(mailTitle[1])
        }
        htmlBody += fmt.Sprintf(templates["mail"], contentProp["cMailAddress"], cssBody, contentProp["cMailTitle"])

      case "$points":
        listPoints := strings.Split(property, contentProp["cDelimiter"])
        var allPoints string
        for _, point := range listPoints {
          allPoints += fmt.Sprintf(templates["points"], strings.TrimSpace(point))
        }
        var pointsBody string = fmt.Sprintf(templates["pointsBody"], contentProp["cPointsType"], cssBody, allPoints, contentProp["cPointsType"])
        htmlBody += pointsBody

      case "$table":
        var tableBody string
        tableRows := strings.Split(property, contentProp["cTableDelimiter"])
        for _, rowValues := range tableRows {
          values := strings.Split(rowValues, contentProp["cDelimiter"])
          var rowBody string
          for _, addValues := range values {
            rowBody += fmt.Sprintf(templates["tableValues"], templates["tableBorder"], strings.TrimSpace(addValues))
          }
          tableBody += fmt.Sprintf(templates["tableBody"], templates["tableBorder"], rowBody)
        }
        htmlBody += fmt.Sprintf(templates["tableComplete"], cssBody, templates["tableBorder"], tableBody)

      case "$check":
        checkPoints := strings.Split(property, contentProp["cDelimiter"])
        var checkPointsBody string
        for _, points := range checkPoints {
          checkPointsBody += fmt.Sprintf(templates["checkbox"], strMultiply("&nbsp;", 2) + strings.TrimSpace(points))
        }
        htmlBody += fmt.Sprintf(templates["checkboxBody"], cssBody, checkPointsBody)

      case "$quote":
        htmlBody += fmt.Sprintf(templates["quote"], property)

      case "$pic":
        htmlBody += fmt.Sprintf(templates["image"], contentProp["cAlign"], contentProp["cBox"], contentProp["cBoxStyle"], contentProp["cWidth"], contentProp["cHeight"], property)

      case "$html":
        htmlBody += fmt.Sprintf("\t\t%s\n", property)
    }

  }

  var htmlTopBody string = fmt.Sprintf(templates["htmlTopBody"], pageProp["~title"])

  if pageProp["~theme"] != "" {
    pageProp["~bg"] = themes[pageProp["~theme"]]["bg"]
  } else {
    pageProp["theme"] = "Default"
  }

  var htmlCSS string = fmt.Sprintf(templates["htmlCSS"], pageProp["~bg"], pageProp["~img"], pageProp["~box"], pageProp["~box-style"])
  htmlCSS = setTheme(htmlCSS, pageProp["~theme"])
  var htmlComplete string = templates["waveMark"] + htmlTopBody + htmlCSS + htmlBody + templates["htmlEnd"]

  var htmlFileName string = makeHTML(sourceName)
  htmlFile, err := os.Create(htmlFileName)
  if err != nil {
    fmt.Printf("Unable to create file: %s\n\nSource Code for the Document:\n\n%s\n", htmlFileName, htmlComplete)
    os.Exit(3)
  }
  htmlFile.WriteString(htmlComplete)
  htmlFile.Close()
}
