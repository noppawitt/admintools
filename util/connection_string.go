package util

import "fmt"

// GetConnectionString returns db connection string
func GetConnectionString(username string, password string, host string, port int, dbName string) string {
	return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&encrypt=disable",
		username,
		password,
		host,
		port,
		dbName,
	)
}
