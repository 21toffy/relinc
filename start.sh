#!/bin/sh

set -e

echo "run migrations"

which /app/migrate
echo "source $DB_SOURCE"

/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start app"
exec "$@"