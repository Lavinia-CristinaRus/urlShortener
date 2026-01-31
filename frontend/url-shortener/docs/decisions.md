# Frontend Design Decisions
## React for a single-page application
React was chosen to build a web application.
Vite as build and development tool

Vite provides:
 - fast development server startup
 - hot module replacement
 - minimal configuration

## Page-based structure
Each major feature is implemented as a page under pages/, making the application flow easy to follow and understand.

## Client-side routing
react-router-dom is used to:
 - enable navigation without page reloads
 - separate public and protected routes
 - keep navigation logic declarative

## Centralized API communication

All backend communication is handled through utils/api.js to:
 - avoid duplicated Axios configuration
 - ensure consistent error handling
 - simplify future API changes

## Minimalistic frontend

The frontend intentionally avoids:
 - business logic
 - data validation rules
 - analytics computation
 - QR code generation
