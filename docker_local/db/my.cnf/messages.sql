CREATE TABLE IF NOT EXISTS "messages" (
	"id" INTEGER, 
	"text" VARCHAR(255), 
	"created_at" DATETIME,
	PRIMARY KEY ("id")
	);
INSER INTO messages (id, text, created_at) VALUES (1, "test", "1000-01-01 00:00:00.000000")