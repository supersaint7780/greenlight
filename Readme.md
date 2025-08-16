# Greenlight API

In this project, we’re going to work through the start-to-finish build of an application called Greenlight — a JSON API for retrieving and managing information about movies. You can think of the core functionality as being a bit like the Open Movie Database API. This project is based on the book "Let's Go Further" by Alex Edwards, which I am following to develop this application and enhance my Go skills along the way.

## Endpoints and Actions

Our Greenlight API will support the following endpoints and actions:

| Method | URL Pattern | Action |
|--------|-------------|--------|
| GET    | /v1/healthcheck | Show application health and version information |
| GET    | /v1/movies | Show the details of all movies |
| POST   | /v1/movies | Create a new movie |
| GET    | /v1/movies/:id | Show the details of a specific movie |
| PATCH  | /v1/movies/:id | Update the details of a specific movie |
| DELETE | /v1/movies/:id | Delete a specific movie |
| POST   | /v1/users | Register a new user |
| PUT    | /v1/users/activated | Activate a specific user |
| PUT    | /v1/users/password | Update the password for a specific user |
| POST   | /v1/tokens/authentication | Generate a new authentication token |
| POST   | /v1/tokens/password-reset | Generate a new password-reset token |
| GET    | /debug/vars | Display application metrics |

## Learning Goals

- Develop a JSON API using Go.
- Implement various HTTP methods and endpoints.
- Manage user authentication and authorization.
- Enhance Go programming skills through practical application.

## Getting Started

To get started with the Greenlight API, follow the instructions in the book and refer to this README for endpoint details.

## Project Structure

```
.
├── bin
├── cmd
│   └── api
├── internal
├── migrations
├── remote
├── go.mod
└── Makefile
```

### Directory and File Overview

- **bin**: Contains the compiled application binaries, ready for deployment to a production server.
- **cmd/api**: Contains the application-specific code for the Greenlight API application, including the server code, HTTP request handling, and authentication management.
- **internal**: Contains various ancillary packages used by the API, such as database interaction, data validation, and email sending. This directory holds reusable code that is not application-specific.
- **migrations**: Contains the SQL migration files for the database.
- **remote**: Contains the configuration files and setup scripts for the production server.
- **go.mod**: Declares the project dependencies, versions, and module path.
- **Makefile**: Contains recipes for automating common administrative tasks, such as auditing Go code, building binaries, and executing database migrations.

### Postgress docker setup command
```sh
docker run --name greenlight_postgres -e POSTGRES_PASSWORD=password -p 5432:5432 -v postgres_data:/var/lib/postgresql/data -d postgres:latest
```

### DB Migration Issue
If you are using PostgreSQL v15, you may see error: 
pq: permission denied for schema public... when running the above command. This is because v15+ revokes CREATE from all users except a database owner. You can fix it with:
```sql
-- Set owner:
ALTER DATABASE greenlight OWNER TO greenlight;

-- And if that didn't work:
GRANT CREATE ON DATABASE greenlight TO greenlight;
```