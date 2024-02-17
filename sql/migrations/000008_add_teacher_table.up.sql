CREATE TABLE IF NOT EXISTS public."user" (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	first_name varchar NOT NULL,
	last_name varchar NOT NULL,
	email varchar NOT NULL,
	CONSTRAINT user_unique UNIQUE (email),
	CONSTRAINT user_pk PRIMARY KEY (id)
);
