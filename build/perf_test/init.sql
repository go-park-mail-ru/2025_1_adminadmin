--
-- PostgreSQL database dump
--

-- Dumped from database version 16.2 (Debian 16.2-1.pgdg120+2)
-- Dumped by pg_dump version 16.2 (Debian 16.2-1.pgdg120+2)

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
-- Name: pg_trgm; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pg_trgm WITH SCHEMA public;


--
-- Name: EXTENSION pg_trgm; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pg_trgm IS 'text similarity measurement and index searching based on trigrams';


--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


--
-- Name: issue_promocode_on_paid(); Type: FUNCTION; Schema: public; Owner: admin_test
--

CREATE FUNCTION public.issue_promocode_on_paid() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE
    new_promocode TEXT := 'PROMO_' || substr(md5(random()::text), 0, 8);
BEGIN
    IF NEW.status = 'paid' AND OLD.status IS DISTINCT FROM NEW.status THEN
        INSERT INTO promocodes (promocode, discount, user_id, created_at, expires_at)
        VALUES (
            new_promocode,
            0.10,
            NEW.user_id,
            now(),
            now() + INTERVAL '7 days'
        );
    END IF;

    RETURN NEW;
END;
$$;


ALTER FUNCTION public.issue_promocode_on_paid() OWNER TO admin_test;

--
-- Name: make_product_tsvector(text, text); Type: FUNCTION; Schema: public; Owner: admin_test
--

CREATE FUNCTION public.make_product_tsvector(name text, category text) RETURNS tsvector
    LANGUAGE plpgsql IMMUTABLE
    AS $$
BEGIN
    RETURN (
        setweight(to_tsvector('ru', coalesce(name, '')), 'A') ||
        setweight(to_tsvector('ru', coalesce(category, '')), 'B')
    );
END;
$$;


ALTER FUNCTION public.make_product_tsvector(name text, category text) OWNER TO admin_test;

--
-- Name: make_restaurant_tsvector(text, text); Type: FUNCTION; Schema: public; Owner: admin_test
--

CREATE FUNCTION public.make_restaurant_tsvector(name text, description text) RETURNS tsvector
    LANGUAGE plpgsql IMMUTABLE
    AS $$
BEGIN
    RETURN (
        setweight(to_tsvector('ru', coalesce(name, '')), 'A') ||
        setweight(to_tsvector('ru', coalesce(description, '')), 'B')
    );
END;
$$;


ALTER FUNCTION public.make_restaurant_tsvector(name text, description text) OWNER TO admin_test;

--
-- Name: set_order_in_delivery(uuid); Type: FUNCTION; Schema: public; Owner: admin_test
--

CREATE FUNCTION public.set_order_in_delivery(order_id uuid) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE orders SET status = 'in_delivery', created_at = NOW() WHERE id = order_id;
END;
$$;


ALTER FUNCTION public.set_order_in_delivery(order_id uuid) OWNER TO admin_test;

--
-- Name: update_product_tsvector(); Type: FUNCTION; Schema: public; Owner: admin_test
--

CREATE FUNCTION public.update_product_tsvector() RETURNS trigger
    LANGUAGE plpgsql
    AS $$ 
BEGIN
    NEW.tsvector_column := make_product_tsvector(NEW.name, NEW.category);
    RETURN NEW;
END;
$$;


ALTER FUNCTION public.update_product_tsvector() OWNER TO admin_test;

--
-- Name: update_restaurant_rating(); Type: FUNCTION; Schema: public; Owner: admin_test
--

CREATE FUNCTION public.update_restaurant_rating() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE restaurants
    SET 
        rating = ROUND(
            COALESCE((rating * rating_count + NEW.rating) / (rating_count + 1), NEW.rating)::numeric,1)::float,
        rating_count = COALESCE(rating_count, 0) + 1
    WHERE id = NEW.restaurant_id;

    RETURN NEW;
END;
$$;


ALTER FUNCTION public.update_restaurant_rating() OWNER TO admin_test;

--
-- Name: update_restaurant_tsvector(); Type: FUNCTION; Schema: public; Owner: admin_test
--

CREATE FUNCTION public.update_restaurant_tsvector() RETURNS trigger
    LANGUAGE plpgsql
    AS $$ 
