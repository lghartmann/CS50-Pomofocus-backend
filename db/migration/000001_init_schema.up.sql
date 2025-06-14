CREATE TABLE
    IF NOT EXISTS pomodoro (
        id SERIAL PRIMARY KEY,
        user_id TEXT NOT NULL,
        duration TEXT NOT NULL,
        pause_duration TEXT NOT NULL,
        effort NUMERIC(3, 1) NOT NULL CHECK(effort BETWEEN 0.0 AND 10.0),
        distraction NUMERIC(3, 1) NOT NULL CHECK(distraction BETWEEN 0.0 AND 10.0),
        productivity NUMERIC(3, 1) NOT NULL CHECK(productivity BETWEEN 0.0 AND 10.0),
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        deleted_at TIMESTAMPTZ DEFAULT NULL
    );