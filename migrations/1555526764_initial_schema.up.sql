-- functions
-- DROP FUNCTION public.update_modified_at();

CREATE or REPLACE FUNCTION update_modified_at()
  RETURNS trigger
LANGUAGE 'plpgsql'
  COST 100
  VOLATILE NOT LEAKPROOF
AS $BODY$
BEGIN
  NEW.modified_at := NOW();
  RETURN NEW;
END;
$BODY$;

ALTER FUNCTION update_modified_at()
  OWNER TO vida;

-- SCHEMA: vida

-- DROP SCHEMA vida ;

CREATE SCHEMA IF NOT EXISTS vida
  AUTHORIZATION vida;

COMMENT ON SCHEMA vida
  IS 'vida represents the dev/production schema for the vida server';



CREATE TABLE vida.movies
(
  id serial NOT NULL,
  imdb_id character varying(10) unique,
  title text NOT NULL,
  synopsis text,
  image_url text,
  trailer_url text,
  playback_uri text,
  starring text,
  duration character varying(8),
  year integer NOT NULL,
  imdb_json json,
  search text,
  release_date timestamp without time zone,
  created_at timestamp without time zone  DEFAULT now(),
  updated_at timestamp without time zone  DEFAULT now()
)
  WITH (
    OIDS = FALSE
  );

ALTER TABLE vida.movies
  OWNER to vida;
COMMENT ON TABLE vida.movies
  IS 'movies is the table to hold all movies. ';

CREATE TRIGGER movies_modified_at_trigger
  BEFORE UPDATE
  ON vida.movies
  FOR EACH ROW
EXECUTE PROCEDURE update_modified_at();