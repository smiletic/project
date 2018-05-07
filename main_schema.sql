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
	filled_test,
	test,
	examination,
	patient,
	doctor,
	administrator,
	nurse,
	role,
	specialty,
	system_user,
	employee,
	person
	CASCADE;

CREATE TABLE person (
    uid uuid NOT NULL DEFAULT uuid_generate_v1(),
	name text NOT NULL,
	surname text NOT NULL,
	JMBG text NOT NULL,
	date_of_birth date,
	address text,
	email text,
	CONSTRAINT person_uid_pkey PRIMARY KEY (uid)
);

ALTER TABLE person OWNER TO postgres;


CREATE TABLE patient (
    uid uuid NOT NULL DEFAULT uuid_generate_v1(),
	person_uid uuid NOT NULL,
	medical_record_id text NOT NULL,
	health_card_id text NOT NULL,
	health_card_valid_until date NOT NULL,
	CONSTRAINT patient_uid_pkey PRIMARY KEY (uid),
	CONSTRAINT fk_person FOREIGN KEY (person_uid)
		REFERENCES person (uid) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE
);

ALTER TABLE patient OWNER TO postgres;

CREATE TABLE employee (
    uid uuid NOT NULL DEFAULT uuid_generate_v1(),
	person_uid uuid NOT NULL,
	work_document_id text NOT NULL,
	role_id int,
	CONSTRAINT employee_uid_pkey PRIMARY KEY (uid),
	CONSTRAINT fk_person FOREIGN KEY (person_uid)
		REFERENCES person (uid) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE
);

ALTER TABLE employee OWNER TO postgres;


CREATE TABLE specialty (
	specialty_id int NOT NULL,
	specialty_name text NOT NULL,
	CONSTRAINT specialty_id_pkey PRIMARY KEY (specialty_id)
);

ALTER TABLE specialty OWNER TO postgres;

CREATE TABLE doctor (
    uid uuid NOT NULL DEFAULT uuid_generate_v1(),
	employee_uid uuid NOT NULL,
	license_number text NOT NULL,
	specialty_id int NOT NULL,
	CONSTRAINT doctor_uid_pkey PRIMARY KEY (uid),
	CONSTRAINT fk_employee FOREIGN KEY (employee_uid)
		REFERENCES employee (uid) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE,
	CONSTRAINT fk_specialty FOREIGN KEY (specialty_id)
		REFERENCES specialty (specialty_id) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE
);

ALTER TABLE doctor OWNER TO postgres;

CREATE TABLE administrator (
    uid uuid NOT NULL DEFAULT uuid_generate_v1(),
	employee_uid uuid NOT NULL,
	CONSTRAINT administrator_uid_pkey PRIMARY KEY (uid),
	CONSTRAINT fk_employee FOREIGN KEY (employee_uid)
		REFERENCES employee (uid) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE
);

ALTER TABLE administrator OWNER TO postgres;

CREATE TABLE nurse (
    uid uuid NOT NULL DEFAULT uuid_generate_v1(),
	employee_uid uuid NOT NULL,
	CONSTRAINT nurse_uid_pkey PRIMARY KEY (uid),
	CONSTRAINT fk_employee FOREIGN KEY (employee_uid)
		REFERENCES employee (uid) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE
);

ALTER TABLE nurse OWNER TO postgres;

CREATE TABLE role (
	role_id int NOT NULL,
	role_name text NOT NULL,
	CONSTRAINT role_id_pkey PRIMARY KEY (role_id)
);

ALTER TABLE role OWNER TO postgres;

CREATE TABLE system_user (
    uid uuid NOT NULL DEFAULT uuid_generate_v1(),
	employee_uid uuid NOT NULL,
	role int NOT NULL,
	username text NOT NULL,
	password text NOT NULL,
	CONSTRAINT system_user_uid_pkey PRIMARY KEY (uid),
	CONSTRAINT fk_employee FOREIGN KEY (employee_uid)
		REFERENCES employee (uid) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE,
	CONSTRAINT fk_role FOREIGN KEY (role)
		REFERENCES role (role_id) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE
);

ALTER TABLE system_user OWNER TO postgres;

CREATE TABLE examination (
    uid uuid NOT NULL DEFAULT uuid_generate_v1(),
	doctor_uid uuid NOT NULL,
	patient_uid uuid NOT NULL,
	examination_date date NOT NULL DEFAULT now(),
	CONSTRAINT examination_uid_pkey PRIMARY KEY (uid),
	CONSTRAINT fk_doctor FOREIGN KEY (doctor_uid)
		REFERENCES doctor (uid) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE,
	CONSTRAINT fk_patient FOREIGN KEY (patient_uid)
		REFERENCES patient (uid) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE
);

ALTER TABLE examination OWNER TO postgres;

CREATE TABLE test (
    uid uuid NOT NULL DEFAULT uuid_generate_v1(),
	specialty_id int NOT NULL,
	questions jsonb NOT NULL,
	CONSTRAINT test_uid_pkey PRIMARY KEY (uid),
	CONSTRAINT fk_specialty FOREIGN KEY (specialty_id)
		REFERENCES specialty (specialty_id) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE
);

ALTER TABLE test OWNER TO postgres;

CREATE TABLE filled_test (
    uid uuid NOT NULL DEFAULT uuid_generate_v1(),
	examination_uid uuid NOT NULL,
	test_uid uuid NOT NULL,
	answers jsonb NOT NULL,
	CONSTRAINT filled_test_uid_pkey PRIMARY KEY (uid),
	CONSTRAINT fk_examination FOREIGN KEY (examination_uid)
		REFERENCES examination (uid) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE,
	CONSTRAINT fk_test FOREIGN KEY (test_uid)
		REFERENCES test (uid) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE
);

ALTER TABLE filled_test OWNER TO postgres;
