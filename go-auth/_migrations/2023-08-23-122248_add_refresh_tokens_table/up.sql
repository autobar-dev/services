CREATE TABLE refresh_tokens (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID DEFAULT NULL,
    module_serial_number VARCHAR(6) DEFAULT NULL,
    token TEXT NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,

    CONSTRAINT xor_user_id_module_serial_number CHECK (
        (user_id IS NOT NULL AND module_serial_number IS NULL) OR
        (user_id IS NULL AND module_serial_number IS NOT NULL)
    )
);
