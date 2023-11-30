CREATE TABLE
  public.movies (
    id serial NOT NULL,
    title character varying(255) NOT NULL,
    description text NULL,
    rating smallint NOT NULL DEFAULT 0,
    image character varying(255) NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now(),
    deleted_at timestamp without time zone NULL
  );

ALTER TABLE public.movies ADD CONSTRAINT movies_pkey PRIMARY KEY (id);
