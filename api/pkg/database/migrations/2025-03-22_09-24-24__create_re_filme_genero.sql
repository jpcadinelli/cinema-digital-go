CREATE TABLE re_filme_genero (
                                   id UUID,
                                   id_filme UUID,
                                   id_genero UUID,
                                   CONSTRAINT pk_re_filme_genero PRIMARY KEY (id),
                                   CONSTRAINT fk_re_filme_genero_filme FOREIGN KEY (id_filme) REFERENCES filme(id),
                                   CONSTRAINT fk_re_filme_genero_genero FOREIGN KEY (id_genero) REFERENCES genero(id)
);

COMMENT ON COLUMN re_filme_genero.id IS 'Identificador único da tabela re_filme_genero';
COMMENT ON COLUMN re_filme_genero.id_filme IS 'Referência para o filme na tabela filme';
COMMENT ON COLUMN re_filme_genero.id_genero IS 'Referência para o gênero na tabela genero';