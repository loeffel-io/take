# take

Simple RethinkDB ORM based on [gorethink](https://github.com/GoRethink/gorethink)

## Examples

### Basic Connection

```go
package main

import (
	"os"
	r "gopkg.in/gorethink/gorethink.v4"
	"github.com/loeffel-io/take"
)

var databaseSession *r.Session

func init(){
    // Init database
    databaseSession = take.Connect(take.Connection{
        Address:    os.Getenv("RETHINKDB_ADDRESS") + ":" + os.Getenv("RETHINKDB_PORT"),
        Database:   os.Getenv("RETHINKDB_DATABASE"),
        Username:   os.Getenv("RETHINKDB_USERNAME"),
        Password:   os.Getenv("RETHINKDB_PASSWORD"),
    })
}
```

### Connection Pool

```go
package main

import (
	"os"
	r "gopkg.in/gorethink/gorethink.v4"
	"github.com/loeffel-io/take"
)

var databaseSession *r.Session

func init(){
    // Init database
    databaseSession = take.Connect(take.Connection{
        Address:    os.Getenv("RETHINKDB_ADDRESS") + ":" + os.Getenv("RETHINKDB_PORT"),
        Database:   os.Getenv("RETHINKDB_DATABASE"),
        Username:   os.Getenv("RETHINKDB_USERNAME"),
        Password:   os.Getenv("RETHINKDB_PASSWORD"),
        InitialCap: 10,
        MaxOpen:    10,
    })
}
```

### Secure Connection

e.g.: [compose.io](https://compose.io)

```go
package main

import (
	"os"
	r "gopkg.in/gorethink/gorethink.v4"
	"github.com/loeffel-io/take"
)

var databaseSession *r.Session

func init(){
    // Init database
    databaseSession = take.Connect(take.Connection{
        Address:    os.Getenv("RETHINKDB_ADDRESS") + ":" + os.Getenv("RETHINKDB_PORT"),
        Database:   os.Getenv("RETHINKDB_DATABASE"),
        Username:   os.Getenv("RETHINKDB_USERNAME"),
        Password:   os.Getenv("RETHINKDB_PASSWORD"),
        AuthKey:    os.Getenv("RETHINKDB_AUTHKEY"),
        CertPath:   os.Getenv("RETHINKDB_CERT_PATH"),
    })
}
```

## Sentry support

This package supports [sentry.io](https://sentry.io) real time error reporting.
More informations: [Sentry golang docs](https://docs.sentry.io/clients/go)

```go
package main

import "github.com/getsentry/raven-go"

func init() {
    // Setup sentry
    raven.SetDSN("https://<key>:<secret>@sentry.io/<project>")
    
    // Setup cronjobs ...
}
```