BEGIN
    NEW.tsvector_column := make_restaurant_tsvector(NEW.name, NEW.description);
    RETURN NEW;
END;
$$;


ALTER FUNCTION public.update_restaurant_tsvector() OWNER TO admin_test;

--
-- Name: russian_ispell; Type: TEXT SEARCH DICTIONARY; Schema: public; Owner: admin_test
--

CREATE TEXT SEARCH DICTIONARY public.russian_ispell (
    TEMPLATE = pg_catalog.ispell,
    dictfile = 'russian', afffile = 'russian', stopwords = 'russian' );


ALTER TEXT SEARCH DICTIONARY public.russian_ispell OWNER TO admin_test;

--
-- Name: ru; Type: TEXT SEARCH CONFIGURATION; Schema: public; Owner: admin_test
--

CREATE TEXT SEARCH CONFIGURATION public.ru (
    PARSER = pg_catalog."default" );

ALTER TEXT SEARCH CONFIGURATION public.ru
    ADD MAPPING FOR asciiword WITH english_stem;

ALTER TEXT SEARCH CONFIGURATION public.ru
    ADD MAPPING FOR word WITH public.russian_ispell, russian_stem;

ALTER TEXT SEARCH CONFIGURATION public.ru
    ADD MAPPING FOR numword WITH simple;

ALTER TEXT SEARCH CONFIGURATION public.ru
    ADD MAPPING FOR email WITH simple;

ALTER TEXT SEARCH CONFIGURATION public.ru
    ADD MAPPING FOR url WITH simple;

ALTER TEXT SEARCH CONFIGURATION public.ru
    ADD MAPPING FOR host WITH simple;

ALTER TEXT SEARCH CONFIGURATION public.ru
    ADD MAPPING FOR sfloat WITH simple;

ALTER TEXT SEARCH CONFIGURATION public.ru
    ADD MAPPING FOR version WITH simple;

ALTER TEXT SEARCH CONFIGURATION public.ru
    ADD MAPPING FOR hword_numpart WITH simple;

ALTER TEXT SEARCH CONFIGURATION public.ru
    ADD MAPPING FOR hword_part WITH public.russian_ispell, russian_stem;

ALTER TEXT SEARCH CONFIGURATION public.ru
    ADD MAPPING FOR hword_asciipart WITH english_stem;

ALTER TEXT SEARCH CONFIGURATION public.ru
    ADD MAPPING FOR numhword WITH simple;

ALTER TEXT SEARCH CONFIGURATION public.ru
    ADD MAPPING FOR asciihword WITH english_stem;

ALTER TEXT SEARCH CONFIGURATION public.ru
    ADD MAPPING FOR hword WITH public.russian_ispell, russian_stem;

ALTER TEXT SEARCH CONFIGURATION public.ru
    ADD MAPPING FOR url_path WITH simple;

ALTER TEXT SEARCH CONFIGURATION public.ru
    ADD MAPPING FOR file WITH simple;

ALTER TEXT SEARCH CONFIGURATION public.ru
    ADD MAPPING FOR "float" WITH simple;

ALTER TEXT SEARCH CONFIGURATION public.ru
    ADD MAPPING FOR "int" WITH simple;

ALTER TEXT SEARCH CONFIGURATION public.ru
    ADD MAPPING FOR uint WITH simple;


ALTER TEXT SEARCH CONFIGURATION public.ru OWNER TO admin_test;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: addresses; Type: TABLE; Schema: public; Owner: admin_test
--

CREATE TABLE public.addresses (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    address text NOT NULL,
    user_id uuid,
    is_active boolean DEFAULT false
);


ALTER TABLE public.addresses OWNER TO admin_test;

--
-- Name: orders; Type: TABLE; Schema: public; Owner: admin_test
--

CREATE TABLE public.orders (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    status text NOT NULL,
    address_id text NOT NULL,
    order_products text NOT NULL,
    apartment_or_office text,
    intercom text,
    entrance text,
    floor text,
    courier_comment text,
    leave_at_door boolean DEFAULT false,
    final_price numeric(10,2) NOT NULL,
    order_items uuid[],
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.orders OWNER TO admin_test;

--
-- Name: products; Type: TABLE; Schema: public; Owner: admin_test
--

CREATE TABLE public.products (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    restaurant_id uuid NOT NULL,
    name text NOT NULL,
    price numeric(10,2) NOT NULL,
    image_url text DEFAULT 'default_product.jpg'::text,
    weight integer NOT NULL,
    category text NOT NULL,
    tsvector_column tsvector
);


ALTER TABLE public.products OWNER TO admin_test;

--
-- Name: promocodes; Type: TABLE; Schema: public; Owner: admin_test
--

CREATE TABLE public.promocodes (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    promocode text NOT NULL,
    discount numeric(10,2) NOT NULL,
    user_id uuid NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    expires_at timestamp with time zone DEFAULT (now() + '1 day'::interval) NOT NULL,
    is_used boolean DEFAULT false
);


ALTER TABLE public.promocodes OWNER TO admin_test;

--
-- Name: restaurant_tags; Type: TABLE; Schema: public; Owner: admin_test
--

CREATE TABLE public.restaurant_tags (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name text NOT NULL
);


ALTER TABLE public.restaurant_tags OWNER TO admin_test;

--
-- Name: restaurant_tags_relations; Type: TABLE; Schema: public; Owner: admin_test
--

CREATE TABLE public.restaurant_tags_relations (
    restaurant_id uuid NOT NULL,
    tag_id uuid NOT NULL
);


ALTER TABLE public.restaurant_tags_relations OWNER TO admin_test;

--
-- Name: restaurants; Type: TABLE; Schema: public; Owner: admin_test
--

CREATE TABLE public.restaurants (
    id uuid NOT NULL,
    name text NOT NULL,
    banner_url text DEFAULT 'default_restaurant.jpg'::text,
    address text DEFAULT ''::text,
    rating double precision,
    rating_count double precision,
    description text DEFAULT ''::text,
    working_mode_from integer DEFAULT 8,
    working_mode_to integer DEFAULT 23,
    delivery_time_from integer DEFAULT 50,
    delivery_time_to integer DEFAULT 60,
    tsvector_column tsvector,
    CONSTRAINT restaurants_rating_check CHECK (((rating >= (0)::double precision) AND (rating <= (5)::double precision))),
    CONSTRAINT restaurants_rating_count_check CHECK ((rating_count >= (0)::double precision))
);


ALTER TABLE public.restaurants OWNER TO admin_test;

--
-- Name: reviews; Type: TABLE; Schema: public; Owner: admin_test
--

CREATE TABLE public.reviews (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    restaurant_id uuid NOT NULL,
    review_text text,
    rating integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    CONSTRAINT reviews_rating_check CHECK (((rating >= 1) AND (rating <= 5))),
    CONSTRAINT reviews_review_text_check CHECK ((char_length(review_text) <= 300))
);


ALTER TABLE public.reviews OWNER TO admin_test;

--
-- Name: users; Type: TABLE; Schema: public; Owner: admin_test
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    login text NOT NULL,
    phone_number text,
    first_name text NOT NULL,
    last_name text NOT NULL,
    description text DEFAULT ''::text,
    user_pic text DEFAULT 'default_user.jpg'::text,
    password_hash bytea NOT NULL,
    secret2fa bytea
);


ALTER TABLE public.users OWNER TO admin_test;

--
-- Name: addresses addresses_pkey; Type: CONSTRAINT; Schema: public; Owner: admin_test
--

ALTER TABLE ONLY public.addresses
    ADD CONSTRAINT addresses_pkey PRIMARY KEY (id);


--
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: public; Owner: admin_test
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: admin_test
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: promocodes promocodes_pkey; Type: CONSTRAINT; Schema: public; Owner: admin_test
--

ALTER TABLE ONLY public.promocodes
    ADD CONSTRAINT promocodes_pkey PRIMARY KEY (id);


--
-- Name: restaurant_tags restaurant_tags_name_key; Type: CONSTRAINT; Schema: public; Owner: admin_test
--

ALTER TABLE ONLY public.restaurant_tags
    ADD CONSTRAINT restaurant_tags_name_key UNIQUE (name);


--
-- Name: restaurant_tags restaurant_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: admin_test
--

ALTER TABLE ONLY public.restaurant_tags
    ADD CONSTRAINT restaurant_tags_pkey PRIMARY KEY (id);


--
-- Name: restaurant_tags_relations restaurant_tags_relations_pkey; Type: CONSTRAINT; Schema: public; Owner: admin_test
--

