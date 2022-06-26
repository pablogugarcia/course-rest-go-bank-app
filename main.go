package main

import (
	app "github.com/pablogugarcia/banking/app"
	"github.com/pablogugarcia/banking/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
