CREATE TABLE IF NOT EXISTS certifications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(300) NOT NULL,
    issuer VARCHAR(200) NOT NULL,
    issued_date VARCHAR(50) NOT NULL,
    expired_date VARCHAR(50),
    credential_id VARCHAR(200),
    credential_url VARCHAR(500),
    description TEXT,
    skills JSONB NOT NULL DEFAULT '[]',
    image_url VARCHAR(500),
    category VARCHAR(50) NOT NULL DEFAULT 'core',
    status VARCHAR(50) NOT NULL DEFAULT 'active',
    order_number INTEGER NOT NULL DEFAULT 0,
    is_published BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
