#!/bin/bash

go build -o bookings cmd/web/*.go && ./bookings -dbname=bookings -dbuser=postgresadmin -cache=false -production=false -dbpass=postgresadmin
