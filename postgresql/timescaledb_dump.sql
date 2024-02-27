--
-- PostgreSQL database dump
--

-- Dumped from database version 16.2 (Debian 16.2-1.pgdg120+2)
-- Dumped by pg_dump version 16.2 (Ubuntu 16.2-1.pgdg22.04+1)

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
-- Name: timescaledb; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS timescaledb WITH SCHEMA public;


--
-- Name: EXTENSION timescaledb; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION timescaledb IS 'Enables scalable inserts and complex queries for time-series data (Community Edition)';


--
-- Name: timescaledb_toolkit; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS timescaledb_toolkit WITH SCHEMA public;


--
-- Name: EXTENSION timescaledb_toolkit; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION timescaledb_toolkit IS 'Library of analytical hyperfunctions, time-series pipelining, and other SQL utilities';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: acq; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.acq (
    rx_time timestamp with time zone NOT NULL,
    experiment_id integer,
    signal_id integer,
    doppler double precision,
    code_phase double precision,
    acf_corr real[],
    noise_floor real,
    acq_mode smallint
);


ALTER TABLE public.acq OWNER TO postgres;

--
-- Name: _hyper_3_2_chunk; Type: TABLE; Schema: _timescaledb_internal; Owner: postgres
--

CREATE TABLE _timescaledb_internal._hyper_3_2_chunk (
    CONSTRAINT constraint_2 CHECK (((rx_time >= '2014-08-14 00:00:00+00'::timestamp with time zone) AND (rx_time < '2014-08-21 00:00:00+00'::timestamp with time zone)))
)
INHERITS (public.acq);


ALTER TABLE _timescaledb_internal._hyper_3_2_chunk OWNER TO postgres;

--
-- Name: application; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.application (
    id integer NOT NULL,
    application_name smallint
);


ALTER TABLE public.application OWNER TO postgres;

--
-- Name: application_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.application_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.application_id_seq OWNER TO postgres;

--
-- Name: application_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.application_id_seq OWNED BY public.application.id;


--
-- Name: experiment; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.experiment (
    id integer NOT NULL,
    mission_operation_id integer,
    experiment_type_id integer,
    start_time timestamp without time zone,
    end_time timestamp without time zone
);


ALTER TABLE public.experiment OWNER TO postgres;

--
-- Name: experiment_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.experiment_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.experiment_id_seq OWNER TO postgres;

--
-- Name: experiment_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.experiment_id_seq OWNED BY public.experiment.id;


--
-- Name: experiment_type; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.experiment_type (
    id integer NOT NULL,
    experiment_type_name character varying(32)
);


ALTER TABLE public.experiment_type OWNER TO postgres;

--
-- Name: experiment_type_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.experiment_type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.experiment_type_id_seq OWNER TO postgres;

--
-- Name: experiment_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.experiment_type_id_seq OWNED BY public.experiment_type.id;


--
-- Name: log; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.log (
    rx_time timestamp with time zone NOT NULL,
    experiment_type_id integer,
    priority integer,
    size integer,
    message text
);


ALTER TABLE public.log OWNER TO postgres;

--
-- Name: log_priority; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.log_priority (
    id integer NOT NULL,
    log_priority_name smallint
);


ALTER TABLE public.log_priority OWNER TO postgres;

--
-- Name: log_priority_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.log_priority_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.log_priority_id_seq OWNER TO postgres;

--
-- Name: log_priority_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.log_priority_id_seq OWNED BY public.log_priority.id;


--
-- Name: mission_operation; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mission_operation (
    id integer NOT NULL,
    mission_operation_name integer,
    start_time timestamp without time zone,
    end_time timestamp without time zone
);


ALTER TABLE public.mission_operation OWNER TO postgres;

--
-- Name: mission_operation_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mission_operation_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.mission_operation_id_seq OWNER TO postgres;

--
-- Name: mission_operation_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mission_operation_id_seq OWNED BY public.mission_operation.id;


--
-- Name: nav; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.nav (
    rx_time timestamp with time zone NOT NULL,
    experiment_id integer,
    application_id integer,
    wn integer,
    tow bigint,
    decimals double precision,
    n_sat integer,
    pos_x double precision,
    pos_y double precision,
    pos_z double precision,
    vel_x double precision,
    vel_y double precision,
    vel_z double precision,
    pos_std real,
    vel_std real,
    tim_std double precision,
    clock_bias double precision,
    clock_drift double precision,
    ggto double precision,
    gdop double precision,
    pdop double precision,
    hdop double precision,
    vdop double precision,
    tdop double precision
);


ALTER TABLE public.nav OWNER TO postgres;

--
-- Name: operation_modes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.operation_modes (
    id integer NOT NULL,
    rx_operation_mode_name smallint
);


ALTER TABLE public.operation_modes OWNER TO postgres;

--
-- Name: operation_modes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.operation_modes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.operation_modes_id_seq OWNER TO postgres;

--
-- Name: operation_modes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.operation_modes_id_seq OWNED BY public.operation_modes.id;


--
-- Name: pod; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.pod (
    rx_time timestamp with time zone NOT NULL,
    experiment_id integer,
    application_id integer,
    wn integer,
    tow bigint,
    decimals double precision,
    n_sat integer,
    pos_x double precision,
    pos_y double precision,
    pos_z double precision,
    vel_x double precision,
    vel_y double precision,
    vel_z double precision,
    pos_std real[],
    vel_std real[],
    clock_bias double precision,
    clock_drift double precision,
    ambig_vec real[],
    ambig_acc real[]
);


ALTER TABLE public.pod OWNER TO postgres;

--
-- Name: raw; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.raw (
    rx_time timestamp with time zone NOT NULL,
    experiment_id integer,
    signal_id integer,
    sv_id integer,
    fd_raw double precision,
    fd_rate_raw double precision,
    carrier_phase double precision,
    pr_raw double precision
);


ALTER TABLE public.raw OWNER TO postgres;

--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO postgres;

--
-- Name: signal; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.signal (
    id integer NOT NULL,
    signal_name character varying(8)
);


ALTER TABLE public.signal OWNER TO postgres;

--
-- Name: signal_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.signal_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.signal_id_seq OWNER TO postgres;

--
-- Name: signal_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.signal_id_seq OWNED BY public.signal.id;


--
-- Name: sta; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sta (
    rx_time timestamp with time zone NOT NULL,
    experiment_id integer,
    rx_operation_mode_id integer,
    application_id integer,
    signal_id integer,
    number_of_channels smallint,
    sv_id integer,
    channel_status smallint,
    number_of_apps smallint,
    ecc_error_count bigint,
    cpu_temp real,
    frontend_temp real,
    qn400_version_number bigint
);


ALTER TABLE public.sta OWNER TO postgres;

--
-- Name: application id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application ALTER COLUMN id SET DEFAULT nextval('public.application_id_seq'::regclass);


--
-- Name: experiment id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.experiment ALTER COLUMN id SET DEFAULT nextval('public.experiment_id_seq'::regclass);


--
-- Name: experiment_type id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.experiment_type ALTER COLUMN id SET DEFAULT nextval('public.experiment_type_id_seq'::regclass);


--
-- Name: log_priority id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.log_priority ALTER COLUMN id SET DEFAULT nextval('public.log_priority_id_seq'::regclass);


--
-- Name: mission_operation id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mission_operation ALTER COLUMN id SET DEFAULT nextval('public.mission_operation_id_seq'::regclass);


--
-- Name: operation_modes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.operation_modes ALTER COLUMN id SET DEFAULT nextval('public.operation_modes_id_seq'::regclass);


--
-- Name: signal id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.signal ALTER COLUMN id SET DEFAULT nextval('public.signal_id_seq'::regclass);


--
-- Data for Name: hypertable; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--

INSERT INTO _timescaledb_catalog.hypertable VALUES (1, 'public', 'log', '_timescaledb_internal', '_hyper_1', 1, '_timescaledb_functions', 'calculate_chunk_interval', 0, 0, NULL, 0);
INSERT INTO _timescaledb_catalog.hypertable VALUES (2, 'public', 'raw', '_timescaledb_internal', '_hyper_2', 1, '_timescaledb_functions', 'calculate_chunk_interval', 0, 0, NULL, 0);
INSERT INTO _timescaledb_catalog.hypertable VALUES (3, 'public', 'acq', '_timescaledb_internal', '_hyper_3', 1, '_timescaledb_functions', 'calculate_chunk_interval', 0, 0, NULL, 0);
INSERT INTO _timescaledb_catalog.hypertable VALUES (4, 'public', 'pod', '_timescaledb_internal', '_hyper_4', 1, '_timescaledb_functions', 'calculate_chunk_interval', 0, 0, NULL, 0);
INSERT INTO _timescaledb_catalog.hypertable VALUES (5, 'public', 'sta', '_timescaledb_internal', '_hyper_5', 1, '_timescaledb_functions', 'calculate_chunk_interval', 0, 0, NULL, 0);
INSERT INTO _timescaledb_catalog.hypertable VALUES (6, 'public', 'nav', '_timescaledb_internal', '_hyper_6', 1, '_timescaledb_functions', 'calculate_chunk_interval', 0, 0, NULL, 0);


--
-- Data for Name: chunk; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--

INSERT INTO _timescaledb_catalog.chunk VALUES (2, 3, '_timescaledb_internal', '_hyper_3_2_chunk', NULL, false, 0, false, '2024-02-27 04:10:33.662793+00');


--
-- Data for Name: dimension; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--

INSERT INTO _timescaledb_catalog.dimension VALUES (1, 1, 'rx_time', 'timestamp with time zone', true, NULL, NULL, NULL, 604800000000, NULL, NULL, NULL);
INSERT INTO _timescaledb_catalog.dimension VALUES (2, 2, 'rx_time', 'timestamp with time zone', true, NULL, NULL, NULL, 604800000000, NULL, NULL, NULL);
INSERT INTO _timescaledb_catalog.dimension VALUES (3, 3, 'rx_time', 'timestamp with time zone', true, NULL, NULL, NULL, 604800000000, NULL, NULL, NULL);
INSERT INTO _timescaledb_catalog.dimension VALUES (4, 4, 'rx_time', 'timestamp with time zone', true, NULL, NULL, NULL, 604800000000, NULL, NULL, NULL);
INSERT INTO _timescaledb_catalog.dimension VALUES (5, 5, 'rx_time', 'timestamp with time zone', true, NULL, NULL, NULL, 604800000000, NULL, NULL, NULL);
INSERT INTO _timescaledb_catalog.dimension VALUES (6, 6, 'rx_time', 'timestamp with time zone', true, NULL, NULL, NULL, 604800000000, NULL, NULL, NULL);


--
-- Data for Name: dimension_slice; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--

INSERT INTO _timescaledb_catalog.dimension_slice VALUES (2, 3, 1407974400000000, 1408579200000000);


--
-- Data for Name: chunk_constraint; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--

INSERT INTO _timescaledb_catalog.chunk_constraint VALUES (2, 2, 'constraint_2', NULL);
INSERT INTO _timescaledb_catalog.chunk_constraint VALUES (2, NULL, '2_3_acq_experiment_id_fkey', 'acq_experiment_id_fkey');
INSERT INTO _timescaledb_catalog.chunk_constraint VALUES (2, NULL, '2_4_acq_signal_id_fkey', 'acq_signal_id_fkey');


--
-- Data for Name: chunk_index; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--

