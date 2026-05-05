CREATE TABLE IF NOT EXISTS daily_scores (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL DEFAULT 1,
    score_date DATE NOT NULL,
    diet_score INT NOT NULL DEFAULT 0,
    activity_score INT NOT NULL DEFAULT 0,
    sleep_score INT NOT NULL DEFAULT 0,
    total_score INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (user_id, score_date)
);

