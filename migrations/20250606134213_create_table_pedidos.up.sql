CREATE TABLE pedidos (
    id UUID PRIMARY KEY,
    cliente_nome TEXT NOT NULL,
    endereco_entrega TEXT NOT NULL,
    status TEXT NOT NULL,
    criado_em TIMESTAMP NOT NULL DEFAULT now()
);