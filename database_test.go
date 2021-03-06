package take

import (
	"github.com/stretchr/testify/assert"
	r "gopkg.in/gorethink/gorethink.v4"
	"testing"
)

func connectWithoutDatabase() *r.Session {
	connection := Connection{
		Address:  "localhost:28015",
		Username: "admin",
		Password: "",
	}

	return Connect(connection)
}

func connectWithDatabase() *r.Session {
	connection := Connection{
		Address:  "localhost:28015",
		Database: "test",
		Username: "admin",
		Password: "",
	}

	return Connect(connection)
}

func TestDatabaseCreate(t *testing.T) {
	// Connect without database
	session := connectWithoutDatabase()

	// Create test database
	DatabaseCreate("test", session)

	// Close database connection
	session.Close()

	// Connect to database
	session = connectWithDatabase()

	// Assert
	assert.Equal(t, true, session.IsConnected(), "Should be connected to specific database")
	assert.Equal(t, "test", session.Database(), "Database name should be test")

	// Close the session/connection
	session.Close()
}
