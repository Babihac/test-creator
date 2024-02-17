CREATE TABLE IF NOT EXISTS public.answer (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	value varchar NOT NULL,
	question_id uuid NOT NULL,
	correct bool NOT NULL,
	CONSTRAINT answer_pk PRIMARY KEY (id),
	CONSTRAINT answer_question_fk FOREIGN KEY (question_id) REFERENCES public.question(id) ON DELETE CASCADE
);
