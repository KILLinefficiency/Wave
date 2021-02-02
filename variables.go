package main

var variables = map[string]string {
  "source_name": getSourceName(),
  "file_name": makeHTML(getSourceName()),
}
