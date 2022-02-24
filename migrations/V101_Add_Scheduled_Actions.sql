CREATE TABLE IF NOT EXISTS public.scheduled_action
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    scheduled_action_type text COLLATE pg_catalog."default" NOT NULL,
    action_data json,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone,
    recurring boolean DEFAULT false,
    scheduled_date timestamp with time zone,
    days_to_next_occurence integer,
    last_success_at timestamp with time zone,
    CONSTRAINT scheduled_action_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
