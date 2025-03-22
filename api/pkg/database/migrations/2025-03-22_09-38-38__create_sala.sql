CREATE TABLE sala (
                                   id        UUID,
                                   nome      VARCHAR(100) NOT NULL ,
                                   fileiras  CHAR(1)      NOT NULL ,
                                   poltronas SMALLINT     NOT NULL ,
                                   CONSTRAINT pk_sala PRIMARY KEY (id)
);

COMMENT ON COLUMN sala.id IS 'Identificador único da sala de cinema (UUID).';
COMMENT ON COLUMN sala.nome IS 'Nome da sala de cinema (VARCHAR, até 100 caracteres).';
COMMENT ON COLUMN sala.fileiras IS 'Letras de fileiras na sala (CHAR(1)).';
COMMENT ON COLUMN sala.poltronas IS 'Número total de poltronas na sala (SMALLINT).';