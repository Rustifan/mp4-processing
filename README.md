# MP4 Processing System

This project is a distributed system for processing MP4 files. It consists of multiple services working together to handle file processing, storage, and communication. The system is designed to be platform-agnostic and easy to set up using Docker.

---

## Prerequisites

### Docker Installation

To run this project, you need Docker installed on your machine. Follow the official Docker installation guide for your operating system:

- [Install Docker](https://docs.docker.com/get-started/get-docker/)

After installation, verify Docker is running by executing the following command in your terminal:

```bash
docker --version
```

# Development Setup

1. Clone the repository:

```bash
git clone https://github.com/Rustifan/mp4-processing.git
```

2. Navigate to the project root directory:

```bash
cd mp4-processing
```

3. Run the following command to start all services:

```bash
docker compose up
```

4. Access the services:
   - **API Service**: Available at `http://localhost:${API_PORT}`.
   - **Swagger Documentation**: Available at `http://localhost:${API_PORT}/swagger/`.

# Services Architecture

## Database (PostgreSQL)

- **Container name**: mp4_db
- **Image**: postgres:17
- **Purpose**: Stores file metadata and processing status
- **Persistence**: Data persisted via Docker volume (postgres_data)

## Message Broker (NATS)

- **Container name**: mp4_nats
- **Purpose**: Handles asynchronous communication between services
- **Features**: Enables event-driven architecture for file processing

## API Service

- **Container name**: mp4_api
- **Build**: Custom build from ./api-service
- **Purpose**: Provides REST endpoints for file management
- **Features**:
  - Exposes HTTP API for client applications
  - Provides Swagger documentation
  - Connects to database and message broker

## API Subscriber

- **Container name**: mp4_api_subscriber
- **Build**: Shares codebase with API service
- **Purpose**: Listens for events from the message broker
- **Features**:
  - Updates database based on processing status events
  - Handles asynchronous operations

## Processing Service

- **Container name**: mp4_processor
- **Build**: Custom build from ./processing-service
- **Purpose**: Handles actual MP4 file processing

# API Endpoints

## File Management

- **Get All Files**:
  - **GET** `/v1/files/`
  - Retrieves a list of all files being processed.
- **Start Processing**:
  - **POST** `/v1/files/start-processing`
  - Submits an MP4 file for processing. Ensure the file is placed in the `/files` folder before making this request.
- **Delete File**:
  - **DELETE** `/v1/files/{id}`
  - Removes the file information from the database **and** deletes the processed file from the `/processed-files` directory.

# Environment Configuration

The `.env` file in the root directory contains environment variables for development:

```env
API_PORT=1235
DRIZLE_STUDIO_PORT=9999
```

- `API_PORT`: Port for the API service.
- `DRIZLE_STUDIO_PORT`: Port for Drizzle Studio (if applicable).

If you encounter port conflicts on your machine, you can modify these values in the `.env` file.

# Makefile Helpers

For easier interaction with the services, a `Makefile` is provided. Ensure you have `make` installed on your system to use these commands.

## Available Commands

- `api-exec`: Open a bash shell inside the API service container.

```bash
make api-exec
```

- `api-test`: Run tests inside the API service container.

```bash
make api-test
```

- `proc-test`: Run tests inside the processor service container.

```bash
make proc-test
```

- `db-connect`: Connect to the PostgreSQL database using `psql`.

```bash
make db-connect
```

- `db-push`: Push database schema changes.

```bash
make db-push
```

- `db-generate`: Generate database types.

```bash
make db-generate
```

- `db-studio`: Open Drizzle Studio in your browser.

```bash
make db-studio
```

- `nats-tools`: Run NATS tools inside a temporary container.

```bash
make nats-tools
```

# Folder Mounts

- `{project-root}/files`: Contains MP4 files to be processed. Mounted to `/files` inside the API and processing service containers.
- `{project-root}/processed-files`: Contains processed files (e.g., `{filename}_init.mp4`). Mounted to `/processed-files` inside the API and processing service containers.

# Production Considerations

While the current setup is designed for development, the following steps are recommended for production:

1. **CI/CD Pipeline**: Implement a CI/CD pipeline for automated testing and deployment.
2. **Kubernetes**: Deploy the services using Kubernetes for scalability and resilience.
3. **Security**: Add authentication and encryption for API and NATS communication.
4. **Monitoring**: Integrate monitoring tools (e.g., Prometheus, Grafana) for observability.

# License

This project is licensed under the MIT License. See the LICENSE file for details.
