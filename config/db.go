package config

import "os"

// Env config from env or dev defaults
func Env() map[string]string {
	m := map[string]string{
		"host":     os.Getenv("DB_HOST"),
		"user":     os.Getenv("DB_USER"),
		"password": os.Getenv("DB_PASSWORD"),
		"pool":     os.Getenv("DB_POOL"),
	}

	if m["host"] == "" {
		m["host"] = "localhost"
	}
	if m["user"] == "" {
		m["user"] = "case"
	}
	if m["password"] == "" {
		m["password"] = "password"
	}
	if m["pool"] == "" {
		m["pool"] = "30"
	}

	return m
}
