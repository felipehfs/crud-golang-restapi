package test

import (
	"testing"

	"felipe.com/cadastroUsuario/src/model"
)

// Testdatabase test if the connection is well
func TestDatabase(t *testing.T) {
	s, err := model.NewSession()
	if err != nil {
		t.Error(err)
	}
	defer s.Close()
}
