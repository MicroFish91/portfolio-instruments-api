# AZD Deep Discovery & Analysis Report

## File System Inventory
- Project contains Go source files (>50), Dockerfiles, Docker Compose configs, migrations, SQL artifacts, environment and configuration folders, documentation (README.md, LICENSE, CHANGELOG.md).

## Component Classification Table
| Component        | Type          | Technology      | Location        | Purpose                    |
|------------------|--------------|-----------------|-----------------|----------------------------|
| API Service      | API Web App  | Go, Fiber, JWT, Postgres | cmd/api/main.go, api/ | Main REST API with authentication, rate limiting, DB access
| Migration        | Background Service | Go, golang-migrate, Postgres | cmd/migrate/main.go, migrations/ | Database schema migration utility (local utility, not Azure hosted)
| PgCommands       | Supporting Tool | Go, habx/pg-commands, Postgres | cmd/pg_commands/main.go | Database dump/restore utility (local utility, not Azure hosted)
| TC Redact        | Utility      | Go, JSON        | cmd/tc_redact/main.go | Token redaction for ThunderClient request collections (local utility, not Azure hosted)
| Database         | SQL Service  | Postgres        | db/, migrations/, docker-compose.app.yml | Persistent data store for application

## Dependency Map
- All main services depend on centralized config files (config/), environment variables, and shared logger (logger/)
- API and migration utilities depend on the Postgres database via Docker Compose networking and environment settings
- PgCommands utility interacts with Postgres for backup/restore
- TC Redact utility interacts with ThunderClient collection files

## External Dependencies
- Postgres SQL database (hosted via Docker Compose as 'db' service)
- Configuration values provided via environment variables (Docker Compose)
- ThunderClient for REST API testing

## Communication Patterns
- Synchronous HTTP for REST API (Go Fiber)
- Direct DB connection via Postgres client libraries
- Migration via file-based schema scripts
- CLI-based dump/restore of DB for backups

## Key Supporting Files
- go.mod/go.sum: Go dependencies and version
- Dockerfile/Dockerfile.migrate: Container build specs
- docker-compose.app.yml/docker-compose.db.yml: Service orchestration
- README.md: Documentation

## Summary of Actions Performed
- Scanned and inventoried all top-level and subdirectory files
- Classified main deployable components and supporting tools
- Mapped dependencies and service communication
- Catalogued external requirements and runtime dependencies
- Inspected relevant configuration and orchestration files
- Created this report for architecture planning phase

---

# Azure Architecture Plan

## Azure Service Mapping Table

| Component     | Current Tech     | Azure Service              | Rationale                                                                 |
|--------------|------------------|----------------------------|---------------------------------------------------------------------------|
| API Service   | Go (Fiber)       | Azure Container Apps       | Best fit for scalable REST API, Go runtime, auto-scaling, managed platform |
| Database      | Postgres         | Azure Database for PostgreSQL Flexible Server | Native Postgres compatibility, managed DB, high availability               |

## Hosting Strategy Summary
- Only the API Service will be containerized and hosted in Azure Container Apps for flexible scaling and managed environment.
- Database will be provisioned using Azure Database for PostgreSQL Flexible Server, ensuring native compatibility and managed operations.
- Migration, PgCommands, and TC Redact are local utilities and DO NOT require Azure hosting or deployment. These will remain local CLI tools for administrative/data processing tasks.

## Containerization Plan
- API Service will have its own container with a dedicated Dockerfile and be published to Azure Container Registry.
- Utility containers for migration, PgCommands, and TC Redact are NOT required for Azure and will not be included in Azure deployment.

## Data Storage & Messaging Strategy
- Azure Database for PostgreSQL will handle all persistent data storage.
- No explicit messaging middleware required; synchronous HTTP and direct DB access suffice for current architecture.

## Resource Group & Networking Plan
- All resources will be organized into a single resource group per environment (dev/stage/prod).
- Private networking using Azure Virtual Network for secure DB/API communications.
- Managed Identity for secure secrets access; Key Vault enabled for credentials.

## Integration Patterns
- HTTP for API communication; direct Postgres driver access for Go services.
- Administrative/utility jobs (migration, PgCommands, TC Redact) are managed locally and not integrated as Azure Container jobs.

## Infrastructure as Code File Checklist

Based on the selected Azure services, the following Bicep files need to be generated:

### Core Files (Always Required)
- [ ] ./infra/main.bicep - Primary deployment template
- [ ] ./infra/main.parameters.json - Parameter defaults
- [ ] ./infra/modules/monitoring.bicep - Observability stack

### Service-Specific Modules
- [ ] ./infra/modules/container-apps.bicep - For the API container
- [ ] ./infra/modules/database.bicep - For PostgreSQL database
- [ ] ./infra/modules/keyvault.bicep - For secrets management
- [ ] ./infra/modules/container-registry.bicep - For container image storage

Total files to generate: 7

## Docker File Generation Checklist

Based on the containerization strategy, the following Docker files need to be generated:

### Service Dockerfiles
- [ ] ./api/Dockerfile - For Go API service
- [ ] ./api/.dockerignore - Exclude unnecessary files from API container

Total Docker files to generate: 2
---

# AZD Configuration Summary (azure.yaml)

The following azure.yaml configuration was generated for AZD deployment:

- **Project Name**: portfolio-instruments-api
- **Services**:
  - `api`: Go REST API service (located in ./api)
    - Host: Azure Container Apps
    - Dockerfile: ../Dockerfile (relative to api/)
    - Language: go
- **Infrastructure as Code**:
  - Bicep main entrypoint: ./infra/main.bicep
  - Parameters: ./infra/main.parameters.json

The azure.yaml was validated against the AZD schema and is fully compliant.

This configuration ensures only the API service is deployed to Azure, while utilities remain local tools.
