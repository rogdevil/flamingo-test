package main

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"github.com/rogdevil/flamingo-test/src/agoproduct"
	helloworld "github.com/rogdevil/flamingo-test/src/helloWorld"
)

func main() {
	flamingo.App(
		[]dingo.Module{
			new(helloworld.Module),
			new(agoproduct.Module),
		},
	)
}
