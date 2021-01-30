package main

import (
  "os"
)

var variables = map[string]string {
  "file_name": os.Args[1],
}
