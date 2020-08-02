package sib_db_hook

import (
	"database/sql"
	"fmt"
)

type Datastore interface {
	GetArrivals(arrivalQueryRequest ArrivalQueryRequest) (*ArrivalQueryResponse, error)
	CountArrivals(date string, dateQuery DateQuery) (int, error)
	CountScreenings(date string, dateQuery DateQuery) (int error)
	GetScreenings(request ScreeningsQueryRequest) (*ScreeningsQueryResponse, error)
}

type AppDb struct {
	*sql.DB
}

func CreateConnection(cnf *dbConf) (*AppDb, error) {
	connstr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", cnf.DbUsername, cnf.DbDatabase, cnf.DbPassword, cnf.DbHost)
	db, err := sql.Open(cnf.DbType, connstr)
	if err != nil {
		return nil, err
	}
	return &AppDb{db}, err
}

type DateQuery string

const (
	CreatedAt = "CREATEDAT"
	UpdateAt  = "UPDATEDAT"
)

type ArrivalQueryResponse struct {
	Arrivals   []DbArrival
	NextOffset int
	Total      int
}

type ScreeningsQueryResponse struct {
	Screenings []DbScreening
	NextOffset int
	Total      int
}

type ArrivalQueryRequest struct {
	Date      string
	DateQuery DateQuery
	Cursor    int
	Limit     int
}

type ScreeningsQueryRequest struct {
	Date      string
	DateQuery DateQuery
	Cursor    int
	Limit     int
}

func (db *AppDb) CountArrivals(date string, dateQuery DateQuery) (int, error) {
	sqlStmt := fmt.Sprintf("SELECT COUNT(ID) FROM TH_ARRIVAL WHERE %s>='%s'", dateQuery, date)
	var cnt int
	err := db.QueryRow(sqlStmt).Scan(&cnt)
	if err != nil {
		// No rows were found, so cnt is 0
		if err == sql.ErrNoRows {
			cnt = 0
		} else {
			return cnt, err
		}
	}
	return cnt, err
}

func (db *AppDb) CountScreenings(date string, dateQuery DateQuery) (int, error) {
	sqlStmt := fmt.Sprintf("SELECT COUNT(ID) FROM TH_SCREENING WHERE %s>='%s'", dateQuery, date)
	var cnt int
	err := db.QueryRow(sqlStmt).Scan(&cnt)
	if err != nil {
		// No rows were found, so cnt is 0
		if err == sql.ErrNoRows {
			cnt = 0
		} else {
			return cnt, err
		}
	}
	return cnt, err
}

func (db *AppDb) GetArrivals(arrivalQueryRequest ArrivalQueryRequest) (*ArrivalQueryResponse, error) {
	dateQuery := arrivalQueryRequest.DateQuery
	date := arrivalQueryRequest.Date
	limit := arrivalQueryRequest.Limit
	offset := arrivalQueryRequest.Cursor
	sqlStmt := fmt.Sprintf("SELECT ID, FNAME, MNAME, LNAME, SEX, PASSNUM, PHONE, CONTACTPER, CONTACTPERNUM, NATIONALITY, RESIDENCE, PORT, DATEEMBARKTION, CITYAIRPORT, TRAVELMODE, VESSELNUM, PROVINCE, TRAVELORIGIN, COUNTRYVISITED,PURPOSEOFTRIP, LENGTHSTAY, FACILITYNAME, FACILITY,FACILITYADDRESS, FACILITYDISTRICT, COUNTRYOFBIRTH, MARITALSTATUS, OCCUPATION,CREATEDAT, UPDATEDAT, TRIPID,DATEOFBIRTH FROM TH_ARRIVAL WHERE %s>='%s' OFFSET %d LIMIT %d", dateQuery, date, offset, limit)
	var arrivals []DbArrival

	// Count total records that match this query
	total, err := db.CountArrivals(date, dateQuery)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var arrival DbArrival
		//unmarshal the row object
		err = rows.Scan(
			&arrival.Id,
			&arrival.FirstName,
			&arrival.MiddleName,
			&arrival.LastName,
			&arrival.Sex,
			&arrival.PassportNumber,
			&arrival.PhoneNumber,
			&arrival.ContactPerson,
			&arrival.ContactPersonNumber,
			&arrival.Nationality,
			&arrival.Residence,
			&arrival.Port,
			&arrival.DateEmbarkation,
			&arrival.CityAirport,
			&arrival.TravelMode,
			&arrival.VesselNumber,
			&arrival.Province,
			&arrival.TravelOrigin,
			&arrival.CountryVisited,
			&arrival.PurposeOfTrip,
			&arrival.LengthOfStay,
			&arrival.FacilityName,
			&arrival.Facility,
			&arrival.FacilityAddress,
			&arrival.FacilityDistrict,
			&arrival.CountryOfBirth,
			&arrival.MaritalStatus,
			&arrival.Occupation,
			&arrival.CreatedAt,
			&arrival.UpdatedAt,
			&arrival.TripId,
			&arrival.DateOfBirth,
		)
		if err != nil {
			return nil, err
		}
		arrivals = append(arrivals, arrival)
	}
	var nextOffset int
	if offset+limit < total {
		nextOffset = offset + limit
	}
	resp := &ArrivalQueryResponse{
		Arrivals:   arrivals,
		NextOffset: nextOffset,
		Total:      total,
	}
	return resp, nil
}

func (db *AppDb) GetScreenings(request ScreeningsQueryRequest) (*ScreeningsQueryResponse, error) {
	dateQuery := request.DateQuery
	date := request.Date
	limit := request.Limit
	offset := request.Cursor
	sqlStmt := fmt.Sprintf("SELECT ID, DIAGNOSECOVID, COVIDTEST, CONTACTHEALTH, CONTACTCOVID, SYMPTOMDATE, FEVER, COUGH, SHORTBREATH, DIFBREATH, SORETHROAT, HEADACHE, MALAISE, DIARRHEA, VOMITTING, BLEEDING, JOINT, EYEPAIN, GENERALIZEDRASH,BLURREDVISION, OTHERSYMP, CREATEDAT, UPDATEDAT, TEMPERATURE, TRIPID  FROM TH_SCREENING WHERE %s>='%s' OFFSET %d LIMIT %d", dateQuery, date, offset, limit)
	var screenings []DbScreening

	// Count total records that match this query
	total, err := db.CountScreenings(date, dateQuery)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var screening DbScreening
		err = rows.Scan(
			&screening.Id,
			&screening.DiagnosedWithCovid,
			&screening.CovidTest,
			&screening.ContactedHealthFacility,
			&screening.ContactWithCovidCase,
			&screening.SymptomDate,
			&screening.Fever,
			&screening.Cough,
			&screening.ShortBreath,
			&screening.DifficultyBreathing,
			&screening.SoreThroat,
			&screening.Headache,
			&screening.Malaise,
			&screening.Diarrhea,
			&screening.Vomitting,
			&screening.Bleeding,
			&screening.JointPains,
			&screening.EyePain,
			&screening.GeneralizedRash,
			&screening.BlurredVision,
			&screening.OtherSymptoms,
			&screening.CreatedAt,
			&screening.UpdatedAt,
			&screening.Temperature,
			&screening.TripId)
		if err != nil {
			return nil, err
		}

		screenings = append(screenings, screening)
	}
	var nextOffset int
	if offset+limit < total {
		nextOffset = offset + limit
	}

	resp := &ScreeningsQueryResponse{
		Screenings: screenings,
		NextOffset: nextOffset,
		Total:      total,
	}
	return resp, nil
}
