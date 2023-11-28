--
-- PostgreSQL database dump
--

-- Dumped from database version 14.9 (Ubuntu 14.9-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.9 (Ubuntu 14.9-0ubuntu0.22.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: delete_all_foundrisings_of_found(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.delete_all_foundrisings_of_found() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
begin
		DELETE FROM foundrising_tab WHERE found_id = OLD.id;
		IF NOT FOUND THEN 
			RETURN NULL; 
		END IF;
		DELETE FROM transaction_tab 
		WHERE from_id = OLD.id AND from_essence_type = true;
		IF NOT FOUND THEN 
			RETURN NULL; 
		END IF;
		DELETE FROM transaction_tab 
		WHERE to_id = OLD.id AND to_essence_type = false;
		IF NOT FOUND THEN 
			RETURN NULL; 
		END IF;
	RETURN OLD;
end;
$$;


ALTER FUNCTION public.delete_all_foundrisings_of_found() OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: foundation_tab; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.foundation_tab (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    password character varying(100) NOT NULL,
    cur_foudrising_amount integer,
    closed_foundrising_amount integer,
    fund_balance real,
    income_history real,
    outcome_history real,
    volunteer_amount integer,
    country character varying(100),
    login character varying(100) NOT NULL
);


ALTER TABLE public.foundation_tab OWNER TO postgres;

--
-- Name: foundrising_tab; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.foundrising_tab (
    id integer NOT NULL,
    found_id integer NOT NULL,
    description text,
    required_sum real,
    current_sum real,
    philantrops_amount integer,
    creation_date date,
    closing_date date
);


ALTER TABLE public.foundrising_tab OWNER TO postgres;

--
-- Name: transaction_tab; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transaction_tab (
    id integer NOT NULL,
    from_essence_type boolean,
    from_id integer NOT NULL,
    to_essence_type boolean,
    sum_of_money real,
    comment text,
    to_id integer NOT NULL
);


ALTER TABLE public.transaction_tab OWNER TO postgres;

--
-- Name: user_tab; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_tab (
    id integer NOT NULL,
    login character varying(100) NOT NULL,
    password character varying(100) NOT NULL,
    balance real,
    charity_sum real,
    closed_fing_amount integer
);


ALTER TABLE public.user_tab OWNER TO postgres;

--
-- Name: foundation_tab foundation_tab_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.foundation_tab
    ADD CONSTRAINT foundation_tab_pkey PRIMARY KEY (id);


--
-- Name: foundrising_tab foundrising_tab_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.foundrising_tab
    ADD CONSTRAINT foundrising_tab_pkey PRIMARY KEY (id);


--
-- Name: transaction_tab transaction_tab_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction_tab
    ADD CONSTRAINT transaction_tab_pkey PRIMARY KEY (id);


--
-- Name: user_tab user_tab_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_tab
    ADD CONSTRAINT user_tab_pkey PRIMARY KEY (id);


--
-- Name: foundation_tab foundrisings_deleting; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER foundrisings_deleting AFTER DELETE ON public.foundation_tab FOR EACH ROW EXECUTE FUNCTION public.delete_all_foundrisings_of_found();


--
-- PostgreSQL database dump complete
--

