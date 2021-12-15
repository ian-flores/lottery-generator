package main

import (
	"flag"
	"github.com/ian-flores/lottery-generator/pkg/generators"
)

func main() {
	lottery := flag.String("lottery", "all", "type of lottery")
	flag.Parse()

	switch *lottery {
	case "powerball":
		generators.Powerball()
	case "traditional":
		generators.Traditional()
	default:
		generators.Powerball()
		generators.Traditional()
		generators.Pega(2)
		generators.Pega(3)
		generators.Pega(4)
	}
}
