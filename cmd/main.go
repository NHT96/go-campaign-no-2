package main

import (
	"fmt"

	"github.com/NHT96/go-campaign-no-1/pkg/greetings"
)

func main() {
	message := greetings.Hello("Huu Tai")
	fmt.Println(message)
}
