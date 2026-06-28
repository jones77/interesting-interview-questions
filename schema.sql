CREATE TABLE IF NOT EXISTS questions (
    id            INTEGER PRIMARY KEY,
    short_name    TEXT NOT NULL UNIQUE,
    title         TEXT NOT NULL,
    question_text TEXT NOT NULL,
    solution      TEXT NOT NULL,
    hint          TEXT,
    difficulty    TEXT NOT NULL DEFAULT 'medium',
    category      TEXT NOT NULL DEFAULT 'general'
);
