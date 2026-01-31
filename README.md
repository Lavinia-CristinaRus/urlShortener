# URL Shortener
## Project Description
A full-stack web application for creating shortened URLs with user authentication and basic click statistics. The project combines a Go backend (Gin + MySQL) with a React frontend (Vite). The backend API handles authentication and link data, while the frontend gives users a dashboard to interact with their links.

## What does this project do?
This project lets users:
 - register/login
 - create custom short links pointing to long URLs
 - set an expiration date for the short links
 - track how many times each short link has been visited and from where, and how many unique visitors does it have
 - generate a QR code for each short link

## Key Libraries Used
1. Frontend

 - axios – Sends HTTP requests to the backend API (login, URLs, stats).
 - react-router-dom – Handles client-side routing between pages.
 - chart.js – Renders statistical charts for URL click data.
 - react-chartjs-2 – React wrapper used to display Chart.js charts.
2. Backend
 - gorm – ORM used to interact with the MySQL database.
 - gorm.io/driver/mysql – MySQL driver used by GORM.
 - github.com/golang-jwt/jwt/v5 – Handles JWT-based authentication.
 - golang.org/x/crypto – Used for secure password hashing.
 - github.com/gin-contrib/cors – Enables CORS for frontend–backend communication.
 - github.com/oschwald/geoip2-golang – Resolves visitor IPs to geographical locations.
 - github.com/skip2/go-qrcode – Generates QR codes for shortened URLs.

## How do I run it?
1. Clone the repo
 - git clone https://github.com/Lavinia-CristinaRus/urlShortener
 - cd urlShortener

!! As the database is not in cloud, you need to prepare your MySQL database and update the database config in the backend file backend\database\mysql.go before running !!

2. Backend Setup (Go)
Navigate to the backend folder:
 - cd backend

Install dependencies and run:
 - go mod tidy
 - go run main.go

3. Frontend Setup (React + Vite)
Open a new terminal and go to the frontend:
 - cd frontend/url-shortener

Install dependencies:
 - npm install

Start the development server:
 - npm run dev

Visit the URL shown in the terminal (e.g., http://localhost:5173/) to see the app.

## What should I expect to see when it works?

By accessing the link http://localhost:5173/, you should be able to see the login page
If everything is running correctly:

### Backend
A running API that handles:
 - user signup/login
 - creation of short URLs
 - URL update (short link and expiration date)
 - redirection from short to full URLs
 - statistics tracking (clicks) and retrieval
 - QR code generation

### Frontend
A React interface where you can:
 - register and log in
 - enter a long URL to shorten it
 - update the URL (short link and expiration date)
 - view a list of your shortened URLs
 - see the number of clicks for each link, unique visitors, statistics by city
 - obtain QR code from backend

## How to Use the Application (As user)
1. Sign In / Sign Up page
When you open the application, you are redirected to the Sign In page.

What to do:
 - Enter your Email and Password.
 - Click Sign In if you already have an account.
 - Click or sign up to navigate to the Sign Up page.
 - On the Sign Up page, enter the same fields and click Sign Up to create a new account.
 - After a successful sign-in or sign-up, you are redirected to My Short Links.

2. My Short Links page
This is the main dashboard of the application.

What you see:
 - A sign out option in the header.
 - A list of your existing short links (or a message if none exist).
 - A Create new short link button.

3. Create a new short link
To generate a new short link:

Steps:
 - Click Create new short link.
 - An input field labeled Url appears.
 - Enter the original long URL.
 - Click Generate short link to create it.
 - Click Cancel to close the form without creating a link.

After generation, the new short link appears in the list.

4. Use a short link

For each entry in the list, available information:
 - The original long URL (clickable).
 - The generated short URL (clickable).
 - Expiration date or message if not set

Clicking the short URL opens it in a new tab and redirects to the original URL.

5. Customize a short link
To customize an existing short link, expiration date and short link.

Steps:
 - Click the Customize button next to a short link.
 - You are redirected to the Customize URL page.
 - From here, you can modify settings related to that specific link.

6. View analytics
To view statistics for a short link.

Steps:
 - Click the Analytics button next to a short link.
 - You are redirected to the Analytics page.

Here you can see:
 - Click statistics
 - Visual charts representing usage data

7. Generate a QR code
To generate a QR code for a short link.

Steps:
 - Click the Generate QR button next to a short link.
 - You are redirected to the QR Code page.

A QR code is displayed for the selected short URL.

8. Sign out
To log out of the application.

Steps:
 - Click sign out in the header.

You are redirected back to the Sign In page.
