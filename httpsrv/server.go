package httpsrv

import (
	sib_db_hook "bz.moh.sibdb/hook"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Env struct {
	DB *sib_db_hook.AppDb
}

type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging logs all requests with its path and the time it took to process
func Logging() Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			start := time.Now()
			defer func() {
				log.WithFields(log.Fields{
					"path":  r.URL.Path,
					"since": time.Since(start),
				}).Info("received request")
			}()

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Method ensures that url can only be requested with a specific method, else returns a 400 Bad Request
func Method(m string) Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

func VerifyToken(authorizedToken string) Middleware {

	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if len(strings.Trim(token, "")) == 0 {
				// No Authorization Token was provided
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			if token != authorizedToken {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			f(w, r)
		}
	}
}

// EnableCors enables CORS
func EnableCors() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			f(w, r)
		}
	}
}

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Decode(body io.ReadCloser, o interface{}) error {
	b, err := ioutil.ReadAll(body)
	defer body.Close()
	if err != nil {
		return fmt.Errorf("could not parse the body posted: %v", err)
	}
	err = json.Unmarshal(b, &o)
	log.WithFields(
		log.Fields{
			"userRequest": o,
			"rawBody":     string(b)}).
		Info("unmarshalled to user request")
	if err != nil {
		return err
	}
	return nil
}

// Server handles http requests.
type Server struct {
	// Basepath is the path prefix to match.
	// Default: /api/
	Basepath string

	routes map[string]http.Handler
	// NotFound is the http.Handler to use when a resource is
	// not found.
	NotFound http.Handler
	// OnErr is called when there is an error.
	OnErr func(w http.ResponseWriter, r *http.Request, err error)
}

type ReadArrivalRequest struct {
	Date      string                `json:"date,omitempty"`
	DateQuery sib_db_hook.DateQuery `json:"dateQuery,omitempty"`
	Limit     int                   `json:"limit"`
	Offset    int                   `json:"offset"`
}

// ReadArrivals gets a request to fetch arrivals and returns the
// list of arrivals along with pagination information.
func (env *Env) ReadArrivals(w http.ResponseWriter, r *http.Request) {
	var request ReadArrivalRequest
	err := Decode(r.Body, &request)
	if err != nil {
		log.WithFields(
			log.Fields{
				"body": r.Body,
			}).Error("decoding request body failed")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if len(strings.Trim(request.Date, "")) == 0 {
		http.Error(w, "invalid date", http.StatusBadRequest)
		return
	}

	if len(strings.Trim(string(request.DateQuery), "")) == 0 {
		http.Error(w, "invalid dateQuery", http.StatusBadRequest)
		return
	}

	var limit = 100
	if request.Limit > 0 {
		limit = request.Limit
	}

	queryRequest := sib_db_hook.ArrivalQueryRequest{
		Date:      request.Date,
		DateQuery: request.DateQuery,
		Cursor:    request.Offset,
		Limit:     limit,
	}
	arrivalsDb, err := env.DB.GetArrivals(queryRequest)
	if err != nil {
		http.Error(w, "error executing query", http.StatusInternalServerError)
		return
	}

	var arrivals []sib_db_hook.Arrival
	for _, arrival := range arrivalsDb.Arrivals {
		a := sib_db_hook.Arrival{
			Id:                  arrival.Id,
			FirstName:           arrival.FirstName,
			MiddleName:          arrival.MiddleName,
			LastName:            arrival.LastName,
			Sex:                 arrival.Sex,
			PassportNumber:      arrival.PassportNumber,
			PhoneNumber:         arrival.PhoneNumber,
			ContactPerson:       arrival.ContactPerson,
			ContactPersonNumber: arrival.ContactPersonNumber,
			Nationality:         arrival.Nationality,
			Residence:           arrival.Residence,
			PortOfEntry:         arrival.Port,
			DateEmbarkation:     arrival.DateEmbarkation,
			CityAirport:         arrival.CityAirport,
			TravelMode:          arrival.TravelMode,
			VesselNumber:        arrival.VesselNumber,
			Province:            arrival.Province,
			TravelOrigin:        arrival.TravelOrigin,
			CountryVisited:      arrival.CountryVisited,
			PurposeOfTrip:       arrival.PurposeOfTrip,
			LengthOfStay:        arrival.LengthOfStay,
			FacilityName:        arrival.FacilityName,
			Facility:            arrival.Facility,
			FacilityAddress:     arrival.FacilityAddress,
			FacilityDistrict:    arrival.FacilityDistrict,
			CountryOfBirth:      arrival.CountryOfBirth,
			MaritalStatus:       arrival.MaritalStatus,
			Occupation:          arrival.Occupation,
			CreatedAt:           arrival.CreatedAt,
			UpdatedAt:           arrival.UpdatedAt,
			TripId:              arrival.TripId,
			DateOfBirth:         arrival.DateOfBirth,
		}
		arrivals = append(arrivals, a)
	}

	resp := sib_db_hook.ArrivalsResponse{
		Arrivals:   arrivals,
		Total:      arrivalsDb.Total,
		NextOffset: arrivalsDb.NextOffset,
	}

	// Convert arrivals to JSON and send it down in the response
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "error executing query", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonResp))
}

