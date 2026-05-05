CREATE TABLE IF NOT EXISTS sleep_records (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL DEFAULT 1,
    status VARCHAR(32) NOT NULL DEFAULT 'completed',
    started_at TIMESTAMPTZ,
    ended_at TIMESTAMPTZ,
    duration_m INT NOT NULL DEFAULT 0,
    score INT NOT NULL DEFAULT 0,
    audio_url TEXT NOT NULL DEFAULT '',
    advice TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS sleep_events (
    id BIGSERIAL PRIMARY KEY,
    sleep_record_id BIGINT NOT NULL,
    event_type VARCHAR(32) NOT NULL,
    event_timestamp VARCHAR(16) NOT NULL,
    level VARCHAR(16) NOT NULL DEFAULT 'low',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

