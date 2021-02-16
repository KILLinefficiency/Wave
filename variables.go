package main

var variables = map[string]string {
    "<": "&lt;",
    ">": "&gt;",
    "&": "&amp;",
    "\"": "&quot;",
    "'": "&apos;",
    "-": "&#37;",
    "sp": "&nbsp;",
    "source_name": getSourceName(),
    "file_name": makeHTML(getSourceName()),
}
