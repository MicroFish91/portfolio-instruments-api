## User & Authentication Tests

1. POST /api/v1/register

    - Test for each status type
        - 201 Created (successfully registered user)
        - 400 Bad Request (payload invalid)
        - 409 Conflict (user already exists) 

2. POST /api/v1/login

    - Test for each status type
        - 201 Created (successfully created login session)
        - 400 Bad Request (payload format invalid)
        - 401 Unauthorized (bad password)
        - 404 Not Found (bad email) 

## Order

1. Register
2. Login
3. GET /me
4. GET /user
5. GET /settings