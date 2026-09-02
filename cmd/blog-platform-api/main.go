package main

import (
	"os"

	"github.com/wascool0nce/blog-platform-api/intenal/app"
)

func main() {
	if err := app.Run(); err != nil {
		os.Exit(-1)
	}
}
