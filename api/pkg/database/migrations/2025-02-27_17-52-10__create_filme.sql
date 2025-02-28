create table filme
(
    id uuid not null
        constraint filme_pk
            primary key,
    titulo varchar(250) not null,
    sinopse text not null,
    diretor varchar(250) not null,
    duracao integer not null,
    ano_lancamento date not null,
    classificacao integer not null,
    nota numeric not null,
    criado timestamp not null,
    atualizado timestamp,
    excluido timestamp,
    id_usuario_registro uuid not null
);

COMMENT ON COLUMN filme.id IS 'Identificador único do filme';
COMMENT ON COLUMN filme.titulo IS 'Título do filme';
COMMENT ON COLUMN filme.sinopse IS 'Sinopse do filme';
COMMENT ON COLUMN filme.diretor IS 'Diretor do filme';
COMMENT ON COLUMN filme.duracao IS 'Duração do filme em minutos';
COMMENT ON COLUMN filme.ano_lancamento IS 'Ano de lançamento do filme (apenas 2024)';
COMMENT ON COLUMN filme.classificacao IS 'Classificação indicativa do filme';
COMMENT ON COLUMN filme.nota IS 'Nota do filme';
COMMENT ON COLUMN filme.criado IS 'Data de criação do registro';
COMMENT ON COLUMN filme.atualizado IS 'Data de atualização do registro';
COMMENT ON COLUMN filme.excluido IS 'Data de exclusão do registro';
COMMENT ON COLUMN filme.id_usuario_registro IS 'Identificador do usuário que registrou o filme';