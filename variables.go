package main

var variables = map[string]string {
    "source_name": getSourceName(),
    "file_name": makeHTML(getSourceName()),
    "sp": "&nbsp;",
    "<": "&lt;",
    ">": "&gt;",
    "&": "&amp;",
    "\"": "&quot;",
    "'": "&apos;",
    "-": "&#37;",
}
