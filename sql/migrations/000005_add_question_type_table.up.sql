CREATE TABLE IF NOT EXISTS public.question_type (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	"type" varchar NOT NULL,
	CONSTRAINT question_type_pk PRIMARY KEY (id)
);
