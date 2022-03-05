# relayer

- Custom Authentication System (CAS)
- Generic Database Access Layer (GDA)
- Supported Databases:
  - MySQL
  - PostgreSQL
  - MongoDB
  - CockroachDB

#### Environment Variables

    RELAYER_DATABASE: Database Type (cockroachdb, mongo, mysql, postgresql) Default: postgresql
    RELAYER_DB_CONNECTION_URL: Database Connection URL (e.g. postgresql://user:password@localhost:5432/relayer)
