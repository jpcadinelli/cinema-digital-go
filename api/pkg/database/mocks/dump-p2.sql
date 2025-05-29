-- Gêneros
INSERT INTO genero (id, nome, created_at) VALUES
                                              ('364a1dfb-6744-42f6-aeb0-cb15b8ff3b94', 'Comédia', NOW()),
                                              ('ab243d57-63b6-4f3b-bb5a-08b3f0b58e44', 'Ficção Científica', NOW()),
                                              ('d8b5d320-57f7-4c9f-b4c7-e03a0b0f2f7f', 'Ação', NOW()),
                                              ('bbb5d321-57f7-4c9f-b4c7-e03a0b0f2f7f', 'Aventura', NOW());

-- Filmes
INSERT INTO filme (id, titulo, sinopse, diretor, duracao, ano_lancamento, classificacao, nota, criado, atualizado, excluido, id_usuario_registro, caminho_poster) VALUES
                                                ('2b0c3fe1-4e29-4fd9-b6f6-cd5c9f88f68a', 'Lilo & Stitch', 'Stitch, um alienígena, chega ao planeta Terra após fugir de sua prisão e tenta se passar por um cachorro para se camuflar. As coisas mudam quando Lilo, uma travessa menina, o adota de um abrigo de animais. Juntos, eles aprendem os valores da amizade e família.', 'Dean Fleischer Camp', 117, '2025-05-22', 0, 6.7, NOW(), NULL, NULL, 'af23a83a-126f-4b59-89b6-1f07f6cb7b28', '/img/lilo.png'),
                                                ('f1a3dbb2-11d3-42c2-b0e1-9ea327f228df', 'Interestelar', 'Um grupo de astronautas viaja através de um buraco de minhoca em busca de um novo lar para a humanidade.', 'Christopher Nolan', 169, '2025-05-15', 12, 8.6, NOW(), NULL, NULL, 'af23a83a-126f-4b59-89b6-1f07f6cb7b28', '/img/interestelar.png'),
                                                ('d2a4a772-0fd0-4a1e-8fa9-4c2ea60c6f62', 'Vingadores: Ultimato', 'Após o estalo de Thanos, os heróis sobreviventes unem forças para desfazer o caos e restaurar a ordem no universo.', 'Anthony e Joe Russo', 181, '2025-05-10', 14, 8.4, NOW(), NULL, NULL, 'af23a83a-126f-4b59-89b6-1f07f6cb7b28', '/img/vingadores.png');

-- Relação Filme-Gênero
INSERT INTO re_filme_genero (id, id_filme, id_genero) VALUES
                                                ('c0560d5d-d735-4c84-b957-cce317fe12b2', '2b0c3fe1-4e29-4fd9-b6f6-cd5c9f88f68a', '364a1dfb-6744-42f6-aeb0-cb15b8ff3b94'),
                                                ('b13a52c2-9256-4b27-b230-1612214a7b1e', '2b0c3fe1-4e29-4fd9-b6f6-cd5c9f88f68a', 'ab243d57-63b6-4f3b-bb5a-08b3f0b58e44');

-- Salas
INSERT INTO sala (id, nome, fileiras, poltronas) VALUES
                                                ('b145d15f-1a3c-4f94-82c6-84b37034545d', 'Sala 1', 'E', 10);


-- Sessões
INSERT INTO sessao (id, id_filme, id_sala, data_inicio, data_fim, preco_ingresso, disponibilidade) VALUES
                                               ('d85a3e2b-7e1e-4c59-a046-441377ba7f89', '2b0c3fe1-4e29-4fd9-b6f6-cd5c9f88f68a', 'b145d15f-1a3c-4f94-82c6-84b37034545d', '2025-05-31 18:00:00', '2025-05-31 20:00:00', 30.00, 1),
                                               ('529d13da-8e26-4962-8b0d-ccf078f25616', '2b0c3fe1-4e29-4fd9-b6f6-cd5c9f88f68a', 'b145d15f-1a3c-4f94-82c6-84b37034545d', '2025-06-01 18:00:00', '2025-06-01 20:00:00', 30.00, 1),
                                               ('1dfc3c98-b941-401f-8c43-d85c3b1156ca', '2b0c3fe1-4e29-4fd9-b6f6-cd5c9f88f68a', 'b145d15f-1a3c-4f94-82c6-84b37034545d', '2025-06-02 18:00:00', '2025-06-02 20:00:00', 30.00, 1),
                                               ('2993d3b5-3968-4d73-8e02-d3fca9a2d292', '2b0c3fe1-4e29-4fd9-b6f6-cd5c9f88f68a', 'b145d15f-1a3c-4f94-82c6-84b37034545d', '2025-06-03 18:00:00', '2025-06-03 20:00:00', 30.00, 1),
                                               ('2e470826-7fe7-4605-b09d-12f41520c072', '2b0c3fe1-4e29-4fd9-b6f6-cd5c9f88f68a', 'b145d15f-1a3c-4f94-82c6-84b37034545d', '2025-06-04 18:00:00', '2025-06-04 20:00:00', 30.00, 1),
                                               ('f92c8f3a-bf08-4b76-a3b4-183a7dcab3c1', 'f1a3dbb2-11d3-42c2-b0e1-9ea327f228df', 'b145d15f-1a3c-4f94-82c6-84b37034545d', '2025-06-01 15:00:00', '2025-06-01 17:50:00', 35.00, 1),
                                               ('a75e1f86-8a31-40c4-a348-d2ce60ff89e7', 'f1a3dbb2-11d3-42c2-b0e1-9ea327f228df', 'b145d15f-1a3c-4f94-82c6-84b37034545d', '2025-06-02 15:00:00', '2025-06-02 17:50:00', 35.00, 1),
                                               ('30a7c3ab-3aef-4cc4-b8e2-8ef2f9d9f0c7', 'f1a3dbb2-11d3-42c2-b0e1-9ea327f228df', 'b145d15f-1a3c-4f94-82c6-84b37034545d', '2025-06-03 15:00:00', '2025-06-03 17:50:00', 35.00, 1),
                                               ('ce8c08b1-f82f-4082-95d1-560d8e11622d', 'f1a3dbb2-11d3-42c2-b0e1-9ea327f228df', 'b145d15f-1a3c-4f94-82c6-84b37034545d', '2025-06-04 15:00:00', '2025-06-04 17:50:00', 35.00, 1),
                                               ('b17a2cf0-063e-4969-b8e5-f89c1b98879e', 'd2a4a772-0fd0-4a1e-8fa9-4c2ea60c6f62', 'b145d15f-1a3c-4f94-82c6-84b37034545d', '2025-06-01 21:00:00', '2025-06-01 23:55:00', 40.00, 1),
                                               ('9b7e4c7c-21df-4de2-b04e-c87d73a4038a', 'd2a4a772-0fd0-4a1e-8fa9-4c2ea60c6f62', 'b145d15f-1a3c-4f94-82c6-84b37034545d', '2025-06-02 21:00:00', '2025-06-02 23:55:00', 40.00, 1),
                                               ('c9473dfb-f4e2-4a92-90d3-45d8e49b44cf', 'd2a4a772-0fd0-4a1e-8fa9-4c2ea60c6f62', 'b145d15f-1a3c-4f94-82c6-84b37034545d', '2025-06-03 21:00:00', '2025-06-03 23:55:00', 40.00, 1);
