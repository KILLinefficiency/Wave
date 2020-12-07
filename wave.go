package main

import (
  "os"
  "fmt"
  "strings"
  "io/ioutil"
)

var htmlBody string
var pTitle string = "Wave Document"
var pBGcolor string = "white"
var pBGimage string = "none"
var pAlign string = "left"
var pBox string = "0"
var pBoxStyle string = "hidden"


var cSize string = "17"
var cColor string = "black"
var cBox string = "0"
var cAlign string
var cBGcolor string

func main() {
  var sourceName string = os.Args[1]

  byteStream, err := ioutil.ReadFile(sourceName)
  if err != nil {
    fmt.Printf("Invalid file address: %s\n", sourceName)
    os.Exit(1)
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
      case "~pic":
        pBGimage = property
      case "~align":
        pAlign = property
      case "~box":
        pBox = property
      case "~box-style":
        pBoxStyle = property
    }

    cAlign = pAlign
    cBGcolor = pBGcolor

    switch tokens[0] {
      case "!size":
        cSize = property
      case "!color":
        cColor = property
      case "!box":
        cBox = property
      case "!align":
        cAlign = property
      case "!bg":
        cBGcolor = property
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
              margin: %s;
              border-style: %s;
          }
        </style>
      </head>
    `, pTitle, pBGcolor, pBGimage, pAlign, pBox, pBoxStyle)

  var htmlComplete string = htmlTop + "\n\t\t<body>" + htmlBody + "\n\n" + "\n\t\t</body>\n\n</html>\n"

  fileName := strings.Split(sourceName, ".")
  if len(fileName) == 1 {
    fileName = append(fileName, ".html")
  } else {
    fileName[len(fileName) - 1] = ".html"
  }

  htmlFileName := strings.Join(fileName, "")
  htmlFile, err := os.Create(htmlFileName)
  if err != nil {
    fmt.Printf("Unable to create file: %s\n", htmlFileName)
    os.Exit(1)
  }
  htmlFile.WriteString(htmlComplete)
}
