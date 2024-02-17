CREATE TABLE IF NOT EXISTS public.question (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	question_type uuid NOT NULL,
	points int NOT NULL,
	"name" varchar NOT NULL,
	CONSTRAINT question_pk PRIMARY KEY (id)
);
CREATE INDEX IF NOT EXISTS question_name_idx ON public.question ("name");
