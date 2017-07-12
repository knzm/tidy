package main

import (
	"log"
	"os"

	"github.com/knzm/tidy"
)

func main() {
	ns, err := tidy.ParseInput(os.Stdin)
	if err != nil {
		log.Panic(err)
	}
	for i, n := range ns {
		ans := tidy.Solve(n)
		tidy.PrintOutput(os.Stdout, i, n, ans)
	}
}
