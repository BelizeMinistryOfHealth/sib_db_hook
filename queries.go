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

func GetArrivals(date string, dateQuery DateQuery, db *sql.DB) ([]DbArrival, error) {
	sqlStmt := fmt.Sprintf("SELECT ID, FNAME, MNAME, LNAME, SEX, PASSNUM, PHONE, CONTACTPER, CONTACTPERNUM, NATIONALITY, RESIDENCE, PORT, DATEEMBARKTION, CITYAIRPORT, TRAVELMODE, VESSELNUM, PROVINCE, TRAVELORIGIN, COUNTRYVISITED,PURPOSEOFTRIP, LENGTHSTAY, FACILITYNAME, FACILITY,FACILITYADDRESS, FACILITYDISTRICT, COUNTRYOFBIRTH, MARITALSTATUS, OCCUPATION,CREATEDAT, UPDATEDAT, TRIPID,DATEOFBIRTH FROM TH_ARRIVAL WHERE %s>='%s' LIMIT 200", dateQuery, date)
	var arrivals []DbArrival

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

	return arrivals, nil
}
