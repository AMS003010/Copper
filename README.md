# Copper Backend Pipeline

## 🚀 Introduction
**Copper** is a robust backend pipeline that integrates a PostgreSQL database, Redis caching, OpenTelemetry tracing, and monitoring with Prometheus and Grafana. The system is containerized using Docker and orchestrated with Docker Compose.

## 📦 Services

| Service      | Description |
|-------------|-------------|
| **app** | Main backend application (Go Gin server) |
| **nginx** | Reverse proxy and load balancer |
| **db** | PostgreSQL database |
| **redis** | Redis for caching |
| **prometheus** | Metrics collection and monitoring |
| **alertmanager** | Alerts for Prometheus |
| **grafana** | Visualization dashboard for metrics |
| **jaeger** | Distributed tracing |

## 🏗️ Architecture
The pipeline follows a **microservices-based** approach, with each component running in a dedicated container. The **Nginx** service acts as an entry point, directing traffic to the **app** service, which interacts with PostgreSQL, Redis, and monitoring tools.

![image](https://github.com/user-attachments/assets/e79e4324-9c66-45d6-9f44-c320e06822c5)

## 🔧 Setup & Usage
### 1️⃣ Prerequisites
Ensure you have the following installed:
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### 2️⃣ Clone the Repository
```sh
git clone https://github.com/AMS003010/Copper.git
cd Copper
```

### 3️⃣ Start the Services
Run the following command to build and start all services:
```sh
docker-compose up --build
```

### 4️⃣ Stop the Services
To stop the running containers:
```sh
docker-compose down
```

### 5️⃣ Persistent Data Storage
The PostgreSQL database and Grafana data are stored in Docker volumes:
- **copper_data** → PostgreSQL
- **grafana_data** → Grafana dashboards

## 📡 Endpoints & Ports
| Service | Port |
|---------|------|
| **App (Go Backend)** | `8080` |
| **Nginx Proxy** | `80` |
| **PostgreSQL** | `5432` |
| **Redis** | `6379` |
| **Prometheus** | `9090` |
| **Alertmanager** | `9093` |
| **Grafana** | `3000` |
| **Jaeger UI** | `16686` |

## 📊 Monitoring & Logging
- **Jaeger UI** (Tracing) → [`http://localhost:16686`](http://localhost:16686)
- **Prometheus Metrics** → [`http://localhost:9090`](http://localhost:9090)
- **Grafana Dashboard** → [`http://localhost:3000`](http://localhost:3000) (Default login: `admin/admin`)

## 🛠️ Troubleshooting
### Database Connection Issues
If the app cannot connect to the database, ensure PostgreSQL is ready:
```sh
docker-compose logs db
```

### Restarting Specific Services
To restart a specific service, run:
```sh
docker-compose restart <service_name>
```
Example:
```sh
docker-compose restart db
```

## 👥 Contributors
- **@AMS003010** *(Project Maintainer)*

## 📜 License
This project is licensed under the MIT License.

---
⚡ *Copper Backend Pipeline - Scalable & Monitored* ⚡

