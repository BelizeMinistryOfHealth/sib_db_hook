package sib_db_hook

type Screening struct {
	Id                      string `json:"id"`
	DiagnosedWithCovid      bool   `json:"diagnosedWithCovid"`
	CovidTest               bool   `json:"covidTest"`
	ContactedHealthFacility bool   `json:"contactedHealthFacility"`
	ContactWithCovidCase    bool   `json:"contactWithCovidCase"`
	SymptomDate             string `json:"symptomDate,omitempty"`
	Fever                   bool   `json:"fever"`
	Cough                   bool   `json:"cough"`
	ShortBreath             bool   `json:"shortBreath"`
	DifficultyBreathing     bool   `json:"difficultyBreathing"`
	SoreThroat              bool   `json:"soreThroat"`
	Headache                bool   `json:"headache"`
	Malaise                 bool   `json:"malaise"`
	Diarrhea                bool   `json:"diarrhea"`
	Vomitting               bool   `json:"vomitting"`
	Bleeding                bool   `json:"bleeding"`
	JointPains              bool   `json:"jointPains"`
	EyePain                 bool   `json:"eyePain"`
	GeneralizedRash         bool   `json:"generalizedRash"`
	BlurredVision           bool   `json:"blurredVision"`
	OtherSymptoms           string `json:"otherSymptoms,omitempty"`
	CreatedAt               string `json:"createdAt"`
	UpdatedAt               string `json:"updatedAt"`
	Temperature             int    `json:"temperature"`
	TripId                  string `json:"tripId"`
}

type DbScreening struct {
	Id                      string `json:"ID"`
	DiagnosedWithCovid      bool   `json:"DIAGNOSECOVID"`
	CovidTest               bool   `json:"COVIDTEST"`
	ContactedHealthFacility bool   `json:"CONTACTHEALTH"`
	ContactWithCovidCase    bool   `json:"CONTACTCOVID"`
	SymptomDate             string `json:"SYMPTOMDATE,omitempty"`
	Fever                   bool   `json:"FEVER"`
	Cough                   bool   `json:"COUGH"`
	ShortBreath             bool   `json:"SHORTBREATH"`
	DifficultyBreathing     bool   `json:"DIFBREATH"`
	SoreThroat              bool   `json:"SORETHROAT"`
	Headache                bool   `json:"HEADACHE"`
	Malaise                 bool   `json:"MALAISE"`
	Diarrhea                bool   `json:"DIARRHEA"`
	Vomitting               bool   `json:"VOMITTING"`
	Bleeding                bool   `json:"BLEEDING"`
	JointPains              bool   `json:"JOINT"`
	EyePain                 bool   `json:"EYEPAIN"`
	GeneralizedRash         bool   `json:"GENERALIZEDRASH"`
	BlurredVision           bool   `json:"BLURREDVISION"`
	OtherSymptoms           string `json:"OTHERSYMP,omitempty"`
	CreatedAt               string `json:"CREATEDAT"`
	UpdatedAt               string `json:"UPDATEDAT"`
	Temperature             int    `json:"TEMPERATURE,omitempty"`
	TripId                  string `json:"TRIPID"`
}
