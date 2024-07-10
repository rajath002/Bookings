# Booking and Reservations

This is the repository for bookings and reservations project.

- Built in Go 1.21.0
- Uses the [Chi router](github.com/go-chi/chi/v5)
- Uses [alex edwards SCS](github.com/alexedwards/scs/v2) session management
- Uses [nosurf](github.com/justinas/nosurf)

## Running Tests


> Run all tests in one command using below command (Make sure your in root of your application)
```bash
go test -v ./...
```
### OR

go to `./cmd/web` path

> Run below commands one by one, by going to the test files location


```bash
go test 

// or 

go test -v
```

```bash
go test -cover ./...
```
- add `./...` only when you want to run it from Root (Globally)

```bash
go test -coverprofile=coverage
```

```bash
go tool cover -html=coverage
```

Or you can combine and run last two commands. but it won't work in Windows

```bash
go test -coverprofile=coverage.out && go tool cover -html=coverage.out
```


## Creating Migrations tables

### used below commands to generate migration tables

```bash
soda generate fizz CreateRoomsTable

soda generate fizz CreateRestrictionTable  

soda generate fizz createReservationsTable

soda generate fizz CreateRoomRestrictionsTable

soda generate fizz CreateUsersTable
```

### Add Foreign key to a table

```bash
soda generate fizz CreateFKForReservationsTable
```

### Run below command to create tables and to add Foreign Keys

```bash
soda migrate 
```

### Run below command to Remove the table

```bash
soda migrate down 
```