INSERT INTO _timescaledb_catalog.chunk_index VALUES (2, '_hyper_3_2_chunk_acq_rx_time_idx', 3, 'acq_rx_time_idx');


--
-- Data for Name: compression_chunk_size; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--



--
-- Data for Name: compression_settings; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--



--
-- Data for Name: continuous_agg; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--



--
-- Data for Name: continuous_agg_migrate_plan; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--



--
-- Data for Name: continuous_agg_migrate_plan_step; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--



--
-- Data for Name: continuous_aggs_bucket_function; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--



--
-- Data for Name: continuous_aggs_hypertable_invalidation_log; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--



--
-- Data for Name: continuous_aggs_invalidation_threshold; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--



--
-- Data for Name: continuous_aggs_materialization_invalidation_log; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--



--
-- Data for Name: continuous_aggs_watermark; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--



--
-- Data for Name: metadata; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--

INSERT INTO _timescaledb_catalog.metadata VALUES ('install_timestamp', '2024-02-27 03:23:33.426197+00', true);
INSERT INTO _timescaledb_catalog.metadata VALUES ('timescaledb_version', '2.14.2', false);
INSERT INTO _timescaledb_catalog.metadata VALUES ('exported_uuid', '3be1b85c-2e29-4c7e-a609-0b473724bda0', true);


--
-- Data for Name: tablespace; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--



--
-- Data for Name: bgw_job; Type: TABLE DATA; Schema: _timescaledb_config; Owner: postgres
--



--
-- Data for Name: _hyper_3_2_chunk; Type: TABLE DATA; Schema: _timescaledb_internal; Owner: postgres
--

INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:00:21+00', 1, 0, -5209.43359375, 876.7109985351562, '{129734,69890,87178,10203,60086,403952,236730,23708,441152,120181,703993,9.189891e+06,6.166499e+06,1.5258913e+07,7.164357e+06,3.148576e+06,2.614252e+06,853484,412351,1.837646e+06,530250,1.092197e+06,578547,675235,718086,320423,83859,95526,39351,23998,264348,691907,230050,336801,1.110215e+06,2.559973e+06,1.955043e+06,973143,571514,768852,1.088788e+06,1.197492e+06,279976,248012,765879,544953,230023,436095,772462,190451,751678,1.980154e+06,556689,582543,301850,28500,101773,29160,326941,614511,71320,95644,75097,10289,3244,46190,116720,11496,29753,254796,4.955987e+06,843584,482564,691841,1.899535e+06,720205,50479,306889,47596,52373,463245,208784,442701,223601,235448,9763,181245,68003,298969,326284,142753,379147,284628,273609,115238,86936,278691,320702,1.594504e+06,129857,19721,105754,57649,4060,2870,86143,665164,720224,57195,88952,178409,8935,171930,201479,805684,952097,847191,340891,139954,39067,311803,483912,730909,1.0968085e+07,6.823087e+07,1.5151976e+08,3.6647187e+08,7.419797e+08,6.0084755e+08,3.14255e+08,9.5391136e+07,1.8197824e+07,3.035355e+06,489166,618830,1.214427e+06,2.140037e+06,1.976024e+06,422557,308359,515,33481,13921,124937,665926,336058,110991,1349,1500,7185,33505,1.430654e+06,1.582628e+06,443959,40034,51420,487864,265825,168150,221848,348116,1.710901e+06,901001,720354,1.113227e+06,1.598144e+06,824388,271312,54765,58390,111120,331551,139203,104841,63413,4902,187604,584446,870266,250929,5644,1.459895e+06,1.761611e+06,597618,1.050482e+06,15769,1.757272e+06,421692,281637,29228,399598,172743,252053,480411,453625,128713,1319,688682,222347,202396,248462,138450,740939,392929,76333,481523,3.479512e+06,1.013294e+06,1.625695e+06,7.712927e+06,1.2048963e+07,6.742175e+06,1.304881e+06,242031,146622,36367,90432,4935,47176,196290,111064,309940,164949,154738,100202,19528,502286,139451,2.406048e+06,1.218027e+06,212551,268853,47064,310511,1.026787e+06,220522,855993,527151,346427,2.059194e+06,362244,1.869225e+06,1.354679e+06,850672,274121,49903,14956,135739,363522,2.883171e+06,3.361917e+06,529358,1.382112e+06,44473,979459}', 441.1763, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:00:34+00', 1, 2, -5472.94677734375, 2685.886474609375, '{1030,66,182092,74440,3408,16098,39982,57399,2593,16143,22351,8670,1608,149,18259,241281,436997,24775,17421,324836,12042,199234,2445,5745,43032,5,2404,11638,10198,10351,21507,8161,5942,54358,32733,13733,356761,61319,22129,424355,205021,2889,74948,118074,35732,25923,847,2623,29585,73542,2024,13898,352378,1112,130,39775,25368,12442,14913,2494,122,1198,9939,25787,5360,6718,43707,57548,1275,69196,10076,302107,229985,63631,51871,122122,27785,30556,47443,130403,5798,30841,8364,22328,14729,20296,8409,6627,11601,23135,49507,94646,4266,15399,84555,162884,80825,2495,98985,471324,323022,60500,145361,277218,34148,243016,22701,13933,36348,382844,9101,680,7248,1315,2052,29173,93225,190154,940,11323,458904,58196,222244,1.127884e+06,2.291255e+06,13369,1.6567558e+07,9.292409e+07,1.70691e+06,536911,9.843382e+06,618419,46776,129081,8020,4458,160,28026,1393,2682,14309,5151,1280,32107,10219,4642,9731,1030,913,13018,25348,37176,8561,2912,399,100389,2943,155,7538,155448,10178,16786,23389,16726,118808,8105,154176,173612,92156,14591,141980,246716,44364,6251,335561,12994,3568,3826,12964,85535,11644,2227,207686,206311,1736,15563,98567,3285,23426,32685,136,20987,14365,26451,2861,124,64147,50709,45444,386781,29329,21393,116105,143851,11344,43962,239499,24480,203884,99221,50480,74528,125283,7668,174925,7704,23201,29599,38780,6923,2364,10566,44909,70526,131960,13694,39737,336863,61095,632108,25166,9473,342851,24348,16151,64197,47309,369,5874,22289,32827,32422,8323,14159,76289,203,27,3463,8548,141228,523,518,18477,6703,12025}', 126.36282, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:00:48+00', 1, 2, -9286.4462890625, 1421.45849609375, '{19,14,8,5,6,4,10,9,6,10,11,10,10,12,7,9,15,26,18,9,14,15,21,15,24,13,6,11,8,15,10,12,10,6,9,13,12,8,12,14,13,11,9,10,6,7,13,11,12,8,9,9,8,6,9,9,10,10,6,15,15,8,8,7,14,13,16,12,13,16,29,15,13,12,18,10,15,13,10,9,11,14,12,12,11,13,18,13,9,9,12,8,16,7,7,20,18,8,10,12,15,11,10,19,16,20,12,29,16,20,21,11,6,13,12,13,12,8,14,14,13,9,9,11,19,18,21,71,60,25,22,17,24,29,26,12,13,10,15,10,8,11,7,7,14,10,11,13,9,13,20,16,14,7,9,11,11,12,13,8,12,6,13,13,10,7,22,28,10,15,24,19,13,8,12,13,5,15,11,7,5,10,12,11,12,13,16,18,9,13,13,7,9,6,13,10,5,10,15,13,9,13,13,12,15,14,12,15,15,7,18,8,5,12,6,10,14,10,9,10,9,9,10,8,12,12,11,13,10,10,8,8,7,6,6,8,10,19,11,7,14,16,10,17,7,14,19,15,8,4,6,13,6,7,6}', 124.76464, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:00:48+00', 1, 0, -23428.884765625, 120.8844985961914, '{5,7,10,14,11,8,5,7,5,3,2,5,4,4,5,8,10,10,5,4,3,2,2,3,4,7,9,6,3,2,1,4,7,5,6,2,2,4,4,3,3,5,5,5,5,5,4,6,7,7,3,5,5,7,6,5,5,11,12,12,9,6,6,1,5,8,11,13,16,9,2,3,2,6,6,9,4,7,6,1,1,2,7,4,8,5,5,4,3,5,2,5,4,5,6,4,11,14,13,12,7,8,7,3,3,5,6,8,6,6,4,2,4,1,3,3,5,8,11,10,3,2,7,18,30,72,116,132,115,69,39,11,3,3,4,4,4,2,1,2,3,4,4,4,3,4,2,2,2,3,3,2,1,1,2,3,1,2,2,2,3,4,2,4,6,9,9,7,6,5,2,4,3,1,1,1,2,3,2,2,3,4,8,8,6,4,2,2,2,2,4,3,3,11,11,5,2,4,3,2,3,5,5,5,2,4,3,4,5,6,3,5,3,3,3,2,9,7,10,16,12,6,5,4,5,3,2,4,3,4,3,6,10,9,7,7,9,14,10,4,4,3,3,3,4,12,8,8,10,6,4,7,8,9,6}', 443.89432, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:01:09+00', 1, 2, -15092.4755859375, 2164.156494140625, '{38,46,33,47,30,27,40,27,31,46,40,32,28,38,31,34,34,34,41,50,44,24,36,40,30,33,48,36,47,36,33,48,41,42,46,29,41,37,36,21,23,39,31,30,29,31,29,35,28,45,29,24,28,40,38,45,54,48,34,32,25,37,40,34,32,26,36,43,27,26,29,36,31,29,32,35,31,41,35,35,38,33,35,36,43,42,42,51,59,60,46,42,38,42,38,33,49,45,41,36,30,33,32,29,38,42,45,46,47,35,39,38,29,35,39,49,40,32,45,39,40,29,37,56,52,51,78,111,110,59,63,77,53,39,51,34,34,43,43,39,50,54,39,37,37,31,32,35,34,34,47,42,39,32,28,29,26,31,30,29,24,37,32,31,37,37,33,41,36,40,43,25,36,48,39,37,44,38,39,37,48,35,29,40,45,40,33,35,46,48,40,32,34,34,35,32,39,44,33,33,37,40,35,38,27,24,25,29,57,48,34,47,52,55,46,47,52,54,48,44,38,31,37,49,50,52,46,51,59,46,34,40,33,28,43,43,38,39,32,49,52,36,30,39,28,25,30,33,42,39,42,53,34,36,41}', 126.19298, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:01:14+00', 1, 0, -5535.748046875, 257.28448486328125, '{4,6,5,5,7,8,5,8,5,6,5,5,5,4,5,7,5,10,8,9,7,3,3,2,2,1,1,2,5,10,11,13,6,3,1,1,2,4,2,1,1,4,5,6,6,10,7,5,4,4,4,2,3,3,1,1,1,3,8,8,5,6,4,4,6,4,6,5,3,4,6,4,3,4,1,1,5,3,8,6,2,2,3,2,5,4,3,5,5,8,6,9,11,9,1,4,4,2,4,5,5,4,3,8,10,5,6,2,4,4,4,7,4,2,2,3,2,7,6,7,7,2,2,7,16,30,61,98,93,78,47,25,12,4,4,5,17,10,8,4,3,3,2,4,5,3,5,5,9,9,11,7,8,8,5,3,3,2,5,6,15,9,12,8,3,2,4,2,5,3,2,8,11,2,3,5,7,6,10,6,8,4,3,3,2,1,3,4,2,5,2,1,1,5,4,4,6,7,4,3,2,2,2,2,3,4,4,6,6,5,5,5,5,3,3,6,5,2,2,3,4,4,4,2,3,4,4,4,4,2,1,1,1,6,8,8,6,3,5,4,6,5,6,8,9,4,3,2,2,4,2,6,7,5,4}', 433.93234, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:01:16+00', 1, 2, -453.3031005859375, 106.22149658203125, '{14,14,3,7,11,3,5,7,11,10,8,7,13,13,21,14,16,18,16,14,16,8,7,17,15,14,6,17,17,11,13,11,11,8,7,8,13,13,11,13,15,13,7,14,10,6,13,10,8,11,9,11,9,10,18,11,8,19,15,8,13,11,6,16,13,7,7,14,7,9,12,8,6,10,10,19,10,6,12,8,10,11,17,23,12,9,14,10,8,9,11,9,16,11,8,28,21,13,6,13,8,8,10,8,8,8,14,8,15,15,13,13,9,16,15,19,12,10,9,9,8,6,12,29,28,6,74,95,23,10,24,11,10,11,11,8,15,17,10,9,7,8,8,11,9,7,13,7,9,10,11,17,14,21,15,12,17,13,12,10,10,12,13,12,10,16,8,8,13,12,12,13,10,8,13,5,10,11,18,8,10,12,13,8,7,5,6,9,13,12,9,15,12,13,8,7,11,18,8,10,19,10,18,12,6,9,11,20,13,18,18,15,9,11,6,8,11,11,15,14,6,10,8,8,8,9,11,7,10,9,7,7,8,11,9,17,11,10,17,18,13,14,13,9,9,12,10,8,4,6,8,6,12,16,17}', 124.231926, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:01:17+00', 1, 0, -2413.953125, 897.85302734375, '{8,8,10,8,7,5,7,9,8,6,4,9,5,4,7,7,2,1,1,1,1,3,2,4,8,10,15,10,9,10,4,4,7,4,3,6,7,9,3,1,5,2,1,3,2,2,5,3,5,12,12,9,9,3,4,3,3,4,3,2,4,4,3,2,2,1,1,2,4,6,6,4,3,4,3,5,5,3,3,1,1,1,2,3,2,2,3,6,3,5,5,5,5,5,3,3,4,5,6,5,2,3,1,2,2,2,4,4,1,2,1,3,3,3,7,10,8,7,5,2,3,5,6,10,13,29,42,57,29,20,14,6,3,4,3,4,6,6,6,7,15,5,5,2,4,8,7,6,8,4,3,1,1,2,4,2,1,2,2,3,5,4,4,1,4,3,5,9,9,6,5,4,4,5,4,5,8,10,10,10,11,10,3,4,3,8,8,6,4,6,4,4,4,3,5,9,12,8,6,8,14,7,11,9,8,4,6,5,3,5,4,5,4,5,3,2,2,2,2,3,5,3,4,4,3,3,2,1,3,1,4,5,6,5,6,6,5,5,3,3,4,4,1,1,3,3,2,4,2,2,2,2,4,4,6}', 448.11966, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:02:19+00', 1, 0, -7723.31298828125, 25.063499450683594, '{7,9,10,10,8,8,12,15,12,11,10,7,9,6,9,6,6,12,11,14,14,12,8,12,15,12,12,12,13,13,21,22,19,14,7,6,5,4,5,7,8,7,6,4,3,3,5,7,8,9,7,8,10,14,13,13,10,12,8,9,8,13,14,9,12,7,4,2,4,6,8,14,11,8,11,18,14,11,10,13,12,9,7,9,13,12,10,10,9,6,8,7,8,7,7,10,10,7,8,7,6,8,10,9,9,10,16,21,7,8,6,9,14,10,5,8,9,9,11,20,18,18,9,8,11,37,67,103,80,59,35,13,8,7,13,14,8,6,10,6,11,8,9,7,8,12,11,5,8,8,11,8,6,4,15,15,9,7,8,7,5,5,6,4,6,8,14,13,15,12,8,9,14,11,10,8,5,6,6,6,5,7,7,5,8,5,4,6,8,10,4,3,7,8,10,12,7,5,5,7,10,8,5,6,9,6,9,9,13,15,13,14,9,7,7,6,11,10,10,9,9,12,16,11,9,9,9,8,7,5,5,7,7,9,9,7,10,10,12,9,11,11,11,5,7,8,5,6,11,8,13,15,14,11,6}', 431.54868, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:02:20+00', 1, 0, -11719.7001953125, 239.3820037841797, '{33,36,30,24,20,19,21,23,21,19,26,25,31,33,43,38,31,28,26,25,20,15,19,20,19,14,21,20,15,15,27,26,33,38,31,24,25,27,27,26,27,21,19,14,21,21,30,28,34,27,24,23,22,22,18,19,25,31,39,27,24,20,23,22,19,18,19,19,19,18,14,15,14,19,24,26,17,24,25,23,21,31,26,24,23,22,18,17,17,14,17,17,17,22,16,20,25,29,32,30,38,37,35,31,40,29,36,37,33,32,25,26,30,26,24,24,29,33,27,26,29,27,31,32,38,52,89,96,52,35,20,27,33,36,26,28,18,21,17,22,18,21,36,35,30,28,26,19,20,25,22,23,24,23,27,33,27,28,25,13,13,18,30,17,26,29,25,19,19,20,22,13,14,18,15,20,22,20,24,21,26,28,24,19,26,29,20,16,16,17,13,15,27,39,27,20,13,17,16,12,14,13,18,26,31,35,34,37,33,26,27,16,22,31,40,37,43,44,36,31,29,30,23,19,18,15,19,20,21,25,18,25,27,23,31,34,37,35,36,33,24,17,18,20,32,41,27,27,21,15,16,13,22,23,20}', 424.9613, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:02:40+00', 1, 2, -15008.037109375, 3047.6875, '{61,33,34,53,52,35,42,63,41,42,49,48,62,57,51,56,45,43,31,50,55,38,36,51,52,39,45,50,38,30,39,45,47,43,43,55,46,36,63,45,54,59,49,48,47,36,41,33,44,67,38,43,37,39,41,40,31,51,49,32,41,56,42,40,50,38,43,37,41,55,50,42,35,35,35,39,53,70,54,53,55,53,48,35,46,45,49,42,63,41,42,48,40,50,50,58,39,49,48,30,51,49,36,43,52,57,41,34,38,33,36,43,46,42,52,52,48,61,39,53,44,44,39,65,64,59,79,117,100,63,70,74,46,50,61,39,32,38,44,48,51,45,47,34,40,40,34,53,52,40,50,58,40,47,51,46,29,42,45,41,37,42,47,44,42,52,48,46,50,46,36,42,57,48,37,43,57,34,33,45,42,37,36,26,45,41,52,47,44,41,56,52,53,68,47,42,45,44,40,37,45,35,32,37,49,50,43,39,41,50,44,40,38,32,39,50,49,33,36,54,42,38,36,52,50,55,56,42,45,51,48,52,49,45,45,65,50,50,56,64,53,38,59,45,30,41,32,45,42,40,62,51,53,58,46}', 121.31983, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:05:48+00', 1, 2, 147.2875213623047, 170.5, '{53,47,34,58,45,50,59,54,60,38,43,45,46,44,48,41,28,58,56,41,50,51,43,46,44,48,40,40,52,49,42,45,47,41,44,44,45,43,30,29,47,46,44,67,58,43,50,46,51,57,52,68,51,45,45,54,41,41,42,37,34,36,44,51,41,33,53,44,38,50,39,48,51,57,50,40,40,46,51,30,42,40,36,45,38,61,54,44,51,39,44,41,43,45,46,46,55,36,53,42,33,41,43,40,46,38,49,56,56,52,56,64,49,45,66,45,47,52,52,51,45,53,52,56,66,43,67,124,74,38,58,60,46,39,47,37,42,51,49,33,42,47,45,38,46,70,43,38,41,39,39,40,46,49,29,33,44,41,49,30,38,47,43,56,57,28,48,45,39,44,51,51,49,44,47,53,46,33,39,47,44,30,39,41,46,52,48,48,36,50,45,43,45,41,46,39,43,41,47,41,42,41,30,48,47,41,52,49,50,60,55,68,68,43,60,68,47,55,50,38,67,44,33,37,44,43,46,39,39,40,44,50,68,55,37,41,41,55,40,41,45,29,35,51,54,35,49,41,34,40,46,52,46,56,38}', 119.068306, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:05:49+00', 1, 0, -19126.185546875, 932.9760131835938, '{22,30,35,34,32,26,23,19,21,33,31,36,33,28,26,30,26,33,29,36,38,39,34,33,26,23,24,26,23,19,19,24,27,27,28,14,15,23,33,33,29,28,25,24,29,24,22,31,36,34,37,38,27,23,21,18,23,23,23,19,24,23,21,24,24,27,26,20,22,31,32,35,37,38,29,27,36,33,30,28,21,34,35,42,34,37,24,29,25,26,19,25,22,18,19,29,31,23,23,32,28,30,28,32,30,22,23,23,23,34,32,35,32,28,24,24,33,37,32,31,36,27,25,30,37,51,84,117,111,91,60,51,34,29,33,39,39,30,31,28,35,36,26,23,22,25,23,24,23,21,32,31,33,33,31,25,25,28,37,39,32,33,28,20,18,20,20,32,34,38,42,38,33,38,35,26,38,40,34,37,33,19,24,29,30,27,31,24,31,36,32,36,30,27,25,27,26,24,35,29,36,32,37,37,36,26,29,32,33,29,20,20,19,25,32,29,30,34,33,38,39,36,39,38,33,30,24,30,33,28,18,23,31,36,31,23,20,23,23,22,21,22,31,23,27,23,23,22,27,20,19,19,18,19,30}', 426.8409, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:05:50+00', 1, 0, -2203.399169921875, 683.0230102539062, '{37,46,42,40,45,41,40,36,38,45,48,42,24,23,20,22,24,33,43,45,33,43,32,41,39,39,35,25,27,32,30,31,34,28,28,41,42,41,40,49,34,22,18,23,29,33,33,40,47,49,59,41,49,49,46,38,52,36,27,28,30,27,32,35,32,32,28,36,35,38,35,38,45,25,38,32,27,31,33,37,43,42,43,43,28,33,31,32,42,34,29,23,32,30,30,38,35,42,39,37,36,41,34,34,34,39,45,38,48,44,47,39,34,24,28,24,32,30,39,39,40,24,28,34,54,88,114,122,99,73,53,39,39,40,39,30,33,25,30,31,34,39,47,54,36,37,26,37,26,37,34,39,33,39,29,28,24,23,30,28,27,34,37,25,22,32,30,45,34,42,33,41,41,44,34,32,30,37,34,43,57,55,46,36,32,40,38,34,30,45,47,47,49,51,54,46,46,36,43,40,43,44,42,59,42,37,29,32,34,29,27,29,43,35,22,28,30,29,27,25,27,31,35,31,29,43,38,31,27,30,39,34,30,19,19,23,26,23,19,36,44,43,57,55,53,53,39,30,29,28,32,38,40,46,50}', 423.53802, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:07:26+00', 1, 2, 294.7298889160156, 156.6894989013672, '{70,51,58,49,59,67,57,59,56,40,50,52,44,46,60,69,48,45,57,62,61,63,52,59,56,65,82,75,63,53,46,52,55,50,43,63,45,51,61,63,51,63,72,56,53,70,52,53,54,39,44,64,57,47,54,65,66,84,54,47,53,53,70,45,45,51,58,58,55,53,49,56,67,65,58,55,50,62,57,54,54,50,49,51,56,44,56,71,54,65,70,66,40,57,70,45,69,79,55,80,60,50,67,71,61,59,62,57,52,50,74,43,44,61,55,61,70,59,62,64,65,66,51,61,77,56,73,132,113,54,71,70,66,61,58,56,60,62,39,46,50,64,64,67,68,73,63,62,50,41,52,55,53,70,63,64,66,69,48,54,55,56,50,61,64,68,64,52,44,43,66,69,65,51,68,78,40,50,76,54,56,50,50,62,51,46,74,50,38,54,60,62,67,70,59,54,67,54,66,56,62,85,61,60,65,58,65,52,54,57,49,45,46,44,59,51,50,64,73,53,58,52,55,60,66,60,65,66,64,56,63,69,58,66,53,53,72,52,46,65,85,52,54,74,53,51,57,57,63,58,43,41,49,61,51}', 118.51627, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:07:32+00', 1, 2, -7897.0390625, 3020.74853515625, '{39,56,45,35,41,32,54,46,25,36,52,35,40,49,43,32,41,45,49,42,40,40,40,31,32,44,44,43,52,45,47,43,38,41,55,48,41,46,50,54,50,53,55,53,48,38,50,55,43,39,42,35,39,45,61,34,36,39,47,41,43,40,37,43,57,44,50,49,42,43,35,27,50,45,51,58,42,41,41,55,38,46,43,64,66,48,54,48,53,46,45,48,52,45,39,45,42,46,45,52,47,41,46,47,40,37,40,30,37,49,52,49,46,51,53,43,48,45,41,48,50,45,50,48,78,45,63,122,87,58,54,65,43,33,44,31,36,42,52,48,37,47,54,51,44,38,51,40,26,39,42,31,39,45,45,53,39,44,51,43,38,43,44,35,52,57,40,32,35,32,34,49,39,43,50,60,59,30,43,43,43,36,39,44,33,29,37,45,35,41,55,44,45,50,39,39,59,65,44,48,54,48,42,47,50,50,39,42,46,39,45,37,38,45,47,39,46,56,54,56,44,40,48,46,42,45,45,47,48,35,40,51,51,31,35,37,38,37,32,38,36,32,33,37,43,52,45,46,62,43,38,49,60,43,39}', 121.679375, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:10:45+00', 1, 0, -21143.873046875, 917.9719848632812, '{15,10,14,10,10,10,16,13,14,15,16,25,28,21,15,12,8,10,13,17,16,11,13,16,14,17,15,16,13,16,14,13,15,18,19,22,14,14,14,15,12,10,8,7,14,24,17,16,16,9,14,10,11,21,18,24,21,21,17,11,12,12,17,14,14,17,12,17,12,10,18,15,11,12,12,15,10,12,18,13,13,23,21,16,15,18,15,9,13,10,22,16,13,10,9,8,17,16,26,30,18,15,19,20,25,23,16,24,20,24,24,28,22,25,27,22,14,13,19,17,24,23,24,25,37,56,79,85,76,48,34,26,25,18,8,14,13,16,12,7,15,10,11,12,13,16,11,20,15,17,14,12,13,12,18,18,14,17,10,9,14,9,18,20,9,11,14,13,14,11,17,23,12,16,21,21,16,15,15,15,18,16,13,20,25,18,15,18,23,16,10,15,10,13,16,15,15,10,10,11,10,13,10,8,5,13,10,8,8,9,11,12,15,14,9,10,15,17,10,6,8,9,9,12,13,9,8,9,16,21,17,14,17,16,13,12,13,13,12,14,16,17,14,17,19,20,17,14,15,16,12,10,14,14,12}', 412.4421, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:10:50+00', 1, 2, -20488.109375, 1218.052001953125, '{62,70,56,50,46,50,49,46,51,53,47,44,38,53,53,38,52,64,53,61,53,45,51,44,53,53,35,38,56,41,56,53,56,58,65,49,46,54,55,63,52,52,47,62,61,41,59,52,52,58,50,54,64,61,61,53,64,51,52,50,45,54,52,50,46,53,43,46,51,55,56,57,60,53,62,50,37,51,45,61,45,43,58,59,57,43,37,44,44,56,35,50,55,40,43,46,53,54,49,38,47,38,50,62,43,40,44,61,37,49,50,51,43,46,43,40,51,32,45,52,48,65,58,58,93,75,74,131,85,93,81,57,52,45,50,51,49,44,75,46,54,68,49,43,44,39,52,44,46,62,58,56,44,50,44,52,40,49,56,56,59,49,50,67,59,45,34,37,42,35,47,47,43,48,52,53,47,57,61,43,44,49,53,52,48,54,62,67,44,51,57,55,48,53,67,40,49,48,46,49,48,46,54,43,42,60,48,45,46,43,46,60,56,49,51,50,61,51,55,50,52,49,38,56,48,50,65,62,42,56,60,54,55,50,47,47,55,42,37,62,67,46,49,51,49,50,49,41,33,41,46,51,56,45,48}', 118.80548, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:12:17+00', 1, 2, 742.5266723632812, 59.67499923706055, '{60,58,64,40,42,50,55,46,56,49,64,57,47,55,56,56,43,62,63,49,58,66,56,53,59,53,48,62,66,55,63,62,52,59,52,48,58,64,45,52,63,47,54,53,51,49,56,41,43,52,47,57,43,70,63,59,49,63,57,81,48,41,66,55,66,67,55,50,65,58,49,62,66,68,66,61,64,55,47,51,71,61,49,63,61,64,47,48,49,70,66,52,55,65,58,38,59,71,67,79,63,50,59,58,73,56,55,71,69,52,54,63,60,60,46,53,66,67,58,79,80,51,51,74,54,50,115,132,76,64,86,69,51,64,69,61,46,57,53,44,51,44,48,49,46,51,51,46,54,54,45,48,55,48,54,67,47,70,62,46,51,60,65,55,61,51,49,43,48,55,53,49,54,62,47,54,60,59,57,68,61,53,47,49,61,43,53,67,61,74,61,63,59,59,76,60,52,55,63,70,62,56,56,46,62,58,58,56,46,58,67,48,53,60,52,48,46,67,59,52,51,62,65,51,45,62,71,53,65,82,64,70,51,45,53,49,53,47,54,56,54,67,50,58,62,69,61,71,57,75,49,52,70,57,59}', 118.125145, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:13:53+00', 1, 0, -1568.9417724609375, 252.68099975585938, '{29,34,28,31,26,24,15,27,33,44,36,26,21,19,20,15,21,26,26,25,24,35,26,17,23,26,28,28,25,16,21,24,39,41,40,37,27,17,14,24,25,25,24,22,20,26,33,30,36,38,39,26,25,26,29,29,26,28,22,28,24,23,21,24,24,22,28,15,21,30,31,27,31,26,25,19,21,22,26,33,33,28,37,49,40,35,43,43,32,24,26,32,28,24,31,28,28,23,16,21,28,22,21,31,23,22,21,19,13,16,19,20,22,33,18,32,34,25,22,20,22,23,30,30,31,42,72,106,92,63,69,45,40,21,19,20,20,23,22,19,23,20,21,23,25,26,42,40,43,37,38,26,25,24,19,23,15,20,23,27,35,39,28,26,25,21,30,23,26,38,35,21,18,20,23,25,22,20,22,20,16,37,43,38,27,23,17,19,26,28,28,25,17,22,25,29,39,34,35,25,24,27,20,26,25,23,20,31,27,33,26,28,27,27,15,20,26,24,22,23,18,9,20,16,32,39,29,32,35,24,26,25,23,20,29,32,32,36,43,33,20,21,19,24,31,29,34,37,31,29,26,39,37,32,17}', 423.9312, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:20:53+00', 1, 0, -27228.40234375, 1014.3045043945312, '{22,25,16,17,18,23,24,17,17,12,20,26,24,29,28,20,19,23,19,17,14,17,20,13,18,15,16,18,24,23,20,24,23,15,17,17,22,24,24,22,25,17,20,22,22,22,16,17,27,22,23,22,25,22,28,20,15,13,10,12,16,16,13,18,30,35,24,21,18,22,26,25,20,24,26,31,31,26,23,15,16,16,18,21,21,16,20,19,21,19,16,21,30,33,25,25,27,15,16,17,19,17,9,12,8,12,11,14,12,17,17,13,14,22,21,25,27,31,29,34,28,19,21,32,46,55,81,92,71,57,33,22,15,16,17,17,18,18,14,18,18,20,26,25,18,15,13,9,13,8,9,13,17,19,20,22,12,13,20,12,18,21,31,29,21,24,21,17,23,26,24,21,18,28,40,25,26,26}', 401.6497, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:20:54+00', 1, 0, -28947.755859375, 547.9869995117188, '{39,37,25,36,32,39,38,40,28,31,39,35,39,47,41,40,36,39,34,38,48,52,46,53,40,33,36,31,28,39,27,32,39,39,37,38,58,56,46,40,44,32,31,34,45,40,38,41,36,38,31,37,37,39,27,37,39,37,39,41,42,41,51,54,54,56,38,32,29,31,39,36,32,41,43,36,39,38,30,31,38,35,44,60,51,38,43,38,38,38,28,30,31,32,33,38,37,37,25,29,27,33,26,27,28,30,37,32,32,25,31,30,35,52,37,30,41,28,44,41,44,45,47,63,77,103,121,129,95,85,73,61,57,45,30,40,47,45,47,40,39,40,26,26,24,26,27,33,33,36,35,30,31,34,32,44,44,44,35,45,32,27,37,34,36,42,37,39,38,41,36,39,39,33,29,36,33,30,27,21,26,31,38,39,35,38,39,38,42,34,34,27,26,29,32,36,40,49,43,34,39,45,29,30,49,55,51,28,30,25,37,49,49,53,51,31,28,35,35,44,36,52,48,43,52,46,34,30,29,33,31,32,41,32,38,40,36,43,47,38,49,34,29,41,50,43,36,27,30,39,36,37,38,33,45}', 407.27612, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:24:19+00', 1, 0, -7260.212890625, 597.0910034179688, '{2,1,1,2,5,4,6,7,9,13,12,5,4,1,1,3,7,5,8,11,9,11,7,6,3,3,3,3,3,3,3,2,4,2,6,6,6,9,10,13,6,8,9,8,4,6,6,5,9,15,12,9,9,5,3,4,8,5,3,1,2,3,1,1,3,8,11,4,4,3,3,7,11,10,13,9,4,4,4,1,4,3,4,5,5,4,5,6,5,7,7,6,4,1,2,3,3,2,3,8,11,6,4,2,2,1,3,1,3,2,4,5,4,3,3,3,2,3,3,7,5,2,162,628,1331,2285,3505,3791,2490,1535,773,283,33,8,5,5,6,4,3,6,12,9,8,5,8,4,2,2,3,1,2,1,2,1,5,6,7,6,7,7,5,8,5,5,5,9,6,1,3,2,2,1,1,1,3,1,3,2,8,12,14,16,16,12,4,3,1,1,1,6,8,6,10,9,5,2,3,2,3,1,1,4,6,5,4,6,5,5,5,4,12,21,7,8,9,9,4,3,4,4,4,7,5,3,2,2,3,4,3,5,7,14,16,12,11,7,6,3,2,3,11,18,21,19,35,20,10,3,3,7,7,6,7,10,5}', 436.17133, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:40:03+00', 1, 2, -17974.248046875, 2257.931396484375, '{16,17,11,14,29,11,9,15,23,16,8,15,25,12,7,17,17,15,11,14,13,14,18,17,16,14,14,15,20,18,21,20,36,36,22,35,39,18,20,18,9,15,25,21,17,18,14,10,15,28,24,18,29,28,26,14,18,25,16,12,9,20,20,14,15,17,17,22,12,12,14,12,14,17,22,15,23,15,22,28,19,17,15,21,18,25,24,24,26,23,11,19,15,14,12,19,12,12,14,8,20,14,18,18,12,18,16,20,16,13,14,16,17,13,22,24,26,20,14,17,20,24,24,29,39,41,70,81,49,24,28,39,15,16,15,17,13,7,17,14,8,18,18,20,20,16,24,16,11,21,17,11,19,14,16,15,15,17,17,17,17,9,13,13,9,16,24,10,9,15,18,13,29,19,19,16,13,14,22,23,20,13,18,20,20,16,15,18,18,11,18,13,9,30,23,12,18,28,23,27,25,21,28,26,22,19,20,18,13,12,20,20,13,19,18,24,18,18,13,11,12,23,19,6,15,16,17,19,14,19,16,14,11,14,9,10,20,22,15,15,24,25,25,15,20,21,13,23,27,18,22,23,15,17,18}', 117.78425, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:41:01+00', 1, 2, -7108.67822265625, 1250.447021484375, '{5,4,3,7,3,0,6,4,4,7,17,8,2,5,7,3,3,6,2,2,2,4,6,6,7,5,5,3,2,3,7,7,3,8,11,7,7,3,4,5,8,9,8,6,6,5,4,9,12,9,10,9,7,2,6,6,5,6,11,5,2,6,4,3,4,4,4,4,2,2,3,5,11,7,4,6,4,3,4,4,3,5,5,4,1,11,10,10,11,5,4,8,4,9,4,3,5,4,3,5,3,4,7,5,4,7,8,4,6,5,3,9,9,5,4,7,4,2,6,10,7,2,47,413,375,95,1271,1823,351,145,500,178,38,11,3,7,15,5,2,4,5,6,5,7,7,8,4,6,4,3,3,5,6,6,5,3,2,4,6,3,3,4,4,1,3,2,4,2,3,3,11,4,2,6,7,5,3,2,3,3,2,3,6,9,12,2,6,4,2,5,9,5,5,4,3,2,2,7,7,6,6,13,8,9,6,10,7,5,4,9,7,4,3,4,7,4,6,5,4,4,5,5,6,5,5,7,6,3,7,7,4,5,4,3,9,5,5,5,11,3,2,4,15,3,5,5,3,5,5,6,2,3,2,5,6}', 123.207146, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:54:24+00', 1, 0, -19339.64453125, 426.0794982910156, '{7,5,9,6,5,9,13,13,14,20,16,9,6,3,5,5,10,9,12,6,7,8,9,10,10,10,12,11,9,7,7,6,8,6,7,5,7,9,15,14,14,15,9,15,16,16,11,11,5,11,10,9,13,11,8,5,7,9,9,8,9,10,9,11,7,7,7,9,7,9,6,8,10,7,11,13,19,17,12,6,3,7,8,8,11,11,12,10,9,11,8,11,11,8,11,11,5,8,8,7,3,3,4,12,9,11,6,9,8,9,9,11,11,9,8,5,8,7,10,14,6,10,9,22,47,52,69,71,60,46,30,27,18,16,14,12,11,7,11,14,12,12,14,12,6,8,7,7,6,3,4,7,6,9,16,19,17,14,10,9,10,6,5,7,6,5,7,6,5,11,15,17,16,23,29,24,26,23,14,13,7,4,12,19,21,23,16,14,19,15,11,9,8,9,7,8,8,12,12,11,13,25,30,23,29,26,13,12,15,11,16,14,7,6,6,9,9,9,7,8,6,6,8,9,7,5,5,5,8,18,19,15,10,7,10,7,13,18,19,18,5,6,5,9,13,10,5,5,6,6,13,8,10,11,17}', 437.3469, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:58:43+00', 1, 0, -1688.514404296875, 804.5894775390625, '{19,22,26,30,28,28,34,41,37,35,29,23,22,18,25,28,28,29,34,27,22,15,17,13,8,11,15,11,18,23,31,31,27,24,25,31,37,35,36,37,25,26,23,25,23,20,19,17,18,15,12,11,18,16,14,15,13,12,12,11,11,12,16,16,20,18,15,10,13,9,8,17,19,20,27,27,32,29,25,26,22,22,27,24,24,18,18,12,13,16,16,10,17,31,35,35,27,25,20,22,16,12,16,16,12,11,12,10,14,19,14,14,21,22,21,23,20,23,21,29,32,40,44,52,50,73,91,95,64,66,59,42,29,17,23,24,25,26,26,26,19,10,13,15,19,22,25,23,37,41,50,47,52,49,40,24,20,15,10,10,10,15,18,17,28,26,26,26,24,24,29,38,37,42,33,28,25,18,16,17,14,18,20,22,18,19,22,29,30,26,22,24,37,25,25,23,22,15,14,16,13,20,20,20,21,21,15,13,15,15,12,18,14,19,21,22,23,22,24,16,13,18,23,30,29,26,19,13,11,11,16,21,23,33,29,18,18,20,18,14,13,20,15,20,23,22,31,21,22,19,15,17,14,15,12}', 445.5653, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 20:59:16+00', 1, 0, -20336.146484375, 615.6754760742188, '{29,31,27,27,33,28,31,37,39,41,46,37,37,35,35,34,35,38,44,34,33,29,32,29,40,36,36,39,48,39,37,38,30,27,30,24,24,36,32,27,25,22,26,34,30,28,27,33,28,27,22,28,29,31,32,34,39,40,37,34,31,30,30,23,28,33,30,29,29,29,34,24,28,31,32,29,39,56,43,44,31,24,31,21,20,26,29,32,31,33,35,41,42,32,34,34,30,31,27,35,27,30,31,32,25,25,22,30,29,28,33,32,33,25,28,28,31,40,42,43,40,39,40,51,63,70,95,125,115,79,54,33,27,30,34,23,24,18,31,26,29,24,23,21,16,20,28,25,37,42,41,35,31,40,36,24,27,28,25,28,28,22,31,29,21,21,28,29,38,41,40,43,35,33,23,29,32,25,29,28,32,30,22,17,23,23,29,25,30,30,28,32,24,22,32,24,29,24,22,25,32,30,31,34,29,38,41,33,41,33,36,30,33,26,26,28,31,42,38,44,32,34,31,24,24,28,26,33,32,42,36,34,35,35,35,26,24,28,27,31,29,27,21,26,24,23,30,21,36,33,34,32,31,34,34}', 451.66452, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 21:00:00+00', 1, 0, -15521.08203125, 379.7034912109375, '{6,5,4,3,2,1,1,1,4,4,3,3,3,4,5,9,8,3,4,8,11,19,20,22,24,27,29,29,23,11,4,3,2,6,7,2,1,2,5,3,3,2,4,4,2,3,1,1,0,2,2,5,5,2,2,1,2,2,3,3,4,3,1,1,2,6,13,23,29,38,32,29,23,23,26,20,18,12,10,5,2,2,3,4,7,13,20,9,6,3,2,2,2,2,2,5,3,3,4,4,7,5,3,2,3,3,1,3,6,5,4,1,1,2,2,2,5,3,3,6,8,19,163,622,1402,2495,3932,5085,4494,3016,1870,981,352,78,19,15,6,4,4,6,8,8,4,4,5,3,6,5,6,7,9,5,3,2,3,3,1,1,1,1,2,1,1,2,2,2,7,18,31,25,16,11,10,5,3,3,6,5,14,13,15,11,10,8,10,14,13,11,5,4,6,7,2,3,3,1,1,1,2,8,5,3,3,4,2,1,2,1,1,1,2,1,3,3,5,7,6,5,7,17,33,37,20,5,2,1,10,17,16,15,16,20,18,11,7,2,2,3,3,3,2,3,5,4,3,2,2,2,1,1,2,2,3,4,2}', 504.98605, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 21:00:06+00', 1, 0, -22982.478515625, 322.0744934082031, '{18,15,7,9,7,10,10,11,17,7,9,13,14,14,12,12,9,10,8,7,9,9,15,10,12,10,11,11,13,19,18,21,19,9,8,7,10,12,18,21,19,12,7,6,7,12,18,22,23,14,16,14,9,7,3,8,11,11,10,9,9,10,8,7,9,8,7,5,3,3,4,13,15,11,13,7,11,11,8,12,12,11,10,11,6,7,6,11,14,11,22,24,25,15,11,11,9,9,10,9,9,17,17,13,7,3,3,2,3,5,7,5,5,5,7,12,15,14,6,7,12,16,12,16,23,33,57,67,50,35,29,27,27,24,19,13,8,7,6,9,8,10,8,10,9,11,5,9,10,8,8,16,22,9,9,10,9,10,10,12,13,13,17,24,25,22,13,10,7,5,8,9,8,8,7,6,4,9,9,18,16,13,9,12,9,9,12,13,8,9,19,14,17,20,13,14,12,12,13,15,5,3,5,1,3,2,2,3,7,9,11,8,10,12,18,26,24,17,12,15,10,12,23,21,20,17,19,13,11,12,7,5,5,6,10,8,15,20,23,17,8,15,12,17,18,14,7,10,13,26,28,20,17,17,17}', 447.61404, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 21:00:07+00', 1, 0, -18779.484375, 323.4385070800781, '{17,15,14,12,8,10,13,14,16,13,14,12,17,10,8,15,17,14,8,12,12,12,12,14,13,17,14,11,7,18,22,24,27,26,21,20,14,12,11,12,9,11,15,12,16,14,12,15,9,11,13,13,11,17,17,13,10,10,10,10,16,11,12,13,9,13,20,20,25,31,20,18,13,20,25,22,23,23,19,19,18,18,17,16,14,14,20,15,19,24,18,11,12,11,14,14,19,16,16,19,27,24,19,17,11,15,20,16,12,15,14,14,10,15,19,17,15,9,11,10,14,18,12,18,41,54,69,89,81,64,36,28,18,15,12,11,14,20,22,20,16,17,14,15,16,18,16,15,16,12,10,10,12,18,16,16,13,14,15,16,16,14,15,12,10,10,9,11,12,16,16,16,7,10,7,10,10,10,9,6,9,12,18,22,15,16,24,16,17,15,9,9,17,16,17,15,17,15,15,11,14,11,13,17,18,19,29,20,24,19,20,27,25,22,22,17,12,13,13,16,17,21,16,22,17,15,12,11,16,13,13,19,19,18,17,14,13,9,8,13,10,12,12,12,15,16,15,6,10,18,15,15,17,17,14}', 450.41037, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 21:02:06+00', 1, 0, -22731.7578125, 18.7549991607666, '{14,26,15,18,11,6,6,10,10,14,11,7,9,7,4,4,11,17,19,15,13,10,6,4,5,11,14,12,17,19,24,21,12,9,6,8,9,10,12,11,13,13,8,7,8,11,10,8,11,11,11,11,18,15,17,8,8,6,13,18,15,11,9,9,8,8,8,4,4,9,23,29,19,16,14,15,13,9,6,6,5,6,7,7,7,8,13,17,17,16,10,8,9,15,8,13,7,4,7,9,10,16,10,12,14,15,16,22,40,54,73,66,73,64,39,29,21,17,13,11,10,13,10,4,8,17,15,14,9,13,6,9,13,12,11,15,15,10,8,5,5,10,7,7,7,9,13,11,11,12,12,10,6,8,6,8,11,12,12,15,7,11,18,17,14,13,10,11,8,8,19,28,19,11,8,14,20,19,17,13,14,16,13,9,9,10,9,14,16,15,15,5,6,6,7,7,9,15,27,16,24,22,13,10,8,7,13,16,21,17,16,19,9,13,14,21,19,19,14,11,5,4,5,10,3,5,5,4,2,3,3,6,9,10,13,15,8,8,8,11,13,14,14,11,7,4,8,7,5,7,10,8,6,7,8}', 451.65, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 21:02:40+00', 1, 2, -21604.7890625, 2101.071533203125, '{46,45,50,36,41,53,39,38,37,50,43,32,42,40,34,45,52,40,37,36,36,55,32,30,50,35,51,52,30,49,51,50,45,44,51,39,53,46,40,48,48,47,45,36,32,36,39,50,41,27,44,41,26,42,46,49,51,46,48,34,30,40,49,48,41,40,43,60,42,31,39,50,41,44,32,36,48,46,23,50,45,50,39,41,51,52,33,50,52,48,41,41,46,38,43,36,44,45,38,55,38,37,36,50,37,42,35,35,24,40,52,43,29,44,44,33,39,49,36,44,47,46,46,54,53,85,102,117,101,75,49,47,59,48,51,37,56,54,52,37,40,45,39,40,35,25,44,46,35,43,41,50,42,56,43,59,44,39,45,44,46,37,46,40,28,38,34,37,25,43,39,26,40,48,38,49,42,41,44,33,49,44,40,47,37,35,59,72,38,29,50,55,27,43,37,33,29,29,44,28,30,42,37,34,45,44,36,31,42,50,46,36,35,24,30,34,27,24,37,57,47,29,37,37,36,43,26,37,44,46,42,31,31,40,28,33,44,38,35,30,36,36,31,34,39,36,45,45,39,39,36,34,25,34,49}', 119.66867, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 21:04:02+00', 1, 0, -22494.453125, 50.97949981689453, '{40,45,33,33,34,38,36,34,23,26,24,30,38,28,33,30,28,30,34,35,33,27,34,27,38,39,28,36,45,44,46,39,33,27,19,25,31,38,36,29,29,27,38,37,35,29,29,29,23,32,27,24,27,26,28,33,34,38,31,32,26,30,35,29,32,21,23,33,31,26,27,27,23,31,22,28,33,31,30,27,26,27,27,25,30,18,19,26,30,39,42,38,39,24,22,26,28,23,36,30,30,24,25,21,21,21,28,22,30,32,35,28,32,28,30,35,30,27,21,18,21,35,30,45,67,78,108,142,133,101,79,52,34,29,32,31,29,33,28,22,16,25,27,27,27,32,32,30,30,35,32,27,24,32,37,33,25,27,26,21,23,25,25,32,19,29,41,41,43,32,27,35,35,35,35,28,21,21,21,23,34,42,33,35,33,28,35,40,40,51,47,38,29,24,27,31,34,43,32,41,27,30,26,30,31,36,32,28,30,38,38,31,30,25,26,21,27,31,28,39,40,46,46,39,40,30,30,29,30,24,37,39,34,37,39,37,40,42,44,40,37,32,41,31,28,27,21,21,23,30,27,35,38,33,36}', 414.2926, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 21:04:37+00', 1, 2, -21461.74609375, 3724.06103515625, '{49,40,55,77,56,54,68,53,41,56,54,51,41,46,43,42,51,61,50,40,44,50,39,47,39,47,55,32,45,44,34,42,48,35,35,36,43,41,37,53,54,75,51,46,47,51,53,27,53,39,36,45,36,42,43,51,35,35,53,44,33,42,37,44,37,39,39,52,53,45,48,62,64,57,53,46,58,67,58,63,59,48,44,38,46,46,33,33,34,44,48,48,46,56,53,34,36,51,46,38,43,42,35,40,42,41,44,45,54,51,40,46,59,33,31,57,35,50,34,42,47,49,50,71,67,91,111,133,118,70,83,62,45,51,36,45,52,66,46,35,36,47,51,53,51,53,44,38,37,40,53,50,42,49,44,43,56,58,65,54,57,55,46,56,49,46,49,40,36,37,61,52,51,51,47,50,62,57,40,64,49,48,41,34,46,45,36,43,39,36,55,54,69,48,49,53,44,46,37,46,35,49,45,46,37,39,35,38,49,57,50,38,39,50,44,45,39,41,50,47,49,56,49,49,48,42,51,61,39,49,65,41,76,47,34,46,35,31,44,39,52,59,42,53,54,57,55,50,40,50,51,46,34,44,44}', 116.29815, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 21:06:27+00', 1, 2, -21324.521484375, 1175.0860595703125, '{28,19,32,48,41,39,42,56,38,26,48,37,27,34,38,34,29,35,26,32,37,29,38,33,30,46,37,28,35,32,37,36,35,36,34,44,35,34,41,37,27,28,28,34,35,31,28,40,43,28,38,26,31,27,36,44,31,43,49,38,26,28,41,42,32,41,40,33,28,32,31,41,35,36,36,30,38,28,41,46,32,29,29,30,23,35,27,32,34,31,33,37,41,48,44,29,50,50,27,35,36,28,28,32,40,39,33,51,42,27,35,37,27,32,54,45,23,30,44,46,40,29,46,48,48,62,93,106,92,61,38,50,53,43,44,32,42,34,45,43,36,30,35,31,27,31,26,24,39,32,25,50,38,41,37,29,36,30,33,42,41,28,26,28,37,33,34,40,33,36,38,29,22,32,38,30,28,36,27,47,37,38,39,28,33,31,35,40,34,21,28,31,36,38,34,28,31,30,40,38,37,30,33,45,33,47,38,31,36,37,30,31,45,42,41,28,33,38,26,27,27,21,29,43,25,24,39,25,42,36,32,31,34,35,37,33,60,54,43,33,36,42,28,43,34,33,34,37,31,33,32,39,35,36,42}', 116.82035, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 21:08:14+00', 1, 2, -21191.560546875, 2647.864990234375, '{60,39,39,47,49,36,35,55,42,36,45,51,54,44,49,55,41,45,37,45,53,29,45,61,48,43,43,39,41,40,49,42,36,44,56,48,47,41,44,39,39,34,35,36,35,34,29,40,52,34,57,46,41,52,57,54,42,53,37,58,53,57,61,56,45,40,44,36,29,35,40,45,49,40,47,39,44,51,42,50,57,49,30,45,47,48,37,38,39,43,53,44,53,54,39,36,54,49,46,54,41,61,40,53,54,37,46,41,45,48,47,48,32,55,51,33,49,49,36,46,52,49,48,61,100,116,106,119,83,67,68,61,46,46,46,55,42,44,44,58,45,44,41,41,45,43,49,29,30,43,31,35,36,38,53,39,39,43,52,50,39,52,59,59,48,55,45,36,34,40,61,56,48,40,49,39,40,37,49,32,28,51,67,39,43,35,48,50,39,40,41,46,35,47,50,46,36,68,48,41,41,52,42,46,41,51,42,42,45,46,47,45,59,44,48,50,36,36,41,44,45,39,44,53,32,38,57,32,39,47,61,50,49,54,50,43,47,46,47,57,41,53,52,40,47,67,47,50,48,48,41,36,43,37,52}', 114.89345, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 21:09:44+00', 1, 2, 1587.42138671875, 3748.27197265625, '{41,53,55,44,51,53,35,48,70,68,55,55,54,45,42,37,38,31,28,49,53,42,37,68,53,51,58,50,51,53,46,52,57,40,45,58,60,56,52,68,53,74,61,55,45,49,53,41,54,50,49,49,64,38,51,62,37,44,44,50,50,49,46,68,43,49,46,36,54,59,58,55,48,57,60,45,42,57,65,50,36,38,53,56,47,49,55,46,52,44,43,47,46,58,57,52,42,45,57,44,43,56,63,58,38,40,55,45,45,41,46,48,58,48,43,45,53,45,58,49,68,47,47,57,52,54,63,125,93,53,60,75,64,54,52,42,53,39,63,54,56,56,50,35,48,51,43,43,51,48,42,50,46,37,54,48,50,45,31,39,46,38,37,39,46,36,44,38,56,50,47,62,66,54,52,49,47,58,51,63,60,58,58,70,52,58,52,44,53,58,66,55,42,48,48,54,42,53,68,52,56,56,52,60,51,55,43,50,50,54,46,48,45,51,50,49,45,44,57,67,41,58,71,59,68,61,50,68,41,50,44,42,40,46,49,45,55,41,43,38,57,51,50,59,61,45,52,55,51,43,40,43,50,52,42}', 114.54812, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 21:13:23+00', 1, 2, 1930.5162353515625, 3499.342041015625, '{45,56,75,47,46,64,60,44,46,60,52,50,61,54,53,69,77,61,60,65,47,50,56,60,61,66,62,57,72,47,54,66,58,69,59,58,51,59,67,54,49,50,71,58,58,61,61,50,65,70,57,63,59,60,60,35,60,51,43,55,60,52,47,70,56,49,58,61,39,56,52,50,72,70,48,75,56,57,56,42,57,63,49,44,65,67,64,67,49,68,52,44,66,63,51,60,55,56,55,63,64,71,62,47,56,55,54,64,61,61,71,40,43,50,41,62,57,46,61,45,44,51,49,67,57,50,121,133,60,61,67,58,46,59,51,61,61,73,69,65,48,45,48,59,46,38,54,60,50,61,68,68,45,60,53,63,62,74,48,61,53,51,73,62,66,67,71,62,48,66,52,60,54,44,54,63,49,61,59,66,81,65,67,71,62,66,58,49,52,61,64,58,62,58,57,55,63,58,60,55,53,70,64,66,51,60,42,61,80,46,76,73,64,72,60,59,68,64,56,76,83,52,73,85,49,53,60,50,44,68,84,58,53,47,61,51,55,63,56,59,53,65,61,67,64,75,67,56,45,65,66,54,48,39,47}', 113.93559, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 21:15:20+00', 1, 2, -17179.8515625, 3938.89111328125, '{56,52,40,49,49,43,41,53,54,42,38,53,70,65,44,39,36,51,47,40,54,43,48,44,56,46,55,75,42,33,51,58,46,48,52,38,41,51,49,44,39,56,43,46,52,42,52,45,44,38,52,58,49,43,51,43,42,37,44,60,38,44,41,28,45,44,38,26,42,48,42,55,48,47,55,53,54,45,42,63,43,41,63,53,52,49,45,39,55,43,45,52,43,36,47,64,41,37,50,52,46,46,53,34,26,26,51,37,64,60,50,62,54,48,51,51,38,42,35,37,49,42,56,70,59,52,78,124,104,63,57,72,61,51,54,43,48,45,34,50,47,50,37,47,62,44,32,46,50,39,38,37,45,56,55,50,37,54,38,38,44,54,48,45,52,50,40,43,41,60,58,49,45,35,37,54,57,42,47,56,49,50,44,42,63,43,53,45,42,48,47,36,56,42,34,47,49,37,46,57,40,50,41,33,32,37,41,48,50,47,39,53,69,50,45,32,39,42,50,53,53,29,32,36,49,54,43,46,59,44,43,50,38,47,38,48,41,46,39,31,47,44,44,39,37,51,47,47,52,33,41,39,32,43,49}', 114.541, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 21:17:05+00', 1, 2, -17071.4453125, 1018.2260131835938, '{39,53,49,41,47,47,42,37,42,38,38,38,36,45,35,42,43,44,45,39,36,37,42,47,29,41,55,48,46,66,46,40,56,49,36,38,52,59,42,30,42,40,37,43,50,53,44,37,55,62,55,51,52,48,59,45,47,52,53,42,39,50,43,41,43,37,42,58,49,35,44,53,38,41,37,40,27,40,46,35,51,48,46,38,41,52,36,49,52,35,40,43,34,46,42,56,53,40,54,47,45,63,55,57,58,50,57,52,49,48,51,49,56,35,40,49,53,40,41,40,70,61,41,61,75,48,64,117,87,53,48,56,41,45,59,56,60,51,44,45,46,38,48,48,43,41,42,56,36,41,62,42,39,63,34,35,39,44,60,42,60,49,41,32,43,48,44,44,44,42,33,32,45,51,41,36,49,44,39,40,32,51,43,45,58,47,59,36,53,54,45,64,45,35,61,42,40,47,47,44,49,49,51,43,51,58,44,44,48,38,59,42,50,50,48,64,49,49,39,47,56,30,33,40,56,53,35,39,76,45,43,59,44,44,51,35,42,42,44,52,46,52,50,53,40,35,48,46,47,36,40,47,43,42,57}', 115.4816, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 21:18:29+00', 1, 0, -16898.158203125, 819.7639770507812, '{5,4,8,9,10,9,11,8,4,3,7,6,1,2,1,1,1,7,6,5,7,6,4,3,3,3,4,2,3,3,3,2,3,3,4,4,3,2,3,7,11,8,2,1,1,2,2,1,1,3,5,4,4,2,2,4,2,4,3,8,6,3,2,3,3,3,4,4,3,6,3,4,4,4,5,3,2,5,5,3,5,7,6,8,1,1,2,2,6,7,7,13,8,7,4,4,6,6,8,7,5,5,4,5,2,2,6,4,5,5,4,2,4,9,8,7,13,16,10,8,7,5,8,11,32,38,61,71,55,40,15,8,5,7,5,5,11,10,13,7,5,9,7,5,3,5,6,2,4,3,6,5,3,4,6,11,11,10,7,3,4,5,6,6,4,2,3,2,4,6,5,3,3,3,1,2,1,2,9,14,8,3,6,5,6,5,5,6,6,6,5,4,7,9,7,6,10,6,7,4,1,3,1,2,2,3,3,4,5,4,7,2,3,4,7,3,3,2,1,2,4,2,3,3,4,3,8,10,12,5,7,5,3,4,7,6,3,3,3,5,3,2,4,3,3,2,2,4,2,6,4,7,6,7,11}', 407.06595, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 21:18:51+00', 1, 2, -16954.236328125, 2192.45947265625, '{46,35,43,57,50,59,45,45,47,41,39,46,48,43,51,54,49,33,65,59,47,54,43,53,43,31,50,56,38,45,45,32,30,46,50,36,43,41,47,49,44,46,67,49,49,48,56,56,52,44,53,41,52,42,48,45,40,34,54,40,37,54,56,45,50,48,48,60,42,30,39,44,44,36,44,42,53,53,48,53,56,64,53,43,51,66,50,42,45,42,50,45,39,49,51,41,50,43,38,42,42,42,48,50,63,41,51,52,52,53,44,38,47,50,42,36,35,36,49,42,58,61,47,56,74,84,114,123,42,55,83,56,37,41,47,55,40,47,60,38,31,64,62,51,60,52,41,47,39,47,47,48,52,51,41,48,43,33,42,39,41,39,39,54,43,43,55,67,51,48,39,39,51,43,36,43,56,47,34,49,42,39,39,50,34,45,44,42,36,36,61,46,43,50,38,32,40,47,57,53,44,47,46,34,41,53,41,31,46,69,40,65,49,48,43,56,56,58,56,45,48,53,39,44,30,31,36,49,34,42,39,33,54,46,39,38,50,37,38,53,63,45,47,47,43,37,42,45,43,53,50,44,55,48,43}', 114.69729, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 21:24:06+00', 1, 0, -26599.84375, 191.8125, '{26,24,22,21,27,21,9,14,17,15,17,22,18,16,24,27,34,29,25,29,31,27,24,18,18,23,21,19,15,15,17,13,20,19,22,23,23,18,16,17,21,19,25,15,17,11,14,14,25,22,21,29,24,20,24,18,24,26,22,16,18,16,20,16,17,24,20,25,29,25,17,23,21,30,25,24,21,24,29,31,29,20,22,17,15,19,18,25,26,27,34,28,29,34,31,37,35,30,24,24,18,20,26,24,26,34,29,33,29,22,23,24,26,30,25,26,18,15,12,15,22,29,36,42,60,82,94,95,73,46,39,33,28,24,24,25,32,29,29,30,20,29,31,27,26,25,25,17,16,17,19,21,29,28,34,32,29,22,31,27,26,19,16,21,16,24,20,17,18,25,26,23,27,25,24,16,20,17,17,15,13,20,33,33,31,30,16,20,23,19,23,21,18,21,29,25,26,25,26,21,27,18,19,19,22,20,21,21,23,28,25,25,24,23,18,25,26,18,10,11,16,24,21,19,21,20,26,22,17,18,12,8,14,18,19,18,21,21,20,26,32,26,36,34,48,42,24,24,14,22,30,28,37,37,29}', 410.02444, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 21:32:46+00', 1, 2, -1099.8326416015625, 3306.847412109375, '{55,50,63,73,54,56,63,57,47,49,53,55,52,52,58,71,55,46,48,47,38,38,43,57,47,51,53,44,57,56,49,51,43,58,56,51,43,49,50,44,55,35,37,44,38,62,48,45,47,52,56,47,57,60,41,44,55,47,59,59,59,53,47,49,53,27,42,51,47,57,45,70,71,50,47,66,51,51,58,43,54,55,51,68,75,68,51,47,65,53,49,55,53,48,55,55,59,44,40,50,54,50,46,63,67,54,67,70,78,54,55,47,53,53,53,52,44,44,65,70,60,51,47,64,78,57,109,129,66,72,84,53,47,59,64,55,54,60,46,45,51,38,53,46,48,60,52,48,60,46,50,69,53,49,53,55,52,44,45,55,52,38,50,51,43,61,44,53,48,46,53,47,41,64,62,46,46,54,46,48,45,59,44,36,38,47,63,57,44,40,42,60,48,46,68,45,42,51,50,54,52,42,45,38,30,54,45,38,52,54,49,59,48,46,57,64,56,50,52,58,52,42,40,33,57,64,60,58,61,62,60,47,64,66,62,46,59,64,68,64,49,51,50,51,51,53,49,57,57,43,55,41,59,50,45}', 114.14867, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 21:34:28+00', 1, 2, -965.5186157226562, 3376.070556640625, '{25,30,41,27,29,34,30,37,36,50,41,32,43,44,36,36,34,34,37,37,32,36,39,27,23,32,36,25,35,26,30,46,46,32,24,34,25,36,32,41,36,46,36,35,41,33,38,36,34,30,35,40,34,48,42,51,40,34,48,46,29,35,39,46,34,41,53,36,25,38,41,31,41,47,28,41,41,31,32,34,31,43,36,36,37,30,27,41,39,39,30,31,28,34,40,35,44,52,29,32,39,33,34,41,40,26,27,38,37,33,39,34,43,27,33,35,31,32,45,39,28,32,29,36,47,26,54,113,37,35,86,41,35,37,44,30,32,38,37,22,37,30,27,33,40,35,36,39,29,29,27,37,41,37,33,36,39,25,26,38,37,33,30,34,37,35,37,42,34,33,32,34,34,42,44,37,32,17,29,34,36,32,56,57,35,53,42,27,44,36,33,31,21,27,30,32,32,32,42,36,32,30,30,24,37,35,25,31,49,37,39,46,45,38,33,42,35,25,27,32,45,46,34,35,34,25,27,44,49,45,43,25,32,32,22,27,25,17,29,34,34,36,39,39,30,41,27,27,37,34,30,41,46,32,34}', 115.54247, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 21:39:17+00', 1, 2, -9378.4677734375, 1695.4520263671875, '{35,30,44,27,29,36,32,31,27,22,34,41,32,32,28,34,37,36,42,29,33,47,26,35,37,35,54,24,37,40,43,27,30,30,31,36,36,34,36,35,33,36,24,30,25,28,34,48,21,37,35,34,21,40,37,43,37,27,49,51,38,34,47,41,37,40,37,26,28,30,23,28,29,28,28,30,28,36,35,39,54,36,40,35,35,37,44,24,28,37,35,32,39,36,41,42,38,48,52,33,30,42,40,32,29,45,32,28,27,33,41,41,30,35,34,30,31,40,49,41,32,29,43,28,31,46,72,110,46,38,42,38,25,37,22,36,35,26,31,44,31,31,28,33,27,36,27,20,35,39,22,35,49,30,29,41,45,40,38,30,35,32,38,40,33,29,37,30,26,36,39,32,28,29,36,36,37,36,31,43,40,40,43,36,29,26,42,41,37,46,40,34,31,22,16,30,42,25,37,37,37,42,32,37,33,34,30,39,49,51,19,39,49,34,38,33,33,35,34,41,41,44,37,26,18,28,39,49,39,39,45,31,36,43,37,38,45,28,32,33,33,25,36,29,45,43,32,36,42,49,35,40,61,31,29}', 114.9908, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 21:40:57+00', 1, 2, -9217.2998046875, 2293.3955078125, '{35,44,47,49,48,58,49,35,46,52,50,41,37,66,71,53,50,47,50,40,42,60,41,58,50,49,53,48,44,39,29,42,35,55,43,54,48,43,46,59,43,53,37,38,48,39,54,43,40,45,57,46,43,51,49,51,39,43,46,60,37,45,50,46,41,61,44,45,45,41,47,46,35,37,42,34,42,42,42,47,46,43,50,42,39,46,51,39,56,56,51,53,40,39,45,39,52,52,54,53,54,56,46,54,55,57,49,50,61,59,52,55,43,46,50,54,35,39,38,42,48,38,51,60,65,62,83,125,98,62,51,56,50,43,36,43,61,46,44,42,54,52,45,35,43,44,44,52,60,46,46,47,41,54,47,38,49,52,38,39,58,45,41,58,45,44,52,58,68,46,45,41,53,48,53,56,62,46,36,54,62,64,44,38,63,37,36,45,48,41,60,52,45,65,49,58,44,43,54,66,43,44,49,56,48,36,49,35,40,35,54,41,36,36,29,33,52,53,44,58,53,54,44,48,60,46,39,39,37,40,46,40,40,54,43,50,42,62,60,49,51,43,33,47,37,35,55,42,55,52,47,54,51,46,40}', 113.31388, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 21:44:33+00', 1, 0, -26918.36328125, 97.0145034790039, '{27,25,29,30,35,36,29,28,28,39,43,35,32,33,35,24,25,24,27,21,22,26,31,32,27,38,38,29,29,41,44,35,37,41,47,62,52,41,33,22,32,30,28,27,31,38,33,29,28,25,32,46,44,41,33,33,19,21,29,29,23,31,29,29,17,17,19,22,16,20,30,24,35,31,32,36,41,44,46,48,45,40,20,13,24,24,30,31,25,27,34,30,29,31,31,26,34,26,27,30,21,21,26,32,31,28,27,25,26,24,30,34,40,32,26,28,37,40,34,21,22,24,34,59,65,98,112,113,100,61,56,32,30,27,34,37,34,38,36,33,35,26,26,26,21,32,26,27,27,31,38,32,30,20,17,24,31,27,28,39,28,35,26,24,25,34,39,33,28,24,23,22,32,26,23,28,32,29,42,47,38,45,36,36,24,31,26,20,26,17,23,26,29,28,20,30,27,34,35,34,37,27,30,30,32,32,27,29,23,29,39,40,34,26,21,20,14,17,25,29,26,32,23,27,23,24,29,37,33,42,32,30,26,23,26,20,24,26,23,29,36,25,29,32,27,25,30,30,30,31,33,36,35,34,36}', 410.84177, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 21:51:27+00', 1, 0, -23592.07421875, 742.5275268554688, '{21,22,26,23,23,19,13,18,21,19,24,20,29,23,22,25,21,21,19,21,24,25,26,27,27,22,24,34,19,24,24,20,23,29,33,27,21,22,18,14,14,20,23,31,35,38,36,27,26,21,23,24,26,25,29,39,32,28,32,26,23,17,25,21,29,31,36,36,36,34,24,22,23,30,36,28,35,37,35,33,24,32,29,35,47,41,38,33,30,22,31,38,36,43,36,37,44,37,25,23,28,22,18,23,21,27,22,31,36,31,37,40,34,30,25,24,26,26,27,29,30,24,24,31,55,82,97,100,88,64,46,37,26,29,24,22,28,26,20,22,26,18,26,21,16,29,31,23,20,14,19,25,26,26,21,27,18,22,25,20,25,24,23,32,29,21,21,26,29,28,29,27,17,17,19,16,19,26,24,25,11,13,12,13,17,15,16,21,32,38,34,33,25,17,24,27,22,23,16,14,13,16,20,21,24,26,29,23,27,22,24,26,31,32,18,15,20,24,23,25,27,25,18,18,27,24,27,22,20,17,16,19,19,24,19,25,24,22,16,11,19,27,28,25,25,22,27,26,21,19,31,28,19,16,16}', 422.9308, 1);
INSERT INTO _timescaledb_internal._hyper_3_2_chunk VALUES ('2014-08-16 21:53:08+00', 1, 2, -6394.91064453125, 553.9545288085938, '{59,61,58,63,62,70,50,48,55,57,51,51,55,58,63,84,56,47,74,60,54,56,51,64,52,40,57,50,44,55,49,53,49,51,42,64,52,45,55,60,59,79,61,54,58,55,44,47,49,55,56,45,38,54,64,49,57,58,54,66,65,49,66,68,51,42,59,55,50,58,51,51,51,46,36,51,61,37,40,56,58,50,57,54,47,53,62,57,58,62,52,57,45,34,52,39,61,61,50,53,65,59,64,51,57,57,41,58,75,48,43,52,62,49,49,52,37,50,60,54,42,60,65,65,90,57,67,131,116,67,67,63,65,64,79,67,63,62,59,53,54,42,57,78,65,58,52,75,58,62,60,48,52,69,58,71,78,54,61,56,49,52,84,62,44,51,66,60,47,72,63,63,62,55,60,76,63,51,71,72,57,42,48,51,65,53,43,54,58,47,62,73,60,70,57,56,55,46,51,47,45,45,50,63,52,49,56,60,49,42,54,47,54,41,53,54,48,43,61,58,49,53,36,58,61,70,46,61,63,57,63,77,62,47,46,53,53,51,51,63,65,61,55,59,45,55,56,65,67,61,83,73,56,65,71}', 116.65158, 1);


