package take

import (
	"github.com/getsentry/raven-go"
	r "gopkg.in/gorethink/gorethink.v4"
	"log"
)

// Delete data by a specific id
func Delete(table string, id interface{}, session *r.Session) {
	err := r.Table(table).Get(id).Delete().Exec(session)

	if err != nil {
		raven.CaptureError(err, nil)
		log.Fatal(err.Error())
	}
}
