package sib_db_hook

// Arrival is the struct that will be sent over the REST API
type Arrival struct {
	Id                  string `json:"id"`
	FirstName           string `json:"firstName"`
	MiddleName          string `json:"middleName"`
	LastName            string `json:"LastName"`
	Sex                 string `json:"sex"`
	PassportNumber      string `json:"passportNumber"`
	PhoneNumber         string `json:"phoneNumber"`
	ContactPerson       string `json:"contactPerson"`
	ContactPersonNumber string `json:"contactPersonNumber"`
	Nationality         string `json:"nationality"`
	Residence           string `json:"residence"`
	PortOfEntry         string `json:"portOfEntry"`
	DateEmbarkation     string `json:"dateEmbarkation"`
	CityAirport         string `json:"cityAirport"`
	TravelMode          string `json:"travelMode"`
	VesselNumber        string `json:"vesselNumber"`
	Province            string `json:"province"`
	TravelOrigin        string `json:"travelOrigin"`
	CountryVisited      string `json:"countryVisited"`
	PurposeOfTrip       string `json:"purposeOfTrip"`
	LengthOfStay        string `json:"lengthOfStay"`
	FacilityName        string `json:"facilityName"`
	Facility            string `json:"facility"`
	FacilityAddress     string `json:"facilityAddress"`
	FacilityDistrict    string `json:"facilityDistrict"`
	CountryOfBirth      string `json:"countryOfBirth"`
	MaritalStatus       string `json:"maritalStatus"`
	Occupation          string `json:"occupation"`
	CreatedAt           string `json:"createdAt"`
	UpdatedAt           string `json:"updatedAt"`
	TripId              string `json:"tripId"`
	DateOfBirth         string `json:"dateOfBirth"`
}

// DbArrival represents the data we are fetching from the database
type DbArrival struct {
	Id                  string `json:"ID"`
	FirstName           string `json:"FNAME"`
	MiddleName          string `json:"MNAME"`
	LastName            string `json:"LNAME"`
	Sex                 string `json:"SEX"`
	PassportNumber      string `json:"PASSNUM"`
	PhoneNumber         string `json:"PHONE"`
	ContactPerson       string `json:"CONTACTPER"`
	ContactPersonNumber string `json:"CONTACTPERNUM"`
	Nationality         string `json:"NATIONALITY"`
	Residence           string `json:"RESIDENCE"`
	Port                string `json:"PORT"`
	DateEmbarkation     string `json:"DATEEMBARKTION"`
	CityAirport         string `json:"CITYAIRPORT"`
	TravelMode          string `json:"TRAVELMODE"`
	VesselNumber        string `json:"VESSELNUM"`
	Province            string `json:"PROVINCE"`
	TravelOrigin        string `json:"TRAVELORIGIN"`
	CountryVisited      string `json:"COUNTRYVISITED"`
	PurposeOfTrip       string `json:"PURPOSEOFTRIP"`
	LengthOfStay        string `json:"LENGTHSTAY"`
	FacilityName        string `json:"FACILITYNAME"`
	Facility            string `json:"FACILITY"`
	FacilityAddress     string `json:"FACILITYADDRESS"`
	FacilityDistrict    string `json:"FACILITYDISTRICT"`
	CountryOfBirth      string `json:"COUNTRYOFBIRTH"`
	MaritalStatus       string `json:"MARITALSTATUS"`
	Occupation          string `json:"OCCUPATION"`
	CreatedAt           string `json:"CREATEDAT"`
	UpdatedAt           string `json:"UPDATEDAT"`
	TripId              string `json:"TRIPID"`
	DateOfBirth         string `json:"DATEOFBIRTH"`
}

type ArrivalsResponse struct {
	Arrivals   []Arrival
	Total      int
	NextOffset int
}
