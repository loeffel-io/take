
package take

import (
	"crypto/tls"
	"github.com/getsentry/raven-go"
	"github.com/loeffel-io/helper"
	r "gopkg.in/gorethink/gorethink.v4"
	"log"
)

type Connection struct {
	Address    string
	Database   string
	Username   string
	Password   string
	AuthKey    string
	CertPath   string
	InitialCap int
	MaxOpen    int
}

func Connect(connection Connection) *r.Session {
	var connectOpts = r.ConnectOpts{
		Address:  connection.Address,
		Database: connection.Database,
		Username: connection.Username,
		Password: connection.Password,
	}

	if len(connection.AuthKey) > 0 {
		connectOpts.AuthKey = connection.AuthKey
	}

	if len(connection.CertPath) > 0 {
		connectOpts.TLSConfig = &tls.Config{
			RootCAs: helper.CertPool([]string{connection.CertPath}),
		}
	}

	if connection.InitialCap != 0 {
		connectOpts.InitialCap = connection.InitialCap
	}

	if connection.MaxOpen != 0 {
		connectOpts.MaxOpen = connection.MaxOpen
	}

	session, err := r.Connect(connectOpts)

	if err != nil {
		raven.CaptureError(err, nil)
		log.Fatal(err.Error())
	}

	return session
}
