CREATE TABLE
  IF NOT EXISTS gem_source (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    source_id VARCHAR NOT NULL,
    gems FLOAT NOT NULL,
    type VARCHAR(30) NOT NULL,
    status VARCHAR(30) NOT NULL,
    reason TEXT,
    metadata TEXT,
    collected_at TIMESTAMP DEFAULT (NOW () AT TIME ZONE 'utc'),
    created_by VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT (NOW () AT TIME ZONE 'utc'),
    updated_by VARCHAR(100),
    updated_at TIMESTAMP
  );

CREATE INDEX IF NOT EXISTS gem_source_source_id_idx ON gem_source (source_id);

CREATE INDEX IF NOT EXISTS gem_source_user_id_idx ON gem_source (user_id);
