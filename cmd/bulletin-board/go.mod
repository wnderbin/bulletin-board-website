module main

go 1.24.0

replace main/handlers => ../../internal/handlers

replace handlers/database => ../../internal/database

require main/handlers v0.0.0-00010101000000-000000000000

require (
	github.com/mattn/go-sqlite3 v1.14.24 // indirect
	handlers/database v0.0.0-00010101000000-000000000000 // indirect
)
