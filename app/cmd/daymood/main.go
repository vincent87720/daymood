package main

import (
	"fmt"

	"github.com/vincent87720/daymood/app/internal/launch"
)

func main() {
	fmt.Println("Hello Daymood.")
	launch.Launch(true) //debug: true, production: false
}
