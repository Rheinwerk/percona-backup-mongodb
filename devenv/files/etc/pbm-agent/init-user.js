db.getSiblingDB("admin").createRole(
	{
		"role": "pbmAnyAction",
		"privileges": [
			{
				"resource": { "anyResource": true },
				"actions": [ "anyAction" ]
			}
		],
	"roles": []
	}
);

db.getSiblingDB("admin").createUser(
	{
		"user": "pbm",
		"pwd": "secret",
		"roles" : [
			{ "db" : "admin", "role" : "readWrite", "collection": "" },
			{ "db" : "admin", "role" : "backup" },
			{ "db" : "admin", "role" : "clusterMonitor" },
			{ "db" : "admin", "role" : "restore" },
			{ "db" : "admin", "role" : "pbmAnyAction" }
		]
	}
);
