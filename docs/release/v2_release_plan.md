# V2 Release Plan

A central document for recording plans for releasing the new v2 API for Portfolio Instruments.

## Preparations

- Update client extension to v2 to support new compatibility requirements. Test using local Docker build.

## Pre-Deployment

These are all run locally and can be quickly spun up using Docker / Docker Compose.

### Setup V1
1. Start a brand new local test project (v1).
1. Seed with mock data (v1).  Should include at least 5 accounts, 5 holdings, 3 benchmarks, 3 snapshots. One of each category should be deleted / deprecated.
1. Verify everything works with released client extension (v1).

### Test Local Rollback
Test intentionally breaking the app / db.
1. Backup local database.
1. Put database in a fail state.
1. Build / test / verify rollback scripts.

### Setup V2
1. Stop the project (v1).
1. Run SQL migrations (v2).
1. Start the project (v2).
1. Verify everything works as expected with local client extension (v2). 
1. Seed with mock data (v2).  Seed requirements TBD. Should test deletion / deprecation.
1. Verify everything works normally with client extension (v2).

## Deployment
1. Backup remote database.
1. Release portfolio-instruments-api (v2).
1. Stop the remote server.
1. Run SQL migration on remote DB (v2).
1. Deploy portfolio-instruments-api (v2) on the remote server + restart.
1. Release client extension (v2).
1. Run sanity checks.