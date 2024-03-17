CREATE TABLE IF NOT EXISTS pagamentos (
    id SERIAL PRIMARY KEY,
    pedido_id VARCHAR(255) NOT NULL,
    cliente_id bigserial NOT NULL,
    amount NUMERIC(10, 2) NOT NULL,
    status TEXT NOT NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL
    );
CREATE INDEX idx_pagamentos_deleted_at ON pagamentos USING btree (deleted_at);