func (env *Env) ReadScreenings(w http.ResponseWriter, r *http.Request) {
	var request ReadArrivalRequest
	err := Decode(r.Body, &request)
	if err != nil {
		log.WithFields(
			log.Fields{
				"body": r.Body,
			}).Error("decoding request body failed")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if len(strings.Trim(request.Date, "")) == 0 {
		http.Error(w, "invalid date", http.StatusBadRequest)
		return
	}

	if len(strings.Trim(string(request.DateQuery), "")) == 0 {
		http.Error(w, "invalid dateQuery", http.StatusBadRequest)
		return
	}

	var limit = 100
	if request.Limit > 0 {
		limit = request.Limit
	}

	queryRequest := sib_db_hook.ScreeningsQueryRequest{
		Date:      request.Date,
		DateQuery: request.DateQuery,
		Cursor:    request.Offset,
		Limit:     limit,
	}

	dbScreenings, err := env.DB.GetScreenings(queryRequest)
	if err != nil {
		http.Error(w, "error executing query", http.StatusInternalServerError)
		return
	}

	var screenings []sib_db_hook.Screening
	for _, screening := range dbScreenings.Screenings {

		s := sib_db_hook.Screening{
			Id:                      screening.Id,
			DiagnosedWithCovid:      screening.DiagnosedWithCovid,
			CovidTest:               screening.CovidTest,
			ContactedHealthFacility: screening.ContactedHealthFacility,
			ContactWithCovidCase:    screening.ContactWithCovidCase,
			SymptomDate:             screening.SymptomDate.String,
			Fever:                   screening.Fever,
			Cough:                   screening.Cough,
			ShortBreath:             screening.ShortBreath,
			DifficultyBreathing:     screening.DifficultyBreathing,
			SoreThroat:              screening.SoreThroat,
			Headache:                screening.Headache,
			Malaise:                 screening.Malaise,
			Diarrhea:                screening.Diarrhea,
			Vomitting:               screening.Vomitting,
			Bleeding:                screening.Bleeding,
			JointPains:              screening.JointPains,
			EyePain:                 screening.EyePain,
			GeneralizedRash:         screening.GeneralizedRash,
			BlurredVision:           screening.BlurredVision,
			OtherSymptoms:           screening.OtherSymptoms.String,
			CreatedAt:               screening.CreatedAt,
			UpdatedAt:               screening.UpdatedAt,
			Temperature:             screening.Temperature.String,
			TripId:                  screening.TripId,
		}
		screenings = append(screenings, s)
	}

	resp := sib_db_hook.ScreeningResponse{
		Screenings: screenings,
		Total:      dbScreenings.Total,
		NextOffset: dbScreenings.NextOffset,
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "error executing query", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonResp))

}

// HealthCheck returns OK if the server is running
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
