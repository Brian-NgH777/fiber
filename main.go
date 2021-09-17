package main

import (
	"fiber/cmd"
)

var revision string

func main() {
	cmd.SetRevision(revision)
	cmd.Execute()
}
