package subscriber

import "testing"

func TestFindBy(t *testing.T) {
	result, err := FindBy("name", "Kevin")

	if err != nil {
		t.Errorf("error")
	}

	if result.Name == "" {
		t.Errorf("error 2 " + result.Name)
	}
}