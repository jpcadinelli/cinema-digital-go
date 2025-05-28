-- Gêneros
INSERT INTO genero (id, nome, created_at) VALUES
                                              ('364a1dfb-6744-42f6-aeb0-cb15b8ff3b94', 'Comédia', NOW()),
                                              ('ab243d57-63b6-4f3b-bb5a-08b3f0b58e44', 'Ficção Científica', NOW()),
                                              ('d8b5d320-57f7-4c9f-b4c7-e03a0b0f2f7f', 'Ação', NOW()),
                                              ('bbb5d321-57f7-4c9f-b4c7-e03a0b0f2f7f', 'Aventura', NOW());

-- Filmes
INSERT INTO filme (id, titulo, sinopse, diretor, duracao, ano_lancamento, classificacao, nota, criado, atualizado, excluido, id_usuario_registro) VALUES
                                                                                                                                                      ('2b0c3fe1-4e29-4fd9-b6f6-cd5c9f88f68a', 'Lilo & Stitch', 'Stitch, um alienígena, chega ao planeta Terra após fugir de sua prisão e tenta se passar por um cachorro para se camuflar. As coisas mudam quando Lilo, uma travessa menina, o adota de um abrigo de animais. Juntos, eles aprendem os valores da amizade e família.', 'Dean Fleischer Camp', 117, '2025-05-22', 0, 6.7, NOW(), NULL, NULL, 'af23a83a-126f-4b59-89b6-1f07f6cb7b28');

-- Relação Filme-Gênero
INSERT INTO re_filme_genero (id, id_filme, id_genero) VALUES
                                                          ('c0560d5d-d735-4c84-b957-cce317fe12b2', '2b0c3fe1-4e29-4fd9-b6f6-cd5c9f88f68a', '364a1dfb-6744-42f6-aeb0-cb15b8ff3b94'),
                                                          ('b13a52c2-9256-4b27-b230-1612214a7b1e', '2b0c3fe1-4e29-4fd9-b6f6-cd5c9f88f68a', 'ab243d57-63b6-4f3b-bb5a-08b3f0b58e44');

-- Salas
INSERT INTO sala (id, nome, fileiras, poltronas) VALUES
                                                     ('b145d15f-1a3c-4f94-82c6-84b37034545d', 'Sala 1', 'A', 10);

-- Sessões
-- INSERT INTO sessao (id, id_filme, id_sala, data_inicio, data_fim, preco_ingresso, disponibilidade) VALUES
--                                                                                                        ('d85a3e2b-7e1e-4c59-a046-441377ba7f89', '2b0c3fe1-4e29-4fd9-b6f6-cd5c9f88f68a', 'b145d15f-1a3c-4f94-82c6-84b37034545d', '2025-03-23 18:00:00', '2025-03-23 20:00:00', 30.00, 1),
--                                                                                                        ('529d13da-8e26-4962-8b0d-ccf078f25616', '5d07d6f9-4fc3-4f88-b8b1-d92888e7282c', 'f8a3b7ff-e767-450b-9d07-b2cc3c66201f', '2025-03-24 15:00:00', '2025-03-24 17:00:00', 25.00, 1),
--                                                                                                        ('1dfc3c98-b941-401f-8c43-d85c3b1156ca', 'f1ec0f29-fab3-4d33-a660-4b74c7f7d94a', 'e5075a0d-f303-4bb0-a64b-22c439b6be93', '2025-03-25 20:00:00', '2025-03-25 22:00:00', 35.00, 1),
--                                                                                                        ('2993d3b5-3968-4d73-8e02-d3fca9a2d292', 'f925ec02-bd0a-4c81-8137-cc45b5766d77', 'b145d15f-1a3c-4f94-82c6-84b37034545d', '2025-03-26 16:00:00', '2025-03-26 18:00:00', 40.00, 1),
--                                                                                                        ('2e470826-7fe7-4605-b09d-12f41520c072', 'f4efb3e0-3b9b-41d3-8c87-bd573f4b0153', 'f8a3b7ff-e767-450b-9d07-b2cc3c66201f', '2025-03-27 17:00:00', '2025-03-27 19:00:00', 28.00, 1);
-- Criar sessões correspondentes ao lillo & stich