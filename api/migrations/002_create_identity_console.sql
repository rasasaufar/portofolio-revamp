CREATE TABLE IF NOT EXISTS identity_console (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    role VARCHAR(200) NOT NULL,
    headline TEXT NOT NULL,
    description TEXT NOT NULL,
    avatar_url VARCHAR(500),
    current_focus JSONB NOT NULL DEFAULT '[]',
    availability_text VARCHAR(200),
    cta_primary_label VARCHAR(100),
    cta_primary_link VARCHAR(500),
    cta_secondary_label VARCHAR(100),
    cta_secondary_link VARCHAR(500),
    order_number INTEGER NOT NULL DEFAULT 0,
    is_published BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
