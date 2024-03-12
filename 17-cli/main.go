package main

import (
	"cli/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("initial")

	app := app.Generate()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
