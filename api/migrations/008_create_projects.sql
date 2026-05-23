CREATE TABLE IF NOT EXISTS projects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(200) NOT NULL,
    category VARCHAR(50) NOT NULL,
    description TEXT,
    tech_tags JSONB NOT NULL DEFAULT '[]',
    image_url VARCHAR(500),
    demo_url VARCHAR(500),
    repo_url VARCHAR(500),
    is_featured BOOLEAN NOT NULL DEFAULT false,
    order_number INTEGER NOT NULL DEFAULT 0,
    is_published BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
