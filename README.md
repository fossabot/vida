# VIDA (WIP)
[![Build Status](https://travis-ci.com/gangachris/vida.svg?branch=master)](https://travis-ci.com/gangachris/vida)
[![codecov](https://codecov.io/gh/gangachris/vida/branch/master/graph/badge.svg)](https://codecov.io/gh/gangachris/vida)
[![Go Report Card](https://goreportcard.com/badge/github.com/gangachris/vida)](https://goreportcard.com/report/github.com/gangachris/vida)

Vida is a media server built with Golang and VueJS

## System Requirements
1. Docker and Docker Compose

## Run Everything


## Status
Still a WIP but the following works. Right now we can retrieve metadata about movies in a directory using IMDB

1. Make sure the databse is running
```bash
make db
```

2. Run Migrations
```bash
make migrate-up
```

3. Make sure you have a directory with movies with `movie-name.mp4` for the movie names. To test, you can generate fake movie files with
```bash
make data
```

4. Get movie metadata (if you have a directory with movies, you can pass it to the dir flag below)
```bash
go run main.go media search --type movie --dir ./data
```

5. Check your db, it will have metadata for your movies. (POSTGRES: 5432, user: vida, password: vida, database: vida)

...😜 😜 😜 😜 😜

## CLI
The following commands are available to use (before the UI is built)
```bash
# search for media in a specific directory and add the metadata to your database
# type can be "movie" or "series" depending on what is contained in the --dir passed
# only movie supported for now.
./vida media search --type="<type>" --dir="dir" 
```

## Developer Environment
In case you want to set this up for developments, you can use the following steps.

### 1. Run migrations
Migration files in Vida are just `.sql` files and in this case they have a dialect for `postgres`. To create new migration files
run
```bash
make migrate-create name="change_to_sql_schema"
```
This will create two new files `<timestamp>_change_to_sql_schema.up.sql` and `<timestampe>_change_to_sql_schema.up.sql` representing the up and down
migrations respectively.

To run the up migrations, run:
```bash
make migrate-up
```

To run the down migrations, run:
```bash
make migrate-down
```

To rollback the previous migration, run:
```bash
make migrate-rollback #should be able to capture steps as a parameter soon
```
