package main

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"github.com/rogdevil/ago/src/agoproduct"
	helloworld "github.com/rogdevil/ago/src/helloWorld"
)

func main() {
	flamingo.App(
		[]dingo.Module{
			new(helloworld.Module),
			new(agoproduct.Module),
		},
	)
}
