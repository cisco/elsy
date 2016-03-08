package template

import "testing"

func TestEmberTemplate(t *testing.T) {
	if _, err := Get("ember", false); err != nil {
		t.Error("expected ember template to be registered")
	}
}
