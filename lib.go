package main

import (
  "strings"
)

func copyMap(mapOrignal map[string]string, mapCopy map[string]string) {
  for key, value := range mapOrignal {
    mapCopy[key] = value
  }
}

func strMultiply(strText string, times int) string {
  var strFinal string
  for loop := 0; loop < times; loop = loop + 1 {
    strFinal = strFinal + strText
  }
  return strFinal
}

func setTheme(content string, themeName string) string {
  for colName, colCode := range themes[themeName] {
    content = strings.Replace(content, colName, colCode, -1)
  }
  pageProp["pBGcolor"] = themes[themeName]["bg"]
  return content
}
