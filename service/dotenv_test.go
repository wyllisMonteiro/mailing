package service

import (
	"testing"
)

// TestGoDotEnvVariableWorking : Get var env from env file occurs no error
func TestGoDotEnvVariableWorking(t *testing.T) {
	EnvFilePath = "../.env_test"

	usermail, err := GoDotEnvVariable("USERMAIL")

	t.Run("Get username", func(t *testing.T) {
		if err != nil {
			t.Errorf("err = %v, want nil", err)
		}

		if usermail != "myusername" {
			t.Errorf("usermail = %v, want myusername", usermail)
		}
	})
}

// TestGoDotEnvVariableBadFileppath : Get var env from env file occurs error because it doesn't exists
func TestGoDotEnvVariableBadFileppath(t *testing.T) {
	EnvFilePath = "./.env_test"

	_, err := GoDotEnvVariable("USERMAIL")

	t.Run("Get error", func(t *testing.T) {
		if err == nil {
			t.Errorf("err = %v, want an error", "nil")
		}
	})
}

// TestGoDotEnvVariableBadKey : Get an empy variable
func TestGoDotEnvVariableBadKey(t *testing.T) {
	EnvFilePath = "../.env_test"

	nothing, _ := GoDotEnvVariable("BADKEY")

	t.Run("Get error", func(t *testing.T) {
		if nothing != "" {
			t.Errorf("nothing = %v, want ''", nothing)
		}
	})
}

// TestGetVarsMailWorking : Get all env variables to send mail
func TestGetVarsMailWorking(t *testing.T) {
	EnvFilePath = "../.env_test"

	host, port, user, pass, from, to, err := GetVarsMail()

	t.Run("Get all vars for mail", func(t *testing.T) {
		if err != nil {
			t.Errorf("err = %v, want nil", err)
		}

		if host != "smtp.test.io" {
			t.Errorf("host = %v, want smtp.test.io", host)
		}

		if port != 587 {
			t.Errorf("port = %v, want 587", port)
		}

		if user != "myusername" {
			t.Errorf("user = %v, myusername", user)
		}

		if pass != "mypass" {
			t.Errorf("pass = %v, want mypass", pass)
		}

		if from != "me@test.io" {
			t.Errorf("from = %v, myusername", from)
		}

		if to != "me@test.io" {
			t.Errorf("to = %v, want mypass", to)
		}
	})
}

// TestGetVarsMailBadFilepath : Get var env from env file occurs error because it doesn't exists
func TestGetVarsMailBadFilepath(t *testing.T) {
	EnvFilePath = "./.env_test"

	_, _, _, _, _, _, err := GetVarsMail()

	t.Run("Get error", func(t *testing.T) {
		if err == nil {
			t.Errorf("err = %v, want error", err)
		}
	})
}
