CREATE TABLE IF NOT EXISTS activity_records (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL DEFAULT 1,
    source VARCHAR(32) NOT NULL DEFAULT 'manual',
    activity_type VARCHAR(64) NOT NULL,
    steps INT NOT NULL DEFAULT 0,
    minutes INT NOT NULL DEFAULT 0,
    intensity VARCHAR(32) NOT NULL DEFAULT 'moderate',
    cardio_points INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

