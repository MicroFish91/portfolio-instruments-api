# Change Log
## v2.0.0
### Overview
This major release includes numerous quality of life improvements and architectural changes. The settings table has been removed, with its values migrated directly to the snapshot table for better data organization. A new feature allows storing customizable value orderings that persist across clients, ensuring consistent data presentation.

Due to breaking API changes, v1 endpoints are now deprecated. All API requests should use the new `/api/v2` base path.

### Added
* Add support for storing snapshot value ordering ([#54](https://github.com/MicroFish91/portfolio-instruments-api/pull/54))
* Add recommended rebalance threshold field for benchmarks ([#36](https://github.com/MicroFish91/portfolio-instruments-api/pull/36))
* Add an optional rebalance threshold field per snapshot ([#38](https://github.com/MicroFish91/portfolio-instruments-api/pull/38))
* Support momentum-based asset types ([#35](https://github.com/MicroFish91/portfolio-instruments-api/pull/35))

### Changed
* Remove separate settings table. Update rebalance logic and tests to reflect data coming from the new snapshot rebalance field ([#40](https://github.com/MicroFish91/portfolio-instruments-api/pull/40))

### Fixed
* Snapshot calculations should include deprecated items ([#56](https://github.com/MicroFish91/portfolio-instruments-api/pull/56))
* Don't allow duplicate accounts even if one of them is deprecated ([#49](https://github.com/MicroFish91/portfolio-instruments-api/pull/49))
* Don't allow duplicate holdings even if one of them is deprecated ([#47](https://github.com/MicroFish91/portfolio-instruments-api/pull/47))
* Don't allow duplicate benchmarks even if one of them is deprecated ([#46](https://github.com/MicroFish91/portfolio-instruments-api/pull/46))

### Engineering
* Improve docker-compose and build scripts for quickly bootstrapping a local db and app ([#51](https://github.com/MicroFish91/portfolio-instruments-api/pull/51))

## v1.0.1
### Engineering
* [[2d8559](https://github.com/MicroFish91/portfolio-instruments-api/commit/2d85594bddffd640e46ee8d6c380b57238442a33)] Update Go runtime version to 1.25.4 and fix test launch config
* Various Dependabot security updates for package dependencies

## v1.0.0
Official release.