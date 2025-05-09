CREATE TABLE
    IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        email TEXT UNIQUE NOT NULL,
        password TEXT NOT NULL,
        birth_date DATE NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS user_auth_token (
        id SERIAL PRIMARY KEY,
        user_id INTEGER NOT NULL REFERENCES users(id),
        token_hash TEXT NOT NULL,
        token_type VARCHAR(20) NOT NULL CHECK(token_type IN ('access', 'refresh', 'password_reset', 'email_verification')),
        device_info TEXT DEFAULT NULL,
        ip_address INET DEFAULT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        expires_at TIMESTAMPTZ NOT NULL,
        revoked BOOLEAN NOT NULL DEFAULT FALSE,
        last_used_at TIMESTAMPTZ DEFAULT NULL
    );

CREATE TABLE
    IF NOT EXISTS pomodoro (
        id SERIAL PRIMARY KEY,
        user_id INTEGER REFERENCES users(id) NOT NULL,
        duration TEXT NOT NULL,
        pause_duration TEXT NOT NULL,
        effort NUMERIC(3, 1) NOT NULL CHECK(effort BETWEEN 0.0 AND 10.0),
        distraction NUMERIC(3, 1) NOT NULL CHECK(distraction BETWEEN 0.0 AND 10.0),
        productivity NUMERIC(3, 1) NOT NULL CHECK(productivity BETWEEN 0.0 AND 10.0),
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        deleted_at TIMESTAMPTZ DEFAULT NULL
    );