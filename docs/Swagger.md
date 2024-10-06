## 📖 Swagger

The swagger tool I want to use is only supported through Fiber v2, so I will wait for v3 support to come out before adopting and generating the relevant docs.  In the mean time, see the below for a quick summary of the public routes available in this REST API. If you want to understand the inputs and outputs, you'll need to inspect the codebase under the `/api/types` directory as well as inspect the schemas for each service under `/api/services/<insert service name>`.

### Healthcheck
* `GET /ping`

### Auth
* `POST /api/v1/register`
* `POST /api/v1/login`
* `GET /api/v1/me`

### Users
* `GET /api/v1/users/:id`
* `PUT /api/v1/users/:id`
* `DEL /api/v1/users/:id`

* `GET /api/v1/users/:id/settings`
* `PUT /api/v1/users/:id/settings`

### Benchmarks
* `POST /api/v1/benchmarks`
* `GET /api/v1/benchmarks`
* `GET /api/v1/benchmarks/:id`
* `PUT /api/v1/benchmarks/:id`
* `DEL /api/v1/benchmarks/:id`

### Accounts
* `POST /api/v1/accounts`
* `GET /api/v1/accounts`
* `GET /api/v1/accounts/:id`
* `PUT /api/v1/accounts/:id`
* `DEL /api/v1/accounts/:id`

### Holdings
* `POST /api/v1/accounts`
* `GET /api/v1/accounts`
* `GET /api/v1/accounts/:id`
* `PUT /api/v1/accounts/:id`
* `DEL /api/v1/accounts/:id`

### Snapshots
* `POST /api/v1/snapshots`
* `GET /api/v1/snapshots`
* `GET /api/v1/snapshots/:id`
* `GET /api/v1/snapshots/:id/rebalance`
* `PUT /api/v1/snapshots/:id`
* `DEL /api/v1/snapshots/:id`

### SnapshotValues
* `POST /api/v1/snapshots/:s_id/values`
* `GET /api/v1/snapshots/:s_id/values`
* `GET /api/v1/snapshots/:s_id/values/:sv_id`
* `PUT /api/v1/snapshots/:s_id/values/:sv_id`
* `DEL /api/v1/snapshots/:s_id/values/:sv_id`