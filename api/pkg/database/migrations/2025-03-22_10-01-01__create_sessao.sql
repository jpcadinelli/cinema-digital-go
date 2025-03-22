CREATE TABLE sessao (
                        id              UUID           PRIMARY KEY,
                        id_filme        UUID           NOT NULL   ,
                        id_sala         UUID           NOT NULL   ,
                        data_inicio     TIMESTAMP      NOT NULL   ,
                        data_fim        TIMESTAMP      NOT NULL   ,
                        preco_ingresso  NUMERIC(10, 2) NOT NULL   ,
                        disponibilidade SMALLINT       NOT NULL   ,
                        CONSTRAINT fk_sessao_filme FOREIGN KEY (id_filme) REFERENCES filme(id),
                        CONSTRAINT fk_sessao_sala  FOREIGN KEY (id_sala)  REFERENCES sala(id)
);

COMMENT ON COLUMN sessao.id              IS 'Identificador único da sessão de cinema'         ;
COMMENT ON COLUMN sessao.id_filme        IS 'Referência para o filme que será exibido'        ;
COMMENT ON COLUMN sessao.id_sala         IS 'Referência para a sala onde o filme será exibido';
COMMENT ON COLUMN sessao.data_inicio     IS 'Data e hora do início da exibição do filme'      ;
COMMENT ON COLUMN sessao.data_fim        IS 'Data e hora do término da exibição do filme'     ;
COMMENT ON COLUMN sessao.preco_ingresso  IS 'Valor do ingresso para essa sessão'              ;
COMMENT ON COLUMN sessao.disponibilidade IS 'Status da sessão: 1 - Disponível; 2 - Cancelada.';