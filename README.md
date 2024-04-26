# Golang RSS Aggregator Service

Welcome to the Golang RSS Aggregator Service project! This repository serves as a practical exercise for those looking to enhance their Go programming skills by building a fully functional RESTful API. This project isn't just a CRUD application; it includes a continuous service worker that fetches data from various online sources, providing real-time insights into blog content across the internet.

## Project Features

- **RESTful API Development:** Build an API using the robust Go programming language.
- **Database Integration:** Utilize PostgreSQL for data management, complemented by tools such as SQLc, Goose, and pgAdmin for database migrations and administration.
- **Service Worker:** Implement a long-running background worker that periodically scrapes and aggregates data from multiple feeds.
- **Real-World Tools:** Learn to use real-world tools and libraries like `chi` for routing and `cors` for handling cross-origin requests.
- **Environment Management:** Use `godotenv` for loading environment variables from a `.env` file, streamlining configuration management for development and production environments.

## Setup and Running

1. **Environment Setup:**
   Ensure you have Go installed on your machine. You will also need PostgreSQL and the associated tools mentioned above.

2. **Dependencies:**
   Install all necessary Go dependencies:
   ```bash
   go get -u ./...
