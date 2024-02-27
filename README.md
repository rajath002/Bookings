# Booking and Reservations

This is the repository for bookings and reservations project.

- Built in Go 1.21.0
- Uses the [Chi router](github.com/go-chi/chi/v5)
- Uses [alex edwards SCS](github.com/alexedwards/scs/v2) session management
- Uses [nosurf](github.com/justinas/nosurf)

## Running Tests

go to `./cmd/web` path

> Run below commands one by one

then run 

```bash
go test 

// or 

go test -v
```

```bash
go test -cover
```

```bash
go test -coverprofile=coverage
```

```bash
go tool cover -html=coverage
```

Or you can combine and run last two commands. but it won't work in Windows

```bash
go test -coverprofile=coverage.out | go tool cover -html=coverage.out
```