ALTER TABLE ONLY public.restaurant_tags_relations
    ADD CONSTRAINT restaurant_tags_relations_pkey PRIMARY KEY (restaurant_id, tag_id);


--
-- Name: restaurants restaurants_pkey; Type: CONSTRAINT; Schema: public; Owner: admin_test
--

ALTER TABLE ONLY public.restaurants
    ADD CONSTRAINT restaurants_pkey PRIMARY KEY (id);


--
-- Name: reviews reviews_pkey; Type: CONSTRAINT; Schema: public; Owner: admin_test
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT reviews_pkey PRIMARY KEY (id);


--
-- Name: users users_login_key; Type: CONSTRAINT; Schema: public; Owner: admin_test
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_login_key UNIQUE (login);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: admin_test
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_products_tsv; Type: INDEX; Schema: public; Owner: admin_test
--

CREATE INDEX idx_products_tsv ON public.products USING gin (tsvector_column);


--
-- Name: idx_restaurants_tsv; Type: INDEX; Schema: public; Owner: admin_test
--

CREATE INDEX idx_restaurants_tsv ON public.restaurants USING gin (tsvector_column);


--
-- Name: idx_users_user_login; Type: INDEX; Schema: public; Owner: admin_test
--

CREATE UNIQUE INDEX idx_users_user_login ON public.users USING btree (login);


--
-- Name: unique_active_address_per_user; Type: INDEX; Schema: public; Owner: admin_test
--

CREATE UNIQUE INDEX unique_active_address_per_user ON public.addresses USING btree (user_id) WHERE (is_active = true);


--
-- Name: reviews after_review_insert; Type: TRIGGER; Schema: public; Owner: admin_test
--

CREATE TRIGGER after_review_insert AFTER INSERT ON public.reviews FOR EACH ROW EXECUTE FUNCTION public.update_restaurant_rating();


--
-- Name: orders trg_issue_promocode_on_paid; Type: TRIGGER; Schema: public; Owner: admin_test
--

CREATE TRIGGER trg_issue_promocode_on_paid AFTER UPDATE ON public.orders FOR EACH ROW EXECUTE FUNCTION public.issue_promocode_on_paid();


--
-- Name: products trg_update_product_tsv; Type: TRIGGER; Schema: public; Owner: admin_test
--

CREATE TRIGGER trg_update_product_tsv BEFORE INSERT OR UPDATE ON public.products FOR EACH ROW EXECUTE FUNCTION public.update_product_tsvector();


--
-- Name: restaurants trg_update_restaurant_tsv; Type: TRIGGER; Schema: public; Owner: admin_test
--

CREATE TRIGGER trg_update_restaurant_tsv BEFORE INSERT OR UPDATE ON public.restaurants FOR EACH ROW EXECUTE FUNCTION public.update_restaurant_tsvector();


--
-- Name: addresses addresses_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin_test
--

ALTER TABLE ONLY public.addresses
    ADD CONSTRAINT addresses_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE SET NULL;


--
-- Name: orders orders_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin_test
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: products products_restaurant_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin_test
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_restaurant_id_fkey FOREIGN KEY (restaurant_id) REFERENCES public.restaurants(id) ON DELETE CASCADE;


--
-- Name: promocodes promocodes_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin_test
--

ALTER TABLE ONLY public.promocodes
    ADD CONSTRAINT promocodes_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: restaurant_tags_relations restaurant_tags_relations_restaurant_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin_test
--

ALTER TABLE ONLY public.restaurant_tags_relations
    ADD CONSTRAINT restaurant_tags_relations_restaurant_id_fkey FOREIGN KEY (restaurant_id) REFERENCES public.restaurants(id) ON DELETE CASCADE;


--
-- Name: restaurant_tags_relations restaurant_tags_relations_tag_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin_test
--

ALTER TABLE ONLY public.restaurant_tags_relations
    ADD CONSTRAINT restaurant_tags_relations_tag_id_fkey FOREIGN KEY (tag_id) REFERENCES public.restaurant_tags(id) ON DELETE CASCADE;


--
-- Name: reviews reviews_restaurant_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin_test
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT reviews_restaurant_id_fkey FOREIGN KEY (restaurant_id) REFERENCES public.restaurants(id) ON DELETE CASCADE;


--
-- Name: reviews reviews_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin_test
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT reviews_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

