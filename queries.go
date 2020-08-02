package sib_db_hook

import (
	"database/sql"
	"fmt"
)

func CreateConnection(cnf *dbConf) (*sql.DB, error) {
	connstr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", cnf.DbUsername, cnf.DbDatabase, cnf.DbPassword, cnf.DbHost)
	db, err := sql.Open(cnf.DbType, connstr)
	return db, err
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

type ArrivalQueryRequest struct {
	Date      string
	DateQuery DateQuery
	Cursor    int
	Limit     int
}

func countArrivals(date string, dateQuery DateQuery, db *sql.DB) (int, error) {
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

func GetArrivals(arrivalQueryRequest ArrivalQueryRequest, db *sql.DB) (*ArrivalQueryResponse, error) {
	dateQuery := arrivalQueryRequest.DateQuery
	date := arrivalQueryRequest.Date
	limit := arrivalQueryRequest.Limit
	offset := arrivalQueryRequest.Cursor
	sqlStmt := fmt.Sprintf("SELECT ID, FNAME, MNAME, LNAME, SEX, PASSNUM, PHONE, CONTACTPER, CONTACTPERNUM, NATIONALITY, RESIDENCE, PORT, DATEEMBARKTION, CITYAIRPORT, TRAVELMODE, VESSELNUM, PROVINCE, TRAVELORIGIN, COUNTRYVISITED,PURPOSEOFTRIP, LENGTHSTAY, FACILITYNAME, FACILITY,FACILITYADDRESS, FACILITYDISTRICT, COUNTRYOFBIRTH, MARITALSTATUS, OCCUPATION,CREATEDAT, UPDATEDAT, TRIPID,DATEOFBIRTH FROM TH_ARRIVAL WHERE %s>='%s' OFFSET %d LIMIT %d", dateQuery, date, offset, limit)
	var arrivals []DbArrival

	// Count total records that match this query
	total, err := countArrivals(date, dateQuery, db)
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
