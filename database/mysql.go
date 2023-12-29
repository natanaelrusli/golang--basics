package database

var connection string

func init() {
	connection = "mysql"
}

func GetDatabaseConnection() string {
	return connection
}
