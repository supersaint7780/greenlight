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
