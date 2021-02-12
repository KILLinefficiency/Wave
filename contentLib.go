package main

import (
    "fmt"
    "strings"
)

func applyProperties(keyword string, property string, defaultMap *map[string]string) {
    switch keyword {
        case "!bg":
            contentProp["cBGcolor"] = setTheme(property, pageProp["~theme"])

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
            contentProp = copyMap(*defaultMap)

        case "!color":
            property = setTheme(property, pageProp["~theme"])
            contentProp["cColor"] = property

        default:
            if strings.HasPrefix(keyword, "!") {
                contentProp[keyword] = property
            }
    }
}

func genLink(linkContent string) {
    linkTitle := strings.Split(linkContent, contentProp["!sep"])
    if len(linkTitle) > 1 {
        contentProp["cLink"] = strings.TrimSpace(linkTitle[0])
        contentProp["cLinkTitle"] = strings.TrimSpace(linkTitle[1])
    }
}

func genMail(mailContent string) {
    mailTitle := strings.Split(mailContent, contentProp["!sep"])
    if len(mailTitle) > 1 {
        contentProp["cMailAddress"] = strings.TrimSpace(mailTitle[0])
        contentProp["cMailTitle"] = strings.TrimSpace(mailTitle[1])
    }
}

func genPoints(pointsContent string) string {
    listPoints := strings.Split(pointsContent, contentProp["!sep"])
    var allPoints string
    for _, point := range listPoints {
        allPoints += fmt.Sprintf(templates["points"], strings.TrimSpace(point))
    }
    return allPoints
}

func genTable(tableContent string) string {
    var tableBody string
    tableRows := strings.Split(tableContent, contentProp["!colsep"])
    for _, rowValues := range tableRows {
        values := strings.Split(rowValues, contentProp["!sep"])
        var rowBody string
        for _, addValues := range values {
            rowBody += fmt.Sprintf(templates["tableValues"], fmt.Sprintf(templates["tableBorder"], themes[pageProp["~theme"]]["fg"]), strings.TrimSpace(addValues))
        }
        tableBody += fmt.Sprintf(templates["tableBody"], fmt.Sprintf(templates["tableBorder"], themes[pageProp["~theme"]]["fg"]), rowBody)
    }
    return tableBody
}

func genCheck(checkContent string) string {
    var checkPointsBody string
    checkPoints := strings.Split(checkContent, contentProp["!sep"])
    for _, points := range checkPoints {
        checkPointsBody += fmt.Sprintf(templates["checkbox"], strMultiply("&nbsp;", 2) + strings.TrimSpace(points))
    }
    return checkPointsBody
}
