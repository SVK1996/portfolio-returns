# Portfolio Returns Calculator

## Overview
The Portfolio Returns Calculator is a system designed to calculate the top 5 performers based on portfolio returns across different time periods (daily, monthly, yearly, lifetime). The system is modular, extensible, and built with scalability and maintainability in mind.

## Future Enhancements:

## Architecture
The system is divided into several microservices, each responsible for a specific part of the application:
- **Portfolio Service**: Manages portfolio data and calculations.
- **Order Service**: Manages order data and processing.
- **Stock Price Service**: Fetches and caches stock prices.
- **Return Calculator Service**: Calculates returns for different periods.

### Technology Stack
- **Programming Language**: Go
- **Database**: PostgreSQL
- **Caching**: Redis
- **Message Queue**: Kafka
- **Containerization**: Docker
- **Orchestration**: Kubernetes
- **API Gateway**: Kong
- **Observability**: Prometheus, Grafana, Jaeger, ELK Stack

## Setup Instructions

### Prerequisites
- Docker
- Kubernetes
- Kafka
- Redis
- PostgreSQL
- Kong
- Prometheus
- Grafana
- Jaeger
- ELK Stack

### Clone the Repository
```sh
git clone https://github.com/SVK1996/portfolio-returns.git
cd portfolio-returns
