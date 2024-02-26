CREATE TABLE
  IF NOT EXISTS gem_dashboard (
    user_id UUID PRIMARY KEY,
    id UUID NOT NULL,
    pending FLOAT NOT NULL,
    redeemable FLOAT NOT NULL,
    redeem_limitation FLOAT DEFAULT NULL,
    redeemed FLOAT NOT NULL,
    status VARCHAR,
    created_by VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT (NOW () AT TIME ZONE 'utc'),
    updated_by VARCHAR(100),
    updated_at TIMESTAMP NOT NULL DEFAULT (NOW () AT TIME ZONE 'utc')
  );

CREATE INDEX IF NOT EXISTS gem_dashboard_id_idx ON gem_dashboard (user_id);
