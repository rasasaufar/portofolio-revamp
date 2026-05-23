CREATE TABLE IF NOT EXISTS publications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    journal_name VARCHAR(300) NOT NULL,
    published_date VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'published',
    authors TEXT NOT NULL,
    description TEXT,
    tags JSONB NOT NULL DEFAULT '[]',
    publication_url VARCHAR(500),
    order_number INTEGER NOT NULL DEFAULT 0,
    is_published BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
