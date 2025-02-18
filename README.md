# Carveo Car Management API 🚗💨

**Carveo** is a high-performance, scalable REST API built with Go and Gin for effortless car and engine management.

## Features 🚀
- **🔐 Secure Authentication:** JWT-based login & profile management.
- **🔄 Refresh Token:** The application uses JWT-based authentication, allowing users to renew their access token with a refresh token without logging out.
- **🔑 Reset Password:** Users can securely reset their passwords via email verification.
- **🚗 Car Management:** Create, read, update, and delete cars with brand and ID search.
- **🔧 Engine Management:** Full CRUD operations for engines.
- **📊 Real-Time Monitoring:** Logs and analytics for system health.
- **⚡ Lightning Fast:** Optimized performance with Go & Gin.

## Quick Start 🏁

**Install:** Make sure you have Go, Docker, and the required Go packages (see code for details).
# Run with Docker

# Clone the repository
```bash
  git clone https://github.com/sanjaygupta972004/Carveo.git
```

Go to the project directory
```bash
  cd Carveo
```

```bash
docker compose up
```

# Run Locally

Clone the project

```bash
  git clone https://github.com/sanjaygupta972004/Carveo.git
```

Go to the project directory

```bash
  cd Carveo
```

Install dependencies

  ```bash
  go mod download
  ```

Start the server

```bash
  go run main.go
```

# 📖 API Documentation

Access interactive API docs: [Swagger UI](https://carveo.site/swagger/index.html)

# 🌍 Deployment

## AWS + Docker Compose + Nginx + SSL 🔒

- ☁️ **Cloud Hosted:** Deployed on AWS for scalability & reliability.
- 🐳 **Dockerized:** Multi-container setup for seamless management.
- ⚡ **Reverse Proxy:** Nginx ensures smooth request handling & load balancing.
- 🔒 **HTTPS Security:** Certbot-managed SSL for encrypted connections.

## Access the API 🌐
[**Carveo API Live**](https://carveo.site/health)

