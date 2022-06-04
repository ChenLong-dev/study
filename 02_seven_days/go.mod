module seven_days

go 1.17

replace seven_days/gee_frame/gee => ./gee_frame/gee

replace seven_days/gee_cache => ./gee_cache

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/mattn/go-sqlite3 v1.14.13 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
