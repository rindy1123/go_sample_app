#!/bin/sh

atlas migrate apply \
  --url postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=$DB_SSLMODE \
  --dir file://migrations && \
  /bin/server
