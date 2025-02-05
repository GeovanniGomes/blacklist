-- Criar o banco de dados de testes
CREATE DATABASE test;

-- Criar o usuário de testes apenas se ele não existir
DO $$ 
BEGIN
    IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname = 'test') THEN
        CREATE USER test WITH PASSWORD 'test';
    END IF;
END $$;

CREATE SCHEMA IF NOT EXISTS public;

-- Conceder permissões ao usuário test
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON TABLES TO test;
GRANT USAGE, CREATE ON SCHEMA public TO test;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO test;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO test;
GRANT ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA public TO test;