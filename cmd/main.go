package main

import (
	. "bz.moh.sibdb/hook"
	"bz.moh.sibdb/hook/httpsrv"
	_ "github.com/godror/godror"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

// Initialize logger
func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func main() {
	var stage = os.Getenv("MOH_API_STAGE")
	if len(stage) == 0 {
		stage = "dev"
	}
	cnf, err := ReadConf("sample_cnf.yaml", stage)
	if err != nil {
		log.WithFields(
			log.Fields{"error": err}).Error("failure reading cnf file")
		log.Panic("error starting up server")
	}
	db, err := CreateConnection(cnf)
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Panic("error connecting to database")
	}

	env := NewEnv(db)

	http.HandleFunc("/api/arrivals",
		httpsrv.Chain(
			env.ReadArrivals,
			httpsrv.Method("POST"),
			httpsrv.Logging(),
			httpsrv.VerifyToken(cnf.ApiToken),
		),
	)
	http.HandleFunc("/api/check",
		httpsrv.Chain(
			httpsrv.HealthCheck,
			httpsrv.Method("GET"),
			httpsrv.Logging(),
		),
	)
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Panicf("could not start up server: %v", err)
	}
}

func NewEnv(db *AppDb) *httpsrv.Env {
	return &httpsrv.Env{
		db,
	}
}
