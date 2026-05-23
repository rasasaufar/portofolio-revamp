CREATE TABLE IF NOT EXISTS site_settings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    site_title VARCHAR(200) NOT NULL,
    meta_description TEXT,
    favicon_url VARCHAR(500),
    logo_url VARCHAR(500),
    footer_text VARCHAR(300),
    theme_mode VARCHAR(20) NOT NULL DEFAULT 'dark',
    maintenance_mode BOOLEAN NOT NULL DEFAULT false,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
