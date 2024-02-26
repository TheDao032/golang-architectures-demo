CREATE TABLE
  IF NOT EXISTS gem_history (
    id UUID NOT NULL,
    user_id UUID NOT NULL,
    source_id VARCHAR NOT NULL,
    gems FLOAT NOT NULL,
    type VARCHAR NOT NULL,
    status VARCHAR NOT NULL,
    reason TEXT,
    metadata TEXT,
    collected_at TIMESTAMP DEFAULT (NOW () AT TIME ZONE 'utc'),
    created_by VARCHAR(100),
    created_at TIMESTAMP DEFAULT (NOW () AT TIME ZONE 'utc')
  );
