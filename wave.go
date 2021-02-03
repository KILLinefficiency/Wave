package main

import (
  "os"
  "fmt"
  "strings"
  "strconv"
)

func main() {
  if len(os.Args) == 1 {
    fmt.Printf(messageTemplates["help"])
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

    if pageProp["~theme"] != "" {
      pageProp["~bg"] = themes[pageProp["~theme"]]["bg"]
      if contentProp["cColor"] == "black" {
        contentProp["cColor"] = themes[pageProp["~theme"]]["fg"]
      }
    }

    switch tokens[0] {
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

      case "!dim":
        widthHeight := strings.Split(property, "x")
        contentProp["cWidth"] = strings.TrimSpace(widthHeight[0])
        contentProp["cHeight"] = strings.TrimSpace(widthHeight[1])

      case "!default":
        contentProp = copyMap(contentDefaults)

      case "!color":
        property = setTheme(property, pageProp["~theme"])
        contentProp["cColor"] = property

      default:
        if strings.HasPrefix(tokens[0], "!") {
            contentProp[tokens[0]] = property
        }
    }

    cssBody = fmt.Sprintf(templates["css"], contentProp["!font"], contentProp["cColor"], contentProp["cBGcolor"], contentProp["!size"], contentProp["!align"], contentProp["!box"], contentProp["!box-style"], contentProp["!points-style"])

    switch tokens[0] {
      case "$text":
        htmlBody += fmt.Sprintf(templates["text"], cssBody, property)

      case "$file":
        tabNumber, _ := strconv.Atoi(contentProp["!tab"])
        var fileStr string = readFileForHTML(property, tabNumber)
        htmlBody += fmt.Sprintf(templates["text"], cssBody, fileStr)

      case "$nl":
        if len(tokens) == 1 {
          property = "1"
        }
        times, _ := strconv.Atoi(property)
        htmlBody += fmt.Sprintf("\t\t%s\n", strMultiply("<br>", times))

      case "$link":
        linkTitle := strings.Split(property, contentProp["!sep"])
        if len(linkTitle) > 1 {
          contentProp["cLink"] = strings.TrimSpace(linkTitle[0])
          contentProp["cLinkTitle"] = strings.TrimSpace(linkTitle[1])
        }
        htmlBody += fmt.Sprintf(templates["link"], contentProp["cLink"], cssBody, contentProp["cLinkTitle"])

      case "$mail":
        mailTitle := strings.Split(property, contentProp["!sep"])
        if len(mailTitle) > 1 {
          contentProp["cMailAddress"] = strings.TrimSpace(mailTitle[0])
          contentProp["cMailTitle"] = strings.TrimSpace(mailTitle[1])
        }
        htmlBody += fmt.Sprintf(templates["mail"], contentProp["cMailAddress"], cssBody, contentProp["cMailTitle"])

      case "$points":
        listPoints := strings.Split(property, contentProp["!sep"])
        var allPoints string
        for _, point := range listPoints {
          allPoints += fmt.Sprintf(templates["points"], strings.TrimSpace(point))
        }
        var pointsBody string = fmt.Sprintf(templates["pointsBody"], contentProp["cPointsType"], cssBody, allPoints, contentProp["cPointsType"])
        htmlBody += pointsBody

      case "$table":
        var tableBody string
        tableRows := strings.Split(property, contentProp["!colsep"])
        for _, rowValues := range tableRows {
          values := strings.Split(rowValues, contentProp["!sep"])
          var rowBody string
          for _, addValues := range values {
            rowBody += fmt.Sprintf(templates["tableValues"], templates["tableBorder"], strings.TrimSpace(addValues))
          }
          tableBody += fmt.Sprintf(templates["tableBody"], templates["tableBorder"], rowBody)
        }
        htmlBody += fmt.Sprintf(templates["tableComplete"], cssBody, templates["tableBorder"], tableBody)

      case "$check":
        checkPoints := strings.Split(property, contentProp["!sep"])
        var checkPointsBody string
        for _, points := range checkPoints {
          checkPointsBody += fmt.Sprintf(templates["checkbox"], strMultiply("&nbsp;", 2) + strings.TrimSpace(points))
        }
        htmlBody += fmt.Sprintf(templates["checkboxBody"], cssBody, checkPointsBody)

      case "$quote":
        htmlBody += fmt.Sprintf(templates["quote"], property)

      case "$pic":
        htmlBody += fmt.Sprintf(templates["image"], contentProp["!align"], contentProp["!box"], contentProp["!box-style"], contentProp["cWidth"], contentProp["cHeight"], property)

      case "$html":
        htmlBody += fmt.Sprintf("\t\t%s\n", property)
    }

  }

  var htmlTopBody string = fmt.Sprintf(templates["htmlTopBody"], pageProp["~title"])

  var htmlCSS string = fmt.Sprintf(templates["htmlCSS"], pageProp["~bg"], pageProp["~img"], pageProp["~box"], pageProp["~box-style"])
  htmlCSS = setTheme(htmlCSS, pageProp["~theme"])
  var htmlComplete string = templates["waveMark"] + htmlTopBody + htmlCSS + htmlBody + templates["htmlEnd"]

  makeHTMLfile(sourceName, htmlComplete)
}
