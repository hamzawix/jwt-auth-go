// database/database_test.go
package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDB(t *testing.T) {
	err := InitDatabase()
	assert.NoError(t, err)
}
