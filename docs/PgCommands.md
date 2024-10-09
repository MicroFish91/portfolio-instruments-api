# PG Commands

This project includes `make pg-dump` and `make pg-restore` commands that utilize the environment variable files `.env.pgdump` and `.env.pgrestore` for efficiently dumping and restoring database data, especially across databases. A dump should be run first, and then the output dump file should be added to the restore environment variables before finally running the restore.  Example environment files are provided in the `env/` folder at the project root.

The project uses the `github.com/habx/pg-commands` library. Due to the way roles are handled in the library, a minor modification was made to the vendored files to ensure the CLI outputs the correct commands.