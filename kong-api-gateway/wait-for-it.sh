#!/usr/bin/env bash
set -e

HOST="$1"
shift
CMD="$@"

until pg_isready -h "$HOST"; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec $CMD