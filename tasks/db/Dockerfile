# TODO: add steps to create a new
# Docker container image based on `mysql`
# with your `schema.sql` file copied into
# /docker-entrypoint-initdb.d/schema.sql
# see https://drstearns.github.io/tutorials/godb/#secbootstrappingthedatabaseschema
FROM mysql:5.7
COPY schema.sql /docker-entrypoint-initdb.d/schema.sql
