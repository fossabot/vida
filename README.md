# VIDA (WIP)
[![Build Status](https://travis-ci.com/gangachris/vida.svg?branch=master)](https://travis-ci.com/gangachris/vida)
[![codecov](https://codecov.io/gh/gangachris/vida/branch/master/graph/badge.svg)](https://codecov.io/gh/gangachris/vida)
[![Go Report Card](https://goreportcard.com/badge/github.com/gangachris/vida)](https://goreportcard.com/report/github.com/gangachris/vida)

Vida is a media server built with Golang and VueJS

## System Requirements
1. Docker and Docker Compose

## Run Everything
```bash
make start
```

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
