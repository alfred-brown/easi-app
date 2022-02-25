CREATE TABLE public.scheduled_action
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    scheduled_action_type text COLLATE pg_catalog."default" NOT NULL,
    action_data json,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    recurring boolean DEFAULT false,
    scheduled_date timestamp with time zone NOT NULL,
    days_to_next_occurence integer,
    last_success_at timestamp with time zone,
    current_attempts integer NOT NULL DEFAULT 0,
    max_attempts integer NOT NULL DEFAULT 5,
    scheduled_action_status text COLLATE pg_catalog."default" NOT NULL DEFAULT 'ready'::text,
    note text COLLATE pg_catalog."default" DEFAULT 'new scheduled action'::text,
    CONSTRAINT scheduled_action_pkey PRIMARY KEY (id)
)
