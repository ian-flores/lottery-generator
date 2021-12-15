package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
	"generators"
)

func main() {
	lottery := flag.String("lottery", "all", "type of lottery")
	flag.Parse()

	switch *lottery {
	case "powerball":
		powerball()
	case "traditional":
		traditional()
	default:
		powerball()
		traditional()
		pega(2)
		pega(3)
		pega(4)
	}
}
