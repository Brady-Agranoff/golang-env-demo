package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	testVar := os.Getenv("TEST_VAR")
	println(testVar)

}
