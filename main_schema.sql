SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;
CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA pg_catalog;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


SET search_path = public, pg_catalog;

-- Drop current version tables

DROP TABLE IF EXISTS 
	popunjen_test,
	test,
	pregled,
	pacijent,
	doktor,
	administrator,
	rola,
	specijalizacija,
	korisnik,
	zaposleni,
	osoba
	CASCADE;

CREATE TABLE osoba (
    uid uuid NOT NULL DEFAULT uuid_generate_v1(),
	ime text NOT NULL,
	prezime text NOT NULL,
	jmbg text NOT NULL,
	datum_rodjenja date,
	adresa text,
	email text,
	CONSTRAINT osoba_uid_pkey PRIMARY KEY (uid)
);

ALTER TABLE osoba OWNER TO postgres;


CREATE TABLE pacijent (
    uid uuid NOT NULL DEFAULT uuid_generate_v1(),
	osoba_uid uuid NOT NULL,
	broj_kartona text NOT NULL,
	broj_isprave text NOT NULL,
	isprava_vazi_do date NOT NULL,
	CONSTRAINT pacijent_suid_pkey PRIMARY KEY (uid),
	CONSTRAINT fk_osoba FOREIGN KEY (osoba_uid)
		REFERENCES osoba (uid) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE
);

ALTER TABLE pacijent OWNER TO postgres;

CREATE TABLE zaposleni (
    uid uuid NOT NULL DEFAULT uuid_generate_v1(),
	osoba_uid uuid NOT NULL,
	broj_radne_isprave text NOT NULL,
	CONSTRAINT zaposleni_uid_pkey PRIMARY KEY (uid),
	CONSTRAINT fk_osoba FOREIGN KEY (osoba_uid)
		REFERENCES osoba (uid) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE
);

ALTER TABLE zaposleni OWNER TO postgres;


CREATE TABLE specijalizacija (
	id_specijalizacije int NOT NULL,
	ime_speccijalizacije text NOT NULL,
	CONSTRAINT id_specijalizacije_pkey PRIMARY KEY (id_specijalizacije)
);

ALTER TABLE specijalizacija OWNER TO postgres;

CREATE TABLE doktor (
    uid uuid NOT NULL DEFAULT uuid_generate_v1(),
	zaposleni_uid uuid NOT NULL,
	broj_licensce text NOT NULL,
	specijalnost int NOT NULL,
	CONSTRAINT doktor_uid_pkey PRIMARY KEY (uid),
	CONSTRAINT fk_zaposleni FOREIGN KEY (zaposleni_uid)
		REFERENCES zaposleni (uid) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE,
	CONSTRAINT fk_specijalnost FOREIGN KEY (specijalnost)
		REFERENCES specijalizacija (id_specijalizacije) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE
);

ALTER TABLE doktor OWNER TO postgres;

CREATE TABLE administrator (
    uid uuid NOT NULL DEFAULT uuid_generate_v1(),
	zaposleni_uid uuid NOT NULL,
	CONSTRAINT administrator_uid_pkey PRIMARY KEY (uid),
	CONSTRAINT fk_zaposleni FOREIGN KEY (zaposleni_uid)
		REFERENCES zaposleni (uid) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE
);

ALTER TABLE administrator OWNER TO postgres;

CREATE TABLE rola (
	id_role int NOT NULL,
	ime_role text NOT NULL,
	CONSTRAINT id_role_pkey PRIMARY KEY (id_role)
);

ALTER TABLE rola OWNER TO postgres;

CREATE TABLE korisnik (
    uid uuid NOT NULL DEFAULT uuid_generate_v1(),
	zaposleni_uid uuid NOT NULL,
	rola int NOT NULL,
	username text NOT NULL,
	password text NOT NULL,
	CONSTRAINT korisnik_uid_pkey PRIMARY KEY (uid),
	CONSTRAINT fk_zaposleni FOREIGN KEY (zaposleni_uid)
		REFERENCES zaposleni (uid) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE,
	CONSTRAINT fk_rola FOREIGN KEY (rola)
		REFERENCES rola (id_role) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE
);

ALTER TABLE korisnik OWNER TO postgres;

CREATE TABLE pregled (
    uid uuid NOT NULL DEFAULT uuid_generate_v1(),
	doktor_uid uuid NOT NULL,
	pacijent_uid uuid NOT NULL,
	datum_pregleda date NOT NULL DEFAULT now(),
	CONSTRAINT pregled_uid_pkey PRIMARY KEY (uid),
	CONSTRAINT fk_doktor FOREIGN KEY (doktor_uid)
		REFERENCES doktor (uid) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE,
	CONSTRAINT fk_pacijent FOREIGN KEY (pacijent_uid)
		REFERENCES pacijent (uid) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE
);

ALTER TABLE pregled OWNER TO postgres;

CREATE TABLE test (
    uid uuid NOT NULL DEFAULT uuid_generate_v1(),
	pitanja jsonb NOT NULL,
	CONSTRAINT test_uid_pkey PRIMARY KEY (uid)
);

ALTER TABLE test OWNER TO postgres;

CREATE TABLE popunjen_test (
    uid uuid NOT NULL DEFAULT uuid_generate_v1(),
	pregled_uid uuid NOT NULL,
	test_uid uuid NOT NULL,
	odgovori jsonb NOT NULL,
	CONSTRAINT popunjen_test_uid_pkey PRIMARY KEY (uid),
	CONSTRAINT fk_pregled FOREIGN KEY (pregled_uid)
		REFERENCES pregled (uid) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE,
	CONSTRAINT fk_test FOREIGN KEY (test_uid)
		REFERENCES test (uid) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE
);

ALTER TABLE popunjen_test OWNER TO postgres;
