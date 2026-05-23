CREATE TABLE IF NOT EXISTS professional_dossier (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(200) NOT NULL,
    paragraph_1 TEXT NOT NULL,
    paragraph_2 TEXT,
    paragraph_3 TEXT,
    is_published BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
