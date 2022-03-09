# relayer
**relayer is a high performance instant messaging server.** (in development)

- Custom Authentication System (CAS)
- Generic Database Access Layer (GDA)
- Supported Databases:
  - ~~MySQL [TBD]~~
  - PostgreSQL
  - ~~MongoDB [TBD]~~
  - ~~CockroachDB [TBD]~~

**Environment Variables**

| Variable  | Description | Default Value |
| ------------- | ------------- | ------------- |
| `RELAYER_DATABASE`  | Database Type<br />`cockroachdb`<br />`mongo`<br />`mysql`<br />`postgresql`  | `postgresql`  |
| `RELAYER_DB_CONNECTION_URL`  | Database Connection URL  | `postgres://postgres:postgres@localhost:5432/relayer`  |
| `RELAYER_API_KEY` | API key for relayer-server | auto generated on launch |
| `RELAYER_API_SECRET` | API secret for relayer-server | auto generated on launch |
