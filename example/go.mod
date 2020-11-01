module github.com/uthng/protobuf-types/example

go 1.15

require (
	github.com/golang/protobuf v1.4.3
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.8.0
	github.com/spf13/cast v1.3.1 // indirect
	github.com/uthng/protobuf-types v0.0.0-20201031222440-d6367b69b2e3
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.25.0
)

replace github.com/uthng/protobuf-types => ../
