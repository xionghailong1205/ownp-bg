package config

import "testing"

func TestActiveDatabaseURLDevelopment(t *testing.T) {
	t.Setenv("APP_ENV", "development")
	t.Setenv("DATABASE_URL_DEV", "postgres://dev")
	t.Setenv("DATABASE_URL_PROD", "postgres://prod")

	cfg := Load()
	active, err := cfg.ActiveDatabaseURL()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if active != "postgres://dev" {
		t.Fatalf("expected development database url, got %s", active)
	}
}

func TestActiveDatabaseURLProduction(t *testing.T) {
	t.Setenv("APP_ENV", "production")
	t.Setenv("DATABASE_URL_PROD", "postgres://prod")

	cfg := Load()
	active, err := cfg.ActiveDatabaseURL()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if active != "postgres://prod" {
		t.Fatalf("expected production database url, got %s", active)
	}
}

func TestActiveDatabaseURLProductionRequiresURL(t *testing.T) {
	t.Setenv("APP_ENV", "production")
	t.Setenv("DATABASE_URL_PROD", "")

	cfg := Load()
	_, err := cfg.ActiveDatabaseURL()
	if err == nil {
		t.Fatal("expected error when production url is missing")
	}
}
