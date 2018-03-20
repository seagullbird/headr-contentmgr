package config

var (
	// PORT is the serving port, this value is set during compile time
	PORT = "unset"
	// DEVDBHOST is the postgres host used in dev environment
	DEVDBHOST = "postgresql-postgresql"
	// DEVDBPORT is the postgres port used in dev environment
	DEVDBPORT = "5432"
	// DEVDBUSER is the postgres dbuser used in dev environment
	DEVDBUSER = "postgres"
	// DEVDBNAME is the postgres dbname used in dev environment
	DEVDBNAME = "postgres"
	// DEVDBPASSWORD is the postgres password used in dev environment
	DEVDBPASSWORD = "qBDXNlz276"
	// DEVDBSSLMODE is the postgres sslmode used in dev environment
	DEVDBSSLMODE = "disable"
	// Auth0Domain is the domain for my auth0 service
	Auth0Domain = "https://headr.auth0.com"
	// Auth0Audience is the identifier for Auth0 API 'headr-api'
	Auth0Audience = "https://api.headr.io"
)
