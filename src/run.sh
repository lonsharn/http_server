#!/bin/bash

go build http_server.go && cp http_server $1/ && cd $1 && ./http_server
