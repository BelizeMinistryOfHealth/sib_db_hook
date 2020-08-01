package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

// Set routes for
// - getArrivals
// - getScreenings

// - Create a database connection

// Initialize logger
func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}
