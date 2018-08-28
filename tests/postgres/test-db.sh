#!/usr/bin/env sh


psql -c "create database test"

psql -c "create schema test"

psql test -c "CREATE TABLE test_table (
  time TIMESTAMP WITH TIME ZONE,
  value DOUBLE PRECISION ,
  PRIMARY KEY (time, value)
)"