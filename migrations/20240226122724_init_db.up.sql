CREATE TABLE operation_modes (
    id SERIAL PRIMARY KEY,
    rx_operation_mode_name SMALLINT
);
---
CREATE TABLE application (
    id SERIAL PRIMARY KEY,
    application_name SMALLINT
);
---
CREATE TABLE log_priority (
    id SERIAL PRIMARY KEY,
    log_priority_name SMALLINT
);
---
CREATE TABLE mission_operation (
    id SERIAL PRIMARY KEY,
    mission_operation_name INTEGER,
    start_time TIMESTAMP,
    end_time TIMESTAMP
);
---
CREATE TABLE signal (
    id SERIAL PRIMARY KEY,
    signal_name VARCHAR(8)
);
---
CREATE TABLE experiment_type (
    id SERIAL PRIMARY KEY,
    experiment_type_name VARCHAR(32)
);
---
CREATE table experiment (
    id SERIAL PRIMARY KEY,
    mission_operation_id INTEGER REFERENCES mission_operation(id),
    experiment_type_id INTEGER REFERENCES experiment_type(id),
    start_time TIMESTAMP,
    end_time TIMESTAMP
);
---
CREATE table log (
    rx_time TIMESTAMPTZ NOT NULL,
    experiment_type_id INTEGER REFERENCES experiment_type(id),
    priority INTEGER REFERENCES log_priority(id),
    size INTEGER,
    message TEXT
);
SELECT create_hypertable('log', 'rx_time'); 
---
CREATE table raw (
    rx_time TIMESTAMPTZ NOT NULL,
    experiment_id INTEGER REFERENCES experiment(id),
    signal_id INTEGER REFERENCES signal(id),
    sv_id INTEGER,
    fd_raw DOUBLE PRECISION,
    fd_rate_raw DOUBLE PRECISION,
    carrier_phase DOUBLE PRECISION,
    pr_raw DOUBLE PRECISION
);
SELECT create_hypertable('raw', 'rx_time'); 
---
CREATE table acq (
    rx_time TIMESTAMPTZ NOT NULL,
    experiment_id INTEGER REFERENCES experiment(id),
    signal_id INTEGER REFERENCES signal(id),
    doppler DOUBLE PRECISION,
    code_phase DOUBLE PRECISION,
    acf_corr REAL[],
    noise_floor REAL,
    acq_mode SMALLINT
);
SELECT create_hypertable('acq', 'rx_time');
---
CREATE table pod (
    rx_time TIMESTAMPTZ NOT NULL,
    experiment_id INTEGER REFERENCES experiment(id),
    application_id INTEGER REFERENCES application(id),
    wn INTEGER,
    tow BIGINT,
    decimals DOUBLE PRECISION,
    n_sat INTEGER,
    pos_x DOUBLE PRECISION,
    pos_y DOUBLE PRECISION,
    pos_z DOUBLE PRECISION,
    vel_x DOUBLE PRECISION,
    vel_y DOUBLE PRECISION,
    vel_z DOUBLE PRECISION,
    pos_std REAL[],
    vel_std REAL[],
    clock_bias DOUBLE PRECISION,
    clock_drift DOUBLE PRECISION,
    ambig_vec REAL[],
    ambig_acc REAL[]
);
SELECT create_hypertable('pod', 'rx_time');
---
CREATE table sta (
    rx_time TIMESTAMPTZ NOT NULL,
    experiment_id INTEGER REFERENCES experiment(id),
    rx_operation_mode_id INTEGER REFERENCES operation_modes(id),
    application_id INTEGER REFERENCES application(id),
    signal_id INTEGER REFERENCES signal(id),
    number_of_channels SMALLINT,
    sv_id INTEGER,
    channel_status SMALLINT,
    number_of_apps SMALLINT,
    ecc_error_count BIGINT,
    cpu_temp REAL,
    frontend_temp REAL,
    qn400_version_number BIGINT
);
SELECT create_hypertable('sta', 'rx_time');
---
CREATE table nav(
    rx_time TIMESTAMPTZ NOT NULL,
    experiment_id INTEGER REFERENCES experiment(id),
    application_id INTEGER REFERENCES application(id),
    wn INTEGER,
    tow BIGINT,
    decimals DOUBLE PRECISION,
    n_sat INTEGER,
    pos_x DOUBLE PRECISION,
    pos_y DOUBLE PRECISION,
    pos_z DOUBLE PRECISION,
    vel_x DOUBLE PRECISION,
    vel_y DOUBLE PRECISION,
    vel_z DOUBLE PRECISION,
    pos_std REAL,
    vel_std REAL,
    tim_std DOUBLE PRECISION,
    clock_bias DOUBLE PRECISION,
    clock_drift DOUBLE PRECISION,
    ggto DOUBLE PRECISION,
    gdop DOUBLE PRECISION,
    pdop DOUBLE PRECISION,
    hdop DOUBLE PRECISION,
    vdop DOUBLE PRECISION,
    tdop DOUBLE PRECISION
);
SELECT create_hypertable('nav', 'rx_time');
