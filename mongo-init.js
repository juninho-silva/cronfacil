db = db.getSiblingDB("cronfacil");

db.createCollection("logs");
db.logs.createIndex({ job_id: 1 })
db.logs.createIndex({ created_at: 1 }, { expireAfterSeconds: 604800 }) // 7 dias

db.createCollection("jobs");
db.jobs.createIndex({ name: 1 }, { unique: true })