--
-- Data for Name: acq; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: application; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: experiment; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.experiment VALUES (1, 1, 1, NULL, NULL);


--
-- Data for Name: experiment_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.experiment_type VALUES (2, 'Sample capture');
INSERT INTO public.experiment_type VALUES (3, 'Software update');
INSERT INTO public.experiment_type VALUES (1, 'Real time processing');


--
-- Data for Name: log; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: log_priority; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: mission_operation; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.mission_operation VALUES (1, 1, NULL, NULL);
INSERT INTO public.mission_operation VALUES (2, 2, NULL, NULL);


--
-- Data for Name: nav; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: operation_modes; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.operation_modes VALUES (1, 1);
INSERT INTO public.operation_modes VALUES (2, 2);
INSERT INTO public.operation_modes VALUES (3, 3);


--
-- Data for Name: pod; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: raw; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.schema_migrations VALUES (20240226122724, false);


--
-- Data for Name: signal; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.signal VALUES (1, 'signal1');
INSERT INTO public.signal VALUES (2, 'signal2');
INSERT INTO public.signal VALUES (3, 'signal3');
INSERT INTO public.signal VALUES (0, 'signal0');


--
-- Data for Name: sta; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Name: chunk_constraint_name; Type: SEQUENCE SET; Schema: _timescaledb_catalog; Owner: postgres
--

