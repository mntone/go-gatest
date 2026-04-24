CREATE TABLE tasks(
	id INTEGER PRIMARY KEY AUTOINCREMENT
		CONSTRAINT ck_tasks_id
			CHECK (id BETWEEN 1 AND 9007199254740991),
	description TEXT NOT NULL
		CONSTRAINT ck_tasks_description
			CHECK (length(description) <= 255),
	created_at TEXT NOT NULL,
	updated_at TEXT NOT NULL
		CONSTRAINT ck_tasks_updated_at
			CHECK (updated_at >= created_at)
) STRICT;
