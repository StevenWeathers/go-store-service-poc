#!/bin/bash

bash cassandra
sleep 10
cqlsh -f /tmp/schema/store-services.cql
cqlsh -f /tmp/load/stores.cql

pkill -f CassandraDaemon
