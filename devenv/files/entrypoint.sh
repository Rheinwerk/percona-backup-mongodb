#!/bin/bash
set -euo pipefail

/entrypoint-original.sh "$@" &
MONGO_PID=$!

if [ ! -f /etc/pbm-agent/.pbm-initialized ]; then
    /etc/pbm-agent/init.sh | tee /etc/pbm-agent/init.log
    touch /etc/pbm-agent/.pbm-initialized
fi

trap 'echo "Caught signal, shutting down..."; kill $MONGO_PID; wait $MONGO_PID' SIGINT SIGTERM
wait $MONGO_PID

echo "mongod exited. Keeping container alive for pbm restore or manual restart..."
tail -f /dev/null
