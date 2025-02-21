CREATE TABLE genero (
    id UUID NOT NULL,
    nome TEXT NOT NULL,
    created_at TIMESTAMP,
        CONSTRAINT pk_genero PRIMARY KEY (id)
);

COMMENT ON COLUMN genero.id IS 'Identificador único do Gênero';
COMMENT ON COLUMN genero.nome IS 'Nome do Gênero';
COMMENT ON COLUMN genero.created_at IS 'Data de Criação';
