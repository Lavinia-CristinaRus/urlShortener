# Backend Design Decisions
## Go and Gin
Go was selected for its performance and strong standard library. Gin provides a lightweight and fast routing framework.

## MySQL with explicit models
A relational database is used to store:
 - user accounts
 - shortened URLs
 - click records including location data (city, country)

Models are explicitly defined to maintain clarity and control.

## URL redirection and click tracking
When a shortened URL is accessed:
 - The original URL is resolved.
 - A click record is created.
 - The user is redirected immediately.
 - IP address handling

## IP address extraction
If the IP is public, it is used directly.
If the IP is localhost or 127.0.0.1, the backend retrieves the public IP using an external IP lookup service.
This ensures meaningful analytics even during local development.

## Geolocation using GeoLite databases
City and country data are resolved using local GeoLite2 databases stored in the project.
Reasons for this approach:
 - no dependency on external geolocation APIs at runtime
 - faster lookup
 - predictable results
The database was not uploaded on GitHub, as the repository is public and the database was obtained for academic purposes.
The resolved location is stored with each click entry.

## QR code generation
QR codes are generated server-side for shortened URLs to:
 - centralize logic
 - ensure consistency
 - avoid frontend duplication

## JWT-based authentication
JWT is used to secure API endpoints:
 - stateless authentication
 - easy frontend integration
 - scalable design

Clear separation of responsibilities

Frontend: UI and navigation

Backend: security, persistence, analytics, geolocation, QR generation
