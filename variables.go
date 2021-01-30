package main

import (
  "os"
)

var variables = map[string]string {
  "source_name": os.Args[1],
  "file_name": makeHTML(os.Args[1]),
}
