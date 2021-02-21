package main

import (
	"os"
	"fmt"
	"strings"
	"io/ioutil"
)

func copyMap(mapOrignal map[string]string) map[string]string {
	var mapCopy = make(map[string]string)
	for key, value := range mapOrignal {
		mapCopy[key] = value
	}
	return mapCopy
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

func getSourceName() string {
	if len(os.Args) == 1 {
		return ""
	}
	return os.Args[1]
}

func makeHTML(file string) string {
	fileName := strings.Split(file, ".")

	if len(fileName) == 1 {
		var htmlName string = file + ".html"
		return htmlName
	}

	fileName[len(fileName) - 1] = ".html"
	var htmlName string = strings.Join(fileName, "")
	return htmlName
}

func splitFileText(file string) []string {
	byteStream, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Invalid file address: %s\n", file)
		os.Exit(2)
	}

	var script string = string(byteStream)
	lines := strings.Split(strings.TrimSpace(script), "\n")
	return lines
}

func makeHTMLFile(sourceFile string, htmlContent string) {
	var htmlFileName string = makeHTML(sourceFile)
	htmlFile, err := os.Create(htmlFileName)
	if err != nil {
		fmt.Printf(messageTemplates["fileNotFoundError"], htmlFileName, htmlContent)
		htmlFile.Close()
		os.Exit(3)
	}

	_, err = htmlFile.WriteString(htmlContent)
	if err != nil {
		fmt.Printf("could not write to newly created file %s beacuse of the following error %s", htmlFileName, err)
		htmlFile.Close()
		os.Exit(3)
	}
	htmlFile.Close()
}

func readFileForHTML(file string, tabNo int) string {
	textFile, _ := ioutil.ReadFile(file)
	var fileContent string = string(textFile)
	fileContent = strings.Replace(fileContent, "\n", "<br>", -1)
	fileContent = strings.Replace(fileContent, " ", "&nbsp;", -1)
	fileContent = strings.Replace(fileContent, "\t", strMultiply("&nbsp;", tabNo), -1)
	return fileContent
}
