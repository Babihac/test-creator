CREATE TABLE IF NOT EXISTS public.test (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	"name" varchar NOT NULL,
	author_id uuid NOT NULL,
	duration interval minute NOT NULL,
	max_points int NOT NULL,
	"date" timestamp with time zone NOT NULL,
	CONSTRAINT test_pk PRIMARY KEY (id)
);
CREATE INDEX test_name_idx ON public.test ("name");