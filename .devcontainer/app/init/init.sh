#!/bin/bash

mongosh -f pbm-init.js

ROOT="/workspaces/percona-backup-mongodb"
go run ${ROOT}/cmd/pbm/main.go config --file ${ROOT}/.devcontainer/app/init/pbm.yml --mongodb-uri 'mongodb://pbm:secret@127.0.0.1:27017/?authSource=admin'

# Start pbm-agent
# go run ${ROOT}/cmd/pbm-agent/main.go --mongodb-uri 'mongodb://pbm:secret@127.0.0.1:27017/?authSource=admin'

# Show pbm status
# go run ${ROOT}/cmd/pbm/main.go status --mongodb-uri 'mongodb://pbm:secret@127.0.0.1:27017/?authSource=admin'
