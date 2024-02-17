ALTER TABLE public.test ADD CONSTRAINT test_user_fk FOREIGN KEY (teacher_id) REFERENCES public."user"(id) ON DELETE CASCADE;
