# Frontend Architecture

The frontend is a web application built with React and Vite. It is responsible for user interaction, navigation, and sending requests to the backend API.
The frontend does not perform business logic or data persistence; it strictly acts as a client.

## Folder Structure
frontend/url-shortener/src/
├── pages/
│   ├── SignIn.jsx
│   ├── SignUp.jsx
│   ├── UrlList.jsx
│   ├── CustomizeUrl.jsx
│   ├── Analytics.jsx
│   └── QrCode.jsx
├── router/
│   └── Router.jsx
├── utils/
│   └── api.js
├── App.jsx
├── main.jsx
├── App.css
└── index.css

## Responsibilities by Folder
1. pages/
Each file in this folder represents a full page/view in the application:
 - SignIn - Login form for existing users
 - SignUp - Registration form for new users
 - UrlList - Main dashboard displaying the user’s shortened URLs
 - CustomizeUrl - Allows modifying properties of an existing short URL
 - Analytics - Displays click statistics using charts
 - QrCode - Displays the QR code associated with a short URL

2. router/
Router.jsx defines application routes using react-router-dom.
Public routes: Sign In, Sign Up
Protected routes: Url list, customization, analytics, QR code

Routing controls which page is rendered based on the URL and authentication state.

3. utils
 - api.js
Centralizes all HTTP communication with the backend.
Responsibilities:
 - Axios configuration
 - Defining API calls (authentication, URL management, analytics)
 - Attaching authentication headers

## Application Flow
 - main.jsx bootstraps the React application.
 - App.jsx loads global styles and routing.
 - Router.jsx determines which page to render.
 - Pages call functions from utils/api.js.
 - Responses update component state and re-render the UI.
