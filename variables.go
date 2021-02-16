package main

var variables = map[string]string {
    "<": "&lt;",
    ">": "&gt;",
    "&": "&amp;",
    "'": "&apos;",
    "\"": "&quot;",
    "-": "&#37;",
    "sp": "&nbsp;",
    "source_name": getSourceName(),
    "file_name": makeHTML(getSourceName()),
}
