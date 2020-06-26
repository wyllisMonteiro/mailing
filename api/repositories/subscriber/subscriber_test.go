package subscriber

import "testing"

func TestFindBy(t *testing.T) {
	result, err := FindBy("name", "Kevin")

	if err != nil {
		t.Errorf(err.Error())
	}

	if result.Name == "" {
		t.Errorf("error 2 " + result.Name)
	}
}