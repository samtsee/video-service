CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    user_id UUID,
    title TEXT NOT NULL,
    description TEXT,
    location JSONB,
    video_id TEXT,
    created TIMESTAMPTZ DEFAULT NOW(),
    edited TIMESTAMPTZ
);
