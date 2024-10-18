CREATE TABLE IF NOT EXISTS "user" (
    id bigint generated always as identity primary key,
    name text,
    email text,
    password text,
    is_active boolean not null default true,
    created_at timestamptz not null default now()::timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz
)