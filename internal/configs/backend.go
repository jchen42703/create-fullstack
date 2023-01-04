package configs

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
	DBType SQL_DB_TYPE `yaml:"db_type"`
	// Runs this script to create the table etc.
	StartupScript string `yaml:"startup_script"`
}

type PreCommitOptions struct {
	Lint          bool `yaml:"lint"`
	FormatterOpts *struct {
		Formatter string
	} `yaml:"format"`
}

type APIAugmentationOptions struct {
	HuskyOpts     *HuskyOptions     `yaml:"husky"`
	PreCommitOpts *PreCommitOptions `yaml:"pre_commit"`
	AddDockerfile bool              `yaml:"dockerfile"`
	GitOpts       *struct {
		AddIssueTemplates bool `ymal:"issue_templates"`
		AddPRTemplates    bool `yaml:"pr_templates"`
	} `yaml:"git"`

	AddCI string `yaml:"ci"`
}

type BackendConfig struct {
	OutputDirectoryPath string                  `yaml:"output_dir"`
	Base                string                  `yaml:"base"`
	Language            string                  `yaml:"lang"`
	AugmentOpts         *APIAugmentationOptions `yaml:"augment"`
	Databases           struct {
		// Assume one SQL database type.
		// TODO: consider RDS/Spanner support
		SQL *SQLDBOptions `yaml:"sql"`

		// Multiple NoSQL databases isn't uncommon.
		// I.e. MongoDB as your regular DB + Redis for caching
		NoSQL []NO_SQL_DB_TYPE `yaml:"no_sql"`
	} `yaml:"db"`
}
