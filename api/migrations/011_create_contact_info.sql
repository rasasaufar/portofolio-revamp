CREATE TABLE IF NOT EXISTS contact_info (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255),
    phone VARCHAR(50),
    whatsapp_url VARCHAR(500),
    github_url VARCHAR(500),
    linkedin_url VARCHAR(500),
    instagram_url VARCHAR(500),
    location VARCHAR(200),
    contact_description TEXT,
    is_published BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