SELECT pg_catalog.setval('_timescaledb_catalog.chunk_constraint_name', 4, true);


--
-- Name: chunk_id_seq; Type: SEQUENCE SET; Schema: _timescaledb_catalog; Owner: postgres
--

SELECT pg_catalog.setval('_timescaledb_catalog.chunk_id_seq', 2, true);


--
-- Name: continuous_agg_migrate_plan_step_step_id_seq; Type: SEQUENCE SET; Schema: _timescaledb_catalog; Owner: postgres
--

SELECT pg_catalog.setval('_timescaledb_catalog.continuous_agg_migrate_plan_step_step_id_seq', 1, false);


--
-- Name: dimension_id_seq; Type: SEQUENCE SET; Schema: _timescaledb_catalog; Owner: postgres
--

SELECT pg_catalog.setval('_timescaledb_catalog.dimension_id_seq', 6, true);


--
-- Name: dimension_slice_id_seq; Type: SEQUENCE SET; Schema: _timescaledb_catalog; Owner: postgres
--

SELECT pg_catalog.setval('_timescaledb_catalog.dimension_slice_id_seq', 2, true);


--
-- Name: hypertable_id_seq; Type: SEQUENCE SET; Schema: _timescaledb_catalog; Owner: postgres
--

SELECT pg_catalog.setval('_timescaledb_catalog.hypertable_id_seq', 6, true);


