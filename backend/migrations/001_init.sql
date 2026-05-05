CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL DEFAULT '',
    age INT NOT NULL DEFAULT 0,
    gender VARCHAR(32) NOT NULL DEFAULT '',
    height_cm INT NOT NULL DEFAULT 0,
    weight_kg INT NOT NULL DEFAULT 0,
    primary_goal VARCHAR(64) NOT NULL DEFAULT 'maintain',
    secondary_goals JSONB NOT NULL DEFAULT '[]'::jsonb,
    health_flags JSONB NOT NULL DEFAULT '[]'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_settings (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL DEFAULT 1 REFERENCES users(id),
    notifications_enabled BOOLEAN NOT NULL DEFAULT true,
    step_permission_granted BOOLEAN NOT NULL DEFAULT false,
    microphone_permission_granted BOOLEAN NOT NULL DEFAULT false,
    sleep_reminder_enabled BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_stats (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL DEFAULT 1 REFERENCES users(id),
    current_streak INT NOT NULL DEFAULT 0,
    days_tracked INT NOT NULL DEFAULT 0,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
