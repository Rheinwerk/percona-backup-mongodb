// This script is used to initialize the MongoDB database for the PBM (Percona Backup Manager)
// It creates a role and user for the PBM and sets up the necessary privileges.
//
// To run this script, you can use the following command:
//
// ```
// mongosh -f .devcontainer/db/pbm-init.js
// ```

// Create the PBM role
db.getSiblingDB("admin").createRole({ "role": "pbmAnyAction",
    "privileges": [
       { "resource": { "anyResource": true },
         "actions": [ "anyAction" ]
       }
    ],
    "roles": []
 });

// Create the PBM user
db.getSiblingDB("admin").createUser({user: "pbm",
     "pwd": "secret",
     "roles" : [
        { "db" : "admin", "role" : "readWrite", "collection": "" },
        { "db" : "admin", "role" : "backup" },
        { "db" : "admin", "role" : "clusterMonitor" },
        { "db" : "admin", "role" : "restore" },
        { "db" : "admin", "role" : "pbmAnyAction" }
     ]
  });
