package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	fmt.Println("Hello")

	if 1 == 2 {
		return errors.New("random error")
	}
	return nil
}
