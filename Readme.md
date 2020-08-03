# SIB DB Hook

This is an http server that serves as an api for the Ministry of Health's (MOH) need to ingest the data related to guest arrivals in the 
country during the COVID19 pandemic.

## Motivations

The MOH has a number of tools (some built in-house, others donated, like Go.Data) that it uses to manage the COVID19 outbreaks.
These tools assist the Public Health Officers in screening people and quickly make the data available to the Epidemiology
Unit (EPI). The MOH has used the Repatriation Programme to test these tools and will expand their usage to the new protocols
for screening guests travelling to Belize. The Public Health officers will be able to screen arrivals the international 
airport (PGIA) if the demographic data collected by the SIB mobile app feeds into MOH's `Port of Entry` app. To that end,
the MOH requested an API from SIB to ingest the data before the guests arrive in the country. The server in this repository
is an attempt to facilitate that process. It is an http server that exposes an endpoint for retrieving arrivals and another
for retrieving screenings. MOH will periodically make requests to these endpoints and ingest the data and route them
to the pertinent systems.

## Requirements
1. A Golang compiler (https://golang.org)
2. Credentials for the database

## Compiling
To compile for Linux: 
```
make build
```

To compile for MacOS:
```
make buildMacos
```

Compilation will produce a binary under `./bin` called `moh_api_server`. To start the server type `./bin/moh_api_server`
on the command line. The server will then be available at `http://localhost:3000`.

## Add Database credentials
The server makes use of a configuration file called `moh_api_cnf.yaml` which should be placed under the same directory
where the binary is being executed from. Add the relevant credentials to the `moh_api_cnf.yaml` file under the `prod` 
section. Use `oracle` as the value for `DbType` for a production environment.

## Running the server in dev mode
A postgresql database is required for dev mode. A `docker-compose.yml` file is provided to facilitate this if you have
[docker](https://www.docker.com) and [docker compose](https://docs.docker.com/compose/) installed locally.

Start up the database with `docker-compose up`.
Create the databases in the `tables_schema.sql` file. You can populate the database with test data using the
`arrival_test_data.csv` and `screening_test_data.csv` files.
Compile the server and run it.

Retrieve screenings:
```
curl --location --request POST 'http://localhost:3000/api/screenings' \
--header 'Authorization: token' \
--header 'Content-Type: application/json' \
--data-raw '{
    "date": "2020-08-02",
    "dateQuery": "CREATEDAT",
    "limit": 100
}'
```

Retrieve arrivals:
```
curl --location --request POST 'http://localhost:3000/api/arrivals' \
--header 'Authorization: token' \
--header 'Content-Type: application/json' \
--data-raw '{
    "date": "2020-07-28",
    "dateQuery": "CREATEDAT",
    "limit": 100
}'
```