--
-- Name: bgw_job_id_seq; Type: SEQUENCE SET; Schema: _timescaledb_config; Owner: postgres
--

SELECT pg_catalog.setval('_timescaledb_config.bgw_job_id_seq', 1000, false);


--
-- Name: application_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.application_id_seq', 1, false);


--
-- Name: experiment_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.experiment_id_seq', 1, false);


--
-- Name: experiment_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.experiment_type_id_seq', 1, false);


--
-- Name: log_priority_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.log_priority_id_seq', 1, false);


--
-- Name: mission_operation_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mission_operation_id_seq', 1, false);


--
-- Name: operation_modes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.operation_modes_id_seq', 1, false);


--
-- Name: signal_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.signal_id_seq', 1, false);


--
-- Name: application application_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application
    ADD CONSTRAINT application_pkey PRIMARY KEY (id);


--
-- Name: experiment experiment_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.experiment
    ADD CONSTRAINT experiment_pkey PRIMARY KEY (id);


--
-- Name: experiment_type experiment_type_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.experiment_type
    ADD CONSTRAINT experiment_type_pkey PRIMARY KEY (id);


--
-- Name: log_priority log_priority_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.log_priority
    ADD CONSTRAINT log_priority_pkey PRIMARY KEY (id);


--
-- Name: mission_operation mission_operation_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mission_operation
    ADD CONSTRAINT mission_operation_pkey PRIMARY KEY (id);


