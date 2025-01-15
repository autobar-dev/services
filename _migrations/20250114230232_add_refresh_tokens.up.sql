CREATE TABLE refresh_tokens (
    id SERIAL PRIMARY KEY,

    user_id INTEGER DEFAULT NULL REFERENCES users(id) ON DELETE CASCADE,
    module_id INTEGER DEFAULT NULL REFERENCES modules(id) ON DELETE CASCADE,
    token TEXT NOT NULL,
    remember_me BOOLEAN NOT NULL DEFAULT FALSE,
    expires_at TIMESTAMPTZ NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,

    CONSTRAINT xor_user_id_module_id CHECK (
        (user_id IS NOT NULL AND module_id IS NULL) OR
        (user_id IS NULL AND module_id IS NOT NULL)
    )
);