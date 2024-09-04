package main

import (
	"github.com/mingrammer/cfmt"
	"math/rand"
)

func main() {
	opts := ParseOptions()
	rand.Seed(opts.Seed)
	if err := Run(opts); err != nil {
		cfmt.Warningln(err.Error())
	}
}
