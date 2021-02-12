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

        applyProperties(tokens[0], property, &contentDefaults)

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
                genLink(property)
                htmlBody += fmt.Sprintf(templates["link"], cssBody, contentProp["cLink"], cssBody, contentProp["cLinkTitle"])

            case "$mail":
                genMail(property)
                htmlBody += fmt.Sprintf(templates["mail"], cssBody, contentProp["cMailAddress"], cssBody, contentProp["cMailTitle"])

            case "$points":
                var allPoints string = genPoints(property)
                htmlBody += fmt.Sprintf(templates["pointsBody"], contentProp["cPointsType"], cssBody, allPoints, contentProp["cPointsType"])

            case "$table":
                var tableBody string = genTable(property)
                htmlBody += fmt.Sprintf(templates["tableComplete"], cssBody, fmt.Sprintf(templates["tableBorder"], themes[pageProp["~theme"]]["fg"]), tableBody)

            case "$check":
                var checkPointsBody string = genCheck(property)
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
