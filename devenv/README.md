# pbm-agent Initialization Scripts

This directory contains the files to bring up a development environment for PBM.

## Setup

### Bucket `pbm-development`
- Create a bucket `pbm-development` in S3.
- Create a user that has access to the bucket.

### .env
Create a file `.env` with the credentials created above.
```bash
AWS_ACCESS_KEY_ID=XXX
AWS_SECRET_ACCESS_KEY=XXX
```

### Start environment
```bash
docker-compose up -d
```

## Development
```bash
# enter container
❯ docker-compose exec aio bash
[root@68b54666906a workspace]>

# start the agent
[root@68b54666906a workspace]> go run ./cmd/pbm-agent

# In another terminal
# trigger a backup
[root@68b54666906a workspace]❯ go run ./cmd/pbm backup --type=physical
Starting backup '2025-05-20T14:07:43Z'....Backup '2025-05-20T14:07:43Z' to remote store 's3:///pbm-development/mongodb-backups'

# list restores
[root@68b54666906a workspace]❯ go run ./cmd/pbm list
Backup snapshots:
  2025-05-20T13:53:41Z <physical> [restore_to_time: 2025-05-20T13:53:43Z]
  2025-05-20T14:07:43Z <physical> [restore_to_time: 2025-05-20T14:07:45Z]

PITR <off>:

# restore a backup
[root@68b54666906a workspace]❯ go run ./cmd/pbm restore 2025-05-20T14:07:43Z
Starting restore 2025-05-20T14:09:13.248296796Z from '2025-05-20T14:07:43Z'..........................Restore of the snapshot from '2025-05-20T14:07:43Z' has started.
Check restore status with: pbm describe-restore 2025-05-20T14:09:13.248296796Z -c </path/to/pbm.conf.yaml>
No other pbm command is available while the restore is running!
```

## IDE
You can connect to the running container and debug the pbm-agent in the container.

In your IDE select `Dev-Container` -> `Attach to running container`
