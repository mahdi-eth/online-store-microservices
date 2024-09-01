# OnlineStore-Microservices

## Overview
OnlineStore-Microservices is a demonstration of a microservices architecture for an online store. This project showcases how to build and deploy a set of microservices using Go, Kubernetes, Docker, and message broker NATS. 

## Features
- **Product Service**: Manages product information using PostgreSQL.
- **Order Service**: Handles order processing with MongoDB.
- **Inventory Service**: Tracks inventory levels with Redis.
- **Inter-service Communication**: Uses HTTP and message brokers for communication.
- **Deployment & Orchestration**: Managed using Kubernetes for scalable and reliable deployments.

## Technologies
- **Go**: Programming language used for microservices development.
- **Kubernetes**: Orchestration platform for containerized applications.
- **Docker**: Containerization for consistent deployment environments.
- **NATS**: Message brokers for asynchronous communication.
- **PostgreSQL**: Relational database for the Product Service.
- **MongoDB**: NoSQL database for the Order Service.
- **Redis**: In-memory data structure store for the Inventory Service.

## Getting Started

### Prerequisites
- Docker
- Kubernetes
- kubectl (Kubernetes CLI)
- Go (for building the microservices)

### Installation

1. **Clone the repository:**
    ```bash
    git clone https://github.com/mahdi-eth/online-store-microservices.git
    cd online-store-microservices
    ```

2. **Build Docker Images:**
    ```bash
        docker compose build
    ```

3. **Deploy to Kubernetes:**
    - Apply the Kubernetes manifests:
      ```bash
      kubectl apply -f kubernetes/pvcs/
      kubectl apply -f kubernetes/deployments/
      kubectl apply -f kubernetes/services/
      ```
