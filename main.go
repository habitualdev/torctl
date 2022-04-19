package main

import (
	"torctl/ui"
)

func main() {
	t := ui.GetDetails()
	t.New()
	t.StartUI()

}
