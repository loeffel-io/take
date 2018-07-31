package take

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	connectionTests := []Connection{
		{
			Address:  "localhost:28015",
			Username: "admin",
			Password: "",
		},
		{
			Address:    "localhost:28015",
			Username:   "admin",
			Password:   "",
			InitialCap: 10,
			MaxOpen:    10,
		},
	}

	for _, connection := range connectionTests {
		session := Connect(connection)

		assert.Equal(t, true, session.IsConnected(), "Should be connected")
		session.Close()
	}
}
