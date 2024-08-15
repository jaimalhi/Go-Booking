# Go-Booking: REST API
- Building a REST API with Go
- Go-powered "Event Booking" API

### API Endpoints
- GET `/events` -> Get a list of availble events
- GET `events/<id>` -> Get a list of availble events by id
- POST `/events` -> Create a new bookable event (auth required)
- PUT `events/<id>` -> Update an event (auth required & only by creator)
- DELETE `events/<id>` -> Delete an event (auth required & only by creator)
- POST `/signup` -> Create a new user
- POST `/login` -> Authenticate user (Auth Token with JWT)
- POST `events/<id>/register` -> Register for an event (auth required)
- DELETE `events/<id>/register` -> Cancel registration (auth required)

### Information
- [GitHub Repo](https://github.com/jaimalhi/Go-Booking)
- Using [Gin](https://github.com/gin-gonic/gin) package to handle endpoints
- Using [SQLite](https://github.com/mattn/go-sqlite3) for database