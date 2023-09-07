# My Auth Service

A simple authentication and authorization service.

- Written in Golang, based on [gin web framework](https://github.com/gin-gonic/gin).
- Provided OpenAPI interfaces, following [swagger spec](https://swagger.io/specification/).

## Components

- User
  - ID: string, for now is the same as name
  - Name: string, the unique identifier of a user
  - Password: string, stored only hash in memory, never in plain text
- Role
  - ID: string, for now is the same as name
  - Name: string, the unique identifier of a role
- Auth
  - Linked a User to several Roles

All the operation on components are been logged for timestamp.

## API

API Specification: check [API Specification](api/openapi.json)

- Create User: POST /users
- Delete User: DELETE /users/{id}
- Create Role: POST /roles
- Delete Role: DELETE /roles/{id}
- Add Role to User: POST /users/{userID}/roles/{roleID}
- Authenticate: POST /auth/login
- Invalidate: POST /auth/logout
  - auth token in Header required 
- Check Role: GET /users/me/roles/{id}
  - auth token in Header required
- All Roles: GET /users/me/roles
  - auth token in Header required
