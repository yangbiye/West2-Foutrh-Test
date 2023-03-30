package main

import (
	models "github.com/videos/Models"
	"github.com/videos/router"
)

func main() {
	models.Setup()
	router.Setup()
}
