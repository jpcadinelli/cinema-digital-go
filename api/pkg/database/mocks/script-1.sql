-- Gêneros
INSERT INTO genero (id, nome, created_at) VALUES
                                              ('d8b5d320-57f7-4c9f-b4c7-e03a0b0f2f7f', 'Ação', NOW()),
                                              ('364a1dfb-6744-42f6-aeb0-cb15b8ff3b94', 'Comédia', NOW()),
                                              ('2078a9da-739e-4fe7-9a0e-f8b5b84b46d0', 'Drama', NOW()),
                                              ('ab243d57-63b6-4f3b-bb5a-08b3f0b58e44', 'Ficção Científica', NOW()),
                                              ('16b6b4de-437e-496b-9b01-06b6edb2a83d', 'Terror', NOW());

-- Filmes
INSERT INTO filme (id, titulo, sinopse, diretor, duracao, ano_lancamento, classificacao, nota, criado, atualizado, excluido, id_usuario_registro) VALUES
                                                                                                                                                      ('2b0c3fe1-4e29-4fd9-b6f6-cd5c9f88f68a', 'Vingadores: Ultimato', 'Os Vingadores devem reunir as forças para enfrentar Thanos, após a destruição do universo.', 'Anthony Russo, Joe Russo', 181, '2019-04-26', 12, 8.4, NOW(), NULL, NULL, 'af23a83a-126f-4b59-89b6-1f07f6cb7b28'),
                                                                                                                                                      ('5d07d6f9-4fc3-4f88-b8b1-d92888e7282c', 'O Rei Leão', 'Simba, um jovem leão, deve superar a tragédia de sua infância e lutar pelo seu lugar como rei.', 'Jon Favreau', 118, '2019-07-19', 6, 8.5, NOW(), NULL, NULL, 'd1b25936-dedb-42d9-bb52-ff6c4f7f4646'),
                                                                                                                                                      ('f1ec0f29-fab3-4d33-a660-4b74c7f7d94a', 'Matrix', 'Um hacker descobre a verdade sobre a realidade em que vive e se junta a um grupo para combater um sistema de inteligência artificial.', 'Lana Wachowski, Lilly Wachowski', 136, '1999-03-31', 14, 8.7, NOW(), NULL, NULL, 'fb2f4f88-c056-487f-8d29-babef3f02a10'),
                                                                                                                                                      ('f925ec02-bd0a-4c81-8137-cc45b5766d77', 'Coringa', 'A transformação de Arthur Fleck em um dos vilões mais icônicos da DC, o Coringa.', 'Todd Phillips', 122, '2019-10-04', 18, 8.4, NOW(), NULL, NULL, '6c88c5a9-2f3f-4b24-bddf-f9983fa7e084'),
                                                                                                                                                      ('f4efb3e0-3b9b-41d3-8c87-bd573f4b0153', 'Parasita', 'Uma família pobre se infiltra na vida de uma família rica, o que gera consequências inesperadas.', 'Bong Joon-ho', 132, '2019-05-30', 16, 8.6, NOW(), NULL, NULL, '7a29a63d-d9ae-4019-bd0d-ea2ff06f4566');

