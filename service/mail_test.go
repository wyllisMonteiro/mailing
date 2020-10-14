package service

import "testing"

// TestSendMailFail : Get error
func TestSendMailFail(t *testing.T) {
	err := SendMail("subject", "body")

	t.Run("Get error", func(t *testing.T) {
		if err == nil {
			t.Errorf("err = %v, want error", err)
		}
	})
}