--
-- Name: operation_modes operation_modes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.operation_modes
    ADD CONSTRAINT operation_modes_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: signal signal_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.signal
    ADD CONSTRAINT signal_pkey PRIMARY KEY (id);


--
-- Name: _hyper_3_2_chunk_acq_rx_time_idx; Type: INDEX; Schema: _timescaledb_internal; Owner: postgres
--

CREATE INDEX _hyper_3_2_chunk_acq_rx_time_idx ON _timescaledb_internal._hyper_3_2_chunk USING btree (rx_time DESC);


--
-- Name: acq_rx_time_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX acq_rx_time_idx ON public.acq USING btree (rx_time DESC);


--
-- Name: log_rx_time_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX log_rx_time_idx ON public.log USING btree (rx_time DESC);


--
-- Name: nav_rx_time_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX nav_rx_time_idx ON public.nav USING btree (rx_time DESC);


--
-- Name: pod_rx_time_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX pod_rx_time_idx ON public.pod USING btree (rx_time DESC);


--
-- Name: raw_rx_time_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX raw_rx_time_idx ON public.raw USING btree (rx_time DESC);


--
-- Name: sta_rx_time_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX sta_rx_time_idx ON public.sta USING btree (rx_time DESC);


--
-- Name: acq ts_insert_blocker; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER ts_insert_blocker BEFORE INSERT ON public.acq FOR EACH ROW EXECUTE FUNCTION _timescaledb_functions.insert_blocker();