-- Relação Filme-Gênero
INSERT INTO re_filme_genero (id, id_filme, id_genero) VALUES
                                                          ('c0560d5d-d735-4c84-b957-cce317fe12b2', '2b0c3fe1-4e29-4fd9-b6f6-cd5c9f88f68a', 'd8b5d320-57f7-4c9f-b4c7-e03a0b0f2f7f'),
                                                          ('b13a52c2-9256-4b27-b230-1612214a7b1e', '2b0c3fe1-4e29-4fd9-b6f6-cd5c9f88f68a', 'ab243d57-63b6-4f3b-bb5a-08b3f0b58e44'),
                                                          ('af38c3a1-14ed-47b9-b2b9-0901cdd7d91d', '5d07d6f9-4fc3-4f88-b8b1-d92888e7282c', '2078a9da-739e-4fe7-9a0e-f8b5b84b46d0'),
                                                          ('d2d81f36-f978-475b-a33c-8e62b6e975ed', '5d07d6f9-4fc3-4f88-b8b1-d92888e7282c', 'ab243d57-63b6-4f3b-bb5a-08b3f0b58e44'),
                                                          ('94c52ed6-feb5-4d43-aadf-cf73fd8e33f1', 'f1ec0f29-fab3-4d33-a660-4b74c7f7d94a', 'ab243d57-63b6-4f3b-bb5a-08b3f0b58e44'),
                                                          ('b9f658b7-0706-49d5-bd57-b73f7f4c5b94', 'f1ec0f29-fab3-4d33-a660-4b74c7f7d94a', 'd8b5d320-57f7-4c9f-b4c7-e03a0b0f2f7f'),
                                                          ('d5a5f7cc-f4fa-400b-9c3d-b44ab8b1cc7b', 'f925ec02-bd0a-4c81-8137-cc45b5766d77', '2078a9da-739e-4fe7-9a0e-f8b5b84b46d0'),
                                                          ('5332bb9e-2289-4d6e-9519-2464a35c7ef2', 'f925ec02-bd0a-4c81-8137-cc45b5766d77', '16b6b4de-437e-496b-9b01-06b6edb2a83d'),
                                                          ('3642a8da-d045-43b1-88e3-3381f88d3b42', 'f4efb3e0-3b9b-41d3-8c87-bd573f4b0153', '2078a9da-739e-4fe7-9a0e-f8b5b84b46d0'),
                                                          ('d447e1db-5a7d-46a1-b055-577640125fa0', 'f4efb3e0-3b9b-41d3-8c87-bd573f4b0153', '364a1dfb-6744-42f6-aeb0-cb15b8ff3b94');

-- Salas
INSERT INTO sala (id, nome, fileiras, poltronas) VALUES
                                                     ('b145d15f-1a3c-4f94-82c6-84b37034545d', 'Sala 1', 'A', 50),
                                                     ('f8a3b7ff-e767-450b-9d07-b2cc3c66201f', 'Sala 2', 'B', 75),
                                                     ('e5075a0d-f303-4bb0-a64b-22c439b6be93', 'Sala 3', 'C', 100);

-- Sessões
INSERT INTO sessao (id, id_filme, id_sala, data_inicio, data_fim, preco_ingresso, disponibilidade) VALUES
                                                                                                       ('d85a3e2b-7e1e-4c59-a046-441377ba7f89', '2b0c3fe1-4e29-4fd9-b6f6-cd5c9f88f68a', 'b145d15f-1a3c-4f94-82c6-84b37034545d', '2025-03-23 18:00:00', '2025-03-23 20:00:00', 30.00, 1),
                                                                                                       ('529d13da-8e26-4962-8b0d-ccf078f25616', '5d07d6f9-4fc3-4f88-b8b1-d92888e7282c', 'f8a3b7ff-e767-450b-9d07-b2cc3c66201f', '2025-03-24 15:00:00', '2025-03-24 17:00:00', 25.00, 1),
                                                                                                       ('1dfc3c98-b941-401f-8c43-d85c3b1156ca', 'f1ec0f29-fab3-4d33-a660-4b74c7f7d94a', 'e5075a0d-f303-4bb0-a64b-22c439b6be93', '2025-03-25 20:00:00', '2025-03-25 22:00:00', 35.00, 1),
                                                                                                       ('2993d3b5-3968-4d73-8e02-d3fca9a2d292', 'f925ec02-bd0a-4c81-8137-cc45b5766d77', 'b145d15f-1a3c-4f94-82c6-84b37034545d', '2025-03-26 16:00:00', '2025-03-26 18:00:00', 40.00, 1),
                                                                                                       ('2e470826-7fe7-4605-b09d-12f41520c072', 'f4efb3e0-3b9b-41d3-8c87-bd573f4b0153', 'f8a3b7ff-e767-450b-9d07-b2cc3c66201f', '2025-03-27 17:00:00', '2025-03-27 19:00:00', 28.00, 1);
