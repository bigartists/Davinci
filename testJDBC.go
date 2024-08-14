package main

import (
	"fmt"
	"regexp"
)

func getDBTypeFromJdbcUrl(jdbcUrl string) (string, error) {
	re := regexp.MustCompile(`jdbc:([^:]+)://`)
	matches := re.FindStringSubmatch(jdbcUrl)
	if len(matches) < 2 {
		return "", fmt.Errorf("invalid JDBC URL: %s", jdbcUrl)
	}
	return matches[1], nil
}

// getIP extracts the IP address from the JDBC URL
func getIP(jdbcUrl string) (string, error) {
	re := regexp.MustCompile(`jdbc:[^:]+://([^:/]+)`)
	matches := re.FindStringSubmatch(jdbcUrl)
	if len(matches) < 2 {
		return "", fmt.Errorf("invalid JDBC URL: %s", jdbcUrl)
	}
	return matches[1], nil
}

// getPort extracts the port from the JDBC URL
func getPort(jdbcUrl string) (string, error) {
	re := regexp.MustCompile(`jdbc:[^:]+://[^:/]+:(\d+)`)
	matches := re.FindStringSubmatch(jdbcUrl)
	if len(matches) < 2 {
		// No port specified, provide default port based on DB type
		dbType, err := getDBTypeFromJdbcUrl(jdbcUrl)
		if err != nil {
			return "", err
		}
		switch dbType {
		case "mysql":
			return "3306", nil
		case "postgresql":
			return "5432", nil
		// Add other database types and their default ports as needed
		default:
			return "", fmt.Errorf("port not specified in JDBC URL and default port for %s is unknown", dbType)
		}
	}
	return matches[1], nil
}

func main() {
	jdbcUrl := "jdbc:mysql://100.43.2.58:3306"

	dbType, err := getDBTypeFromJdbcUrl(jdbcUrl)
	if err != nil {
		fmt.Println("Error getting DB type:", err)
	} else {
		fmt.Println("DB Type:", dbType)
	}

	ip, err := getIP(jdbcUrl)
	if err != nil {
		fmt.Println("Error getting IP:", err)
	} else {
		fmt.Println("IP:", ip)
	}

	port, err := getPort(jdbcUrl)
	if err != nil {
		fmt.Println("Error getting port:", err)
	} else {
		fmt.Println("Port:", port)
	}
}
