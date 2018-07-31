package take

import (
	"github.com/getsentry/raven-go"
	r "gopkg.in/gorethink/gorethink.v4"
	"log"
)

// TableExists check if table exists
func TableExists(table string, session *r.Session) bool {
	var exists bool
	res, err := r.TableList().Contains(table).Run(session)

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

// TablesCreate creates single or multiple database tables
func TablesCreate(tables []string, session *r.Session) {
	for _, table := range tables {
		if TableExists(table, session) {
			continue
		}

		err := r.TableCreate(table).Exec(session)

		if err != nil {
			raven.CaptureError(err, nil)
			log.Fatal(err.Error())
		}
	}
}
