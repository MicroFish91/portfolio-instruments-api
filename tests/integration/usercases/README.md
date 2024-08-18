## User & Authentication Tests

1. POST /api/v1/register
    - 201 Created (successfully registered user)
    - 400 Bad Request (payload invalid)
    - 409 Conflict (user already exists) 

2. POST /api/v1/login
    - 201 Created (successfully created login session)
    - 400 Bad Request (payload format invalid)
    - 401 Unauthorized (bad password)
    - 404 Not Found (bad email) 

3. GET /api/v1/users/me
    - 200 Ok
    - 401 Unauthorized
    - 404 Not Found

4. GET /api/v1/users/:id
    - 200 Ok
    - 400 Bad Request (not a valid id param)
    - 401 Unauthorized (bad token)
    - 403 Forbidden (param id and token id mismatch)

## Order

1. Register
2. Login
3. GET /me
4. GET /user
5. GET /settings