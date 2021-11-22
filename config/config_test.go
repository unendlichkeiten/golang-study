package config

import "testing"

func TestConfig(t *testing.T) {
	// test new config
	config, err := NewConfig("config.yaml")
	if err != nil {
		t.Errorf("new config err: %s", err)
		return
	}

	// set settings
	if err := config.ReadConfig("Server", &ServerSetting); err != nil {
		t.Errorf("read server config error: %s", err)
		return
	}
	if err := config.ReadConfig("App", &AppSetting); err != nil {
		t.Errorf("read app config error: %s", err)
		return
	}
	if err := config.ReadConfig("Database", &DatabaseSetting); err != nil {
		t.Errorf("read database config error: %s", err)
		return
	}

	t.Logf("test config success %v", config)
}
