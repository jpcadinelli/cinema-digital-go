CREATE TABLE ingresso (
                          id UUID PRIMARY KEY,
                          id_sessao UUID NOT NULL,
                          id_usuario UUID NOT NULL,
                          poltrona VARCHAR(10) NOT NULL,
                          comprado_em TIMESTAMP NOT NULL DEFAULT NOW(),

                          CONSTRAINT fk_ingresso_sessao FOREIGN KEY (id_sessao) REFERENCES sessao(id),
                          CONSTRAINT fk_ingresso_usuario FOREIGN KEY (id_usuario) REFERENCES usuario(id)
);


CREATE INDEX idx_ingresso_id_sessao ON ingresso(id_sessao);
CREATE INDEX idx_ingresso_id_usuario ON ingresso(id_usuario);
