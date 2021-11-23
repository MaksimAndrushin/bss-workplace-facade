#!/bin/sh

goose -dir migrations \
  postgres "user=postgres password=postgres host=0.0.0.0 port=5433 database=bss_workplace_facade sslmode=disable" \
  up