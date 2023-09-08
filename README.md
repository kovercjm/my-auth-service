# My Auth Service

A simple authentication and authorization service.

Provided OpenAPI interfaces, following [OpenAPI Specification](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.1.md).

## Domain

- User
  - ID: string, for now is the same as name
  - Name: string, the unique identifier of a user
  - Password: string, stored only hash in memory, never in plain text
- Role
  - ID: string, for now is the same as name
  - Name: string, the unique identifier of a role

## In memory database design

- Structured data, stored in hash map with key as ID, for better searching performance
  - User: {ID: string, Name: string, PasswordHash: string}
  - Role: {ID: string, Name: string}
  - UserRole: {UserID: string, RoleID: string}
- Trashed structured data, stored in slice, keeping all deleted structured data, with deleted timestamp
- Token Map, storing all signed tokens

## API

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

For detailed specification please check: [openapi.yaml](api/openapi.yaml)

## Dependencies

- Dependency Injection framework: [fx](https://github.com/uber-go/fx) 
- Web framework: [gin](https://github.com/gin-gonic/gin)
- Personal tool kit for Golang, like logger: [tool-go](https://github.com/kovercjm/tool-go) 
- JSON Web Token: [jwt](https://github.com/golang-jwt/jwt)
- Official error library: [errors](https://github.com/pkg/errors)
- Testing framework: [ginkgo](https://github.com/onsi/ginkgo/v2) + [gomega](https://github.com/onsi/gomega)
- Testing data generator: [gofakeit](https://github.com/brianvoe/gofakeit/v6)
- CLI application helper: [cobra](https://github.com/spf13/cobra)
