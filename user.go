package take

import (
	"github.com/getsentry/raven-go"
	r "gopkg.in/gorethink/gorethink.v4"
	"log"
)

type User struct {
	ID          string          `gorethink:"id"`
	Password    string          `gorethink:"password"`
	Permissions UserPermissions `gorethink:"-"`
}

type UserPermissions struct {
	Read   bool
	Write  bool
	Config bool
}

func SetupUser(user User, session *r.Session) {
	err := r.DB(rethinkDatabase).Table(rethinkUserTable).Insert(user, r.InsertOpts{Conflict: "update"}).Exec(session)

	if err != nil {
		raven.CaptureError(err, nil)
		log.Fatal(err.Error())
	}

	err = r.DB(session.Database()).Grant(user.ID, map[string]bool{
		"read":   user.Permissions.Read,
		"write":  user.Permissions.Write,
		"config": user.Permissions.Config,
	}).Exec(session)

	if err != nil {
		raven.CaptureError(err, nil)
		log.Fatal(err.Error())
	}
}
