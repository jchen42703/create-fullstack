package configs

type BackendBase int

const (
	Echo BackendBase = iota
	Express
	FastAPI
)

func (base BackendBase) String() string {
	return []string{"Echo", "Express", "FastAPI"}[base]
}

type SQL_DB_TYPE int

const (
	CockroachDB SQL_DB_TYPE = iota
	PostgreSQL
	MySQL
)

func (base SQL_DB_TYPE) String() string {
	return []string{"CockroachDB", "PostgreSQL", "MySQL"}[base]
}

type NO_SQL_DB_TYPE int

const (
	MongoDB NO_SQL_DB_TYPE = iota
	ScyllaDB
	Redis
	Cassandra
)

func (base NO_SQL_DB_TYPE) String() string {
	return []string{"MongoDB", "ScyllaDB", "Redis", "Cassandra"}[base]
}

type SQLDBOptions struct {
	DBType        SQL_DB_TYPE
	StartupScript string // Runs this script to create the table etc.
}

type BackendConfig struct {
	Base             BackendBase
	DesiredImageName string
	Databases        struct {
		// Assume one SQL database type.
		// TODO: consider RDS/Spanner support
		SQL *SQLDBOptions

		// Multiple NoSQL databases isn't uncommon.
		// I.e. MongoDB as your regular DB + Redis for caching
		NoSQL []NO_SQL_DB_TYPE
	}
}
