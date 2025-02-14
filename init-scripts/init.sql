-- Criar o banco de dados de testes
CREATE DATABASE test;

-- Criar o usuário de testes apenas se ele não existir

-- \c blacklist;

CREATE SCHEMA IF NOT EXISTS public;

-- Criar tabela exemplo dentro do banco blacklist
CREATE TABLE IF NOT EXISTS public.blacklist(
    id TEXT PRIMARY KEY,
    event_id TEXT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    reason TEXT NOT NULL,
    document TEXT NOT NULL,
    scope TEXT NOT NULL,
    user_identifier INT NOT NULL,
    blocked_until TIMESTAMP NULL,
    blocked_type TEXT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS public.auditlog (
		id TEXT PRIMARY KEY,
		event_id TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		user_identifier INT NOT NULL,
		action TEXT NOT NULL,
		details TEXT NOT NULL
);