--
-- Name: log ts_insert_blocker; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER ts_insert_blocker BEFORE INSERT ON public.log FOR EACH ROW EXECUTE FUNCTION _timescaledb_functions.insert_blocker();


--
-- Name: nav ts_insert_blocker; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER ts_insert_blocker BEFORE INSERT ON public.nav FOR EACH ROW EXECUTE FUNCTION _timescaledb_functions.insert_blocker();


--
-- Name: pod ts_insert_blocker; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER ts_insert_blocker BEFORE INSERT ON public.pod FOR EACH ROW EXECUTE FUNCTION _timescaledb_functions.insert_blocker();


--
-- Name: raw ts_insert_blocker; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER ts_insert_blocker BEFORE INSERT ON public.raw FOR EACH ROW EXECUTE FUNCTION _timescaledb_functions.insert_blocker();


--
-- Name: sta ts_insert_blocker; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER ts_insert_blocker BEFORE INSERT ON public.sta FOR EACH ROW EXECUTE FUNCTION _timescaledb_functions.insert_blocker();


--
-- Name: _hyper_3_2_chunk 2_4_acq_signal_id_fkey; Type: FK CONSTRAINT; Schema: _timescaledb_internal; Owner: postgres
--

ALTER TABLE ONLY _timescaledb_internal._hyper_3_2_chunk
    ADD CONSTRAINT "2_4_acq_signal_id_fkey" FOREIGN KEY (signal_id) REFERENCES public.signal(id);


--
-- Name: experiment experiment_experiment_type_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.experiment
    ADD CONSTRAINT experiment_experiment_type_id_fkey FOREIGN KEY (experiment_type_id) REFERENCES public.experiment_type(id);


--
-- Name: experiment experiment_mission_operation_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.experiment
    ADD CONSTRAINT experiment_mission_operation_id_fkey FOREIGN KEY (mission_operation_id) REFERENCES public.mission_operation(id);


--
-- PostgreSQL database dump complete
--

