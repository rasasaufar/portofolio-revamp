CREATE TABLE IF NOT EXISTS work_experiences (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    company_name VARCHAR(200) NOT NULL,
    position VARCHAR(200) NOT NULL,
    start_date VARCHAR(50),
    end_date VARCHAR(50),
    is_current BOOLEAN NOT NULL DEFAULT false,
    description TEXT,
    bullet_points JSONB NOT NULL DEFAULT '[]',
    tech_tags JSONB NOT NULL DEFAULT '[]',
    logo_url VARCHAR(500),
    gallery_images JSONB NOT NULL DEFAULT '[]',
    order_number INTEGER NOT NULL DEFAULT 0,
    is_published BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
