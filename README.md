# Member Service
[![auto test](https://github.com/iamsad5566/member_service_frame/actions/workflows/test.yml/badge.svg)](https://github.com/iamsad5566/member_service_frame/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/iamsad5566/member_service_frame)](https://goreportcard.com/report/github.com/iamsad5566/member_service_frame)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

[中文版](/doc/README-zh.md) , [English](README.md)   

Member Service is an open-source microservice developed in Go that handles user authentication, registration, and login. It integrates various use cases and scenarios related to user management.

## Features
- User Registration
- User Login
- User Authentication
- Password Update
- OAuth2.0 Integration (Google)
- gRPC Service for Authorization
- Automatic Database Creation
- Containerized Deployment

## Getting Started
### Prerequisites
- Go (version 1.16 or later)
- PostgreSQL
- Redis

### Installation
1. Clone the repository:
```
https://github.com/iamsad5566/member_service_frame.git
```
2. Navigate to the project directory:
```
cd member_service_frame
```
3. Build:
```
go build
```

### Configuration
The application uses a configuration file (`example_config.yml`) to store various settings. Here's an overview of the configuration options:
- **valid_login:** Number of days before a user's login expires (default: 14).
- **member_service:**
    - **host:** Host for the Member Service (censored in the example).
    - **port:** Port for the Member Service (default: 888).
    - **RESTfulPort:**  Port for the RESTful API (default: 8112).
    - **gRPCPort:** Port for the gRPC service (default: 8113).
- **jwt:**
    - **secret_key:** Secret key for JSON Web Tokens (censored in the example).
    - **token_expire:** Token expiration time in seconds (default: 86400, which is 24 hours).
- **db:**
    - **psql:**
        - **account:** PostgreSQL database account (censored in the example).
        - **password:** PostgreSQL database password (censored in the example).
        - **host:** PostgreSQL database host (censored in the example).
        - **port:** PostgreSQL database port (default: 5433).
        - **maxIdleConns:** Maximum number of idle connections in the PostgreSQL connection pool (default: 20).
        - **maxOpenConns:** Maximum number of open connections in the PostgreSQL connection pool (default: 20).
        - **maxLifeMinute:** Maximum life time (in minutes) for connections in the PostgreSQL connection pool (default: 10).
    - **redis:**
        - **password:** Redis password (censored in the example).
        - **host:** Redis host (censored in the example).
        - **port:** Redis port (default: 6379).
- **logConfig:**
    - **level:** Log level (default: info).
    - **filename:** Log file name (default: logs/viper_zap_gin.log).
    - **maxsize:**  Maximum size of the log file in megabytes (default: 1).
    - **max_age:**  Maximum age of log files in days (default: 30).
    - **max_backups:** Maximum number of old log files to keep (default: 5).

- **oauth2**
    - **google:** 
        - **client_id:**  Google OAuth2.0 client ID (censored in the example).
        - **client_secret:** Google OAuth2.0 client secret (censored in the example).

> [!NOTE] 
> Please note that the sensitive values (e.g., passwords, secrets) are censored in the provided example for security reasons.

#### Using the Open-Source Setting Implementation
If you prefer to use the open-source version of the setting implementation, you can follow these steps:
1. Rename the `example_config.yml` file to `config.yml` and modify the path in `config.getRoot()`.
2. Update the configuration values according to your environment.

#### Using the Private Setting Implementation (setconf)
Alternatively, you can use the private `setconf` repository (this repository only be applied by the author) for setting management. In this case, you'll need to replace the `config.Setting` object with your own implementation.

### Running the Application
After configuring the settings, you can run the application with the following command:
```
./member_service_frame
```
This will start the Member Service, which includes the RESTful API, gRPC service, and automatic database creation (if the database doesn't exist).

### API Documentation
The Member Service provides a Swagger UI for API documentation. You can access it at `http://localhost:8080/swagger/index.html` (replace `8080` with the configured port if different).

### Deployment
The Member Service can be deployed using Docker containers. One Dockerfile is provided:
1. **Dockerfile:** For building and running the application.

### Building the Docker Image
To build the Docker image, run the following command:
```
docker build --build-arg GITHUB_TOKEN=<your_github_token> --build-arg LATEST_SETCONF_VERSION=<setconf_version> -t member_service .
```

Replace `<your_github_token>` with your GitHub personal access token and `<setconf_version>` with the latest version of the setconf repository. By the way, if you decide to use the open-source implementation of the config setting, please remove everything related to `setconf` and `GOPRIVATE` in the `Dockerfile`, and `./github/workflows/ci-cd.yml`, `./github/workflows/test.yml`.

### Running the Docker Container
After building the Docker image, you can run the container with the following command:

```
docker run -d --name member_service -p 8080:8080 member_service
```

This will start the Member Service container and map port 8080 to the host.

### CI/CD
The Member Service includes GitHub Actions workflows for Continuous Integration (CI) and Continuous Deployment (CD).

#### CI Workflow
The CI workflow (`auto test`) is triggered on every push to the `main` branch. It performs the following steps:
1. Checks out the source code
2. Sets up the `GOPRIVATE` environment variable
3. Logs in to the GitHub Package Registry
4. Builds the test Docker image (passing the GitHub Token and `setconf` version as build arguments)
5. Runs the tests inside the Docker container

#### CD Workflow
The CD workflow (`CI/CD`) is triggered when a new release is published. It performs the following steps:
1. Checks out the source code
2. Sets up the `GOPRIVATE` environment variable
3. Logs in to the GitHub Package Registry and Docker Hub
4. Builds the Docker image (passing the GitHub Token and `setconf` version as build arguments)
5. Tags and pushes the Docker image to Docker Hub
6. Deploys the Docker image to a remote server via SSH

> [!NOTE] 
> Please note that you'll need to configure the required GitHub Secrets (TOKEN, LATEST_SETCONF_VERSION, REMOTE_HOST, SSH_PRIVATE_KEY, DOCKER_USERNAME, and DOCKER_PASSWORD) for the workflows to work correctly.

### Contributing
Contributions to the Member Service are welcome! If you find any issues or want to add new features, please open an issue or submit a pull request.

### License
The Member Service is open-source software released under the MIT License.

