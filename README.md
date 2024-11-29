# Staking API Service

> 2024-11-29 @DanteBartel

## 🌟 Overview

The Staking API Service is focused on serving information about dApps, general message passing (GMP) and staking transactions. These information can be utilised by Scalar Staking app and Scanner app for user to review and interact with.

## 🌟 Key Features

1. DApps
   - Serve information about dApps
   - CRUD methods for dApps
   - Serve information about Custodials
2. GMP
   - Serve information about GMP
3. Staking transactions
   - Serve information about staking transactions

## 🌟 Components

The application is structured with the following key components:

### Core Components

- **API Handlers**

  - `HealthCheck`: Health check endpoint for service monitoring
  - `StakingHandler`: Manages staking operations and queries

- **Service Layer**

  - `StakingService`: Business logic for staking operations
  - `FinalityProviderService`: Manages finality provider operations

- **Database Layer**
  - `PostgreSQL`: Primary data store using GORM

### Infrastructure Components

- **Configuration**

  - `ConfigLoader`: Handles YAML-based configuration
  - `EnvironmentLoader`: Manages environment variables
  - `FinalityProviderConfig`: Manages finality provider settings

- **Middleware**
  - `CORS`: Cross-Origin Resource Sharing handler
  - `Security`: Request security middleware
  - `Swagger`: API documentation middleware

### Directory Structure

```
src/
├── api/          # API handlers and routes
├── config/       # Configuration files
├── docs/         # API documentation
├── internal/     # Internal packages
│   ├── service/  # Logic services
│   ├── models/   # Data models
│   └── db/       # Database interfaces
└── cmd/          # Application entry points
```

The service uses the following key dependencies:

- Chi router for HTTP routing
- GORM for database operations
- Swagger for API documentation
- Prometheus for metrics
- Zerolog for logging

## 🌟 Getting Started

### Prerequisites

- Docker
- Go

### Installation

1. Clone the repository:

```bash
git clone https://github.com/scalarorg/scalar-api
```

2. Prepare the config

Edit the `config/config-local.yml` file to match your local environment.

3. Run the service:

```
make run-local
```

OR, you can run as a docker container

```
make start-xchains-api
```

4. Open your browser and navigate to `http://localhost:<port>` to see the api server running, where `<port>` is the port specified in the `config/config-local.yml` file.

## 🌟 Tests

The service only contains integration tests so far, run below:

```
make tests
```

## 🌟 Update Mocks

1. Make sure the interfaces such as the `DBClient`is up to date
2. Install `mockery`: https://vektra.github.io/mockery/latest/
3. Run `make generate-mock-interface`
