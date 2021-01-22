package main

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
