# Backend Architecture
The backend is a RESTful API written in Go, using the Gin framework. It is responsible for authentication, URL generation, redirection, click tracking, analytics, QR code generation, and geolocation.

## Folder Structure
backend/
├── controllers/
│   ├── user/
│   │   └── user.go
│   ├── url/
│   │   ├── generateUrl.go
│   │   ├── generateQrCode.go
│   │   ├── manageUrl.go
│   │   └── redirectUrl.go
│   └── click/
│       └── analytics.go
├── database/
│   └── mysql.go
├── geoLite/
│   ├── GeoLite2-City.mmdb
│   └── GeoLite2-Country.mmdb
├── models/
│   ├── user.go
│   ├── url.go
│   └── click.go
├── utils/
│   ├── jwt.go
│   └── middleware.go
└── main.go

## Architectural Responsibilities

### Controllers
Define HTTP endpoints
Handle request parsing and responses

### Models
Define database entities (users, URLs, clicks)

### Database
Initializes and manages MySQL connection

### Utils
JWT handling and authentication middleware

### GeoLite
Local databases used for city and country lookup, not uploaded on GitHub, was obtained from the website mentioned on the github.com/oschwald/maxminddb-golang/v2 library for geo location.

### Main
Application entry point and route registration

## Request Lifecycle
HTTP request reaches Gin router.
Authentication middleware validates JWT (if required).
Controller processes request.
Database operations are executed.
Response or redirect is returned.
