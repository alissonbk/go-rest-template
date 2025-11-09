CREATE TABLE IF NOT EXISTS "example" (
    id bigint generated always as identity primary key,
    name text,    
    created_at timestamptz not null default now()::timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz
)