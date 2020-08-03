create table if not exists th_arrival
(
	id varchar,
	fname varchar,
	mname varchar,
	lname varchar,
	sex varchar,
	passnum varchar,
	phone varchar,
	contactper varchar,
	contactpernum varchar,
	nationality varchar,
	residence varchar,
	port varchar,
	dateembarktion varchar,
	cityairport varchar,
	travelmode varchar,
	vesselnum varchar,
	province varchar,
	travelorigin varchar,
	countryvisited varchar,
	purposeoftrip varchar,
	lengthstay varchar,
	facilityname varchar,
	facility varchar,
	facilityaddress varchar,
	facilitydistrict varchar,
	countryofbirth varchar,
	maritalstatus varchar,
	occupation varchar,
	createdat timestamp,
	updatedat timestamp,
	tripid varchar,
	dateofbirth date
);

alter table th_arrival owner to postgres;

create table if not exists th_screening
(
    id varchar,
    diagnosecovid varchar,
    covidtest varchar,
    contactcovid varchar,
    contacthealth varchar,
    symptomdate varchar,
    fever varchar,
    cough varchar,
    shortbreath varchar,
    difbreath varchar,
    sorethroat varchar,
    headache varchar,
    malaise varchar,
    diarrhea varchar,
    vomitting varchar,
    bleeding varchar,
    joint varchar,
    eyepain varchar,
    generalizedrash varchar,
    blurredvision varchar,
    othersymp varchar,
    createdat timestamp,
    updatedat timestamp,
    tripid varchar,
    temperature varchar
);

alter table th_screening owner to postgres;
