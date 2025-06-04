#!/bin/bash
set -xeuo pipefail

# wait till mongod is up
while ! mongo --eval 'db.adminCommand("ping")' &>/dev/null; do
    echo "Waiting for MongoDB to start..."
    sleep 1
done

echo "Initializing Replica Set..."
mongo --eval 'rs.initiate()'

echo "Creating PBM user..."
mongo /etc/pbm-agent/init-user.js

echo "Loading PBM config..."
pbm config --file /etc/pbm-agent/config.yml
