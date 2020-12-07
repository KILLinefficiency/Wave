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

    if tokens[0] == "~title" {
      pTitle = property
    }
    if tokens[0] == "~bg" {
      pBGcolor = property
    }
    if tokens[0] == "~pic" {
      pBGimage = property
    }
    if tokens[0] == "~align" {
      pAlign = property
    }
    if tokens[0] == "~box" {
      pBox = property
    }
    if tokens[0] == "~box-style" {
      pBoxStyle = property
    }

  }

  var htmlTop string = fmt.Sprintf(`
<!--
-->
<!DOCTYPE html>
<html>
    <head>
        <title>%s</title>
        <style>
            background-color: %s;
            background-image: %s;
            text-align: %s;
            margin: %s;
            border-style: %s;
        </style>
      </head>
    `, pTitle, pBGcolor, pBGimage, pAlign, pBox, pBoxStyle)

  var htmlComplete string = htmlTop + "\n\t<body>" + htmlBody + "\n\n" + "\n\t</body>\n</html>\n"

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
