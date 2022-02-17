package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	var message string = GetMessage()

	fmt.Println(message)
}

func GetMessage() string {
	return "Hello " + emoji.Sprint(":world_map:") + "!"
}
