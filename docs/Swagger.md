## 📖 Swagger

The swagger tool I want to use is only supported through Fiber v2, so I will wait for v3 support to come out before adopting and generating the relevant docs.  In the mean time, see the below for a quick summary of the public routes available in this REST API. If you want to understand the inputs and outputs, you'll need to inspect the codebase under the `/api/types` directory as well as inspect the schemas for each service under `/api/services/<insert service name>`.

### Healthcheck
* `GET /ping`

### Auth
* `POST /api/v2/register`
* `POST /api/v2/login`
* `GET /api/v2/me`

### Users
* `GET /api/v2/users` (Admin only)
* `GET /api/v2/users/:id`
* `PUT /api/v2/users/:id/verify` (Admin only)
* `DEL /api/v2/users/:id`

### Benchmarks
* `POST /api/v2/benchmarks`
* `GET /api/v2/benchmarks`
* `GET /api/v2/benchmarks/:id`
* `PUT /api/v2/benchmarks/:id`
* `DEL /api/v2/benchmarks/:id`

### Accounts
* `POST /api/v2/accounts`
* `GET /api/v2/accounts`
* `GET /api/v2/accounts/:id`
* `PUT /api/v2/accounts/:id`
* `DEL /api/v2/accounts/:id`

### Holdings
* `POST /api/v2/holdings`
* `GET /api/v2/holdings`
* `GET /api/v2/holdings/:id`
* `PUT /api/v2/holdings/:id`
* `DEL /api/v2/holdings/:id`

### Snapshots
* `POST /api/v2/snapshots`
* `GET /api/v2/snapshots`
* `GET /api/v2/snapshots/:id`
* `GET /api/v2/snapshots/:id/rebalance`
* `PUT /api/v2/snapshots/:id`
* `PUT /api/v2/snapshots/:id/order`
* `DEL /api/v2/snapshots/:id`

### SnapshotValues
* `POST /api/v2/snapshots/:snap_id/values`
* `GET /api/v2/snapshots/:snap_id/values`
* `GET /api/v2/snapshots/:snap_id/values/:snap_val_id`
* `PUT /api/v2/snapshots/:snap_id/values/:snap_val_id`
* `DEL /api/v2/snapshots/:snap_id/values/:snap_val_id`