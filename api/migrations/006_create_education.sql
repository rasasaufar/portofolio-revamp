CREATE TABLE IF NOT EXISTS education (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    institution_name VARCHAR(200) NOT NULL,
    degree VARCHAR(200) NOT NULL,
    major VARCHAR(200),
    start_year VARCHAR(10),
    end_year VARCHAR(10),
    gpa VARCHAR(20),
    description TEXT,
    image_url VARCHAR(500),
    tags JSONB NOT NULL DEFAULT '[]',
    order_number INTEGER NOT NULL DEFAULT 0,
    is_published BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
