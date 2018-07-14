package take

import (
	"log"
	r "gopkg.in/gorethink/gorethink.v4"
	"github.com/getsentry/raven-go"
)

func DatabaseExists(table string, session *r.Session) bool {
	var exists bool
	res, err := r.DBList().Contains(table).Run(session)

	if err != nil {
		raven.CaptureError(err, nil)
		log.Fatal(err.Error())
	}

	if res.One(&exists) != nil {
		raven.CaptureError(err, nil)
		log.Fatal(err.Error())
	}

	return exists
}

func DatabaseCreate(database string, session *r.Session) {
	if DatabaseExists(database, session) {
		return
	}

	err := r.DBCreate(database).Exec(session)

	if err != nil {
		raven.CaptureError(err, nil)
		log.Fatal(err.Error())
	}
}
