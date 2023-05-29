package postal

import (
	"testing"
)

func TestModuleName(t *testing.T) {
	if ProjectName() != "gopostal" {
		t.Errorf("Project name `%s` incorrect", ProjectName())
	}
}
