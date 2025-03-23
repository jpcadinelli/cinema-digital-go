-- Gêneros
INSERT INTO genero (id, nome, created_at) VALUES
                                              ('f3fa5b38-e2da-496f-8bb9-facb5f3e5b92', 'Aventura', NOW()),
                                              ('0d5c61b4-73c2-4a74-8f92-b6012f7c8b6b', 'Romance', NOW()),
                                              ('feb1d5c7-cf70-4a2d-8e88-d5313d9f9b29', 'Fantasia', NOW()),
                                              ('bf1b32c8-fb33-47a2-9f3a-b6b1be84d03e', 'Mistério', NOW()),
                                              ('3d9cf22d-09fc-4b9c-a9a6-c6b12b1d149d', 'Documentário', NOW());

-- Filmes
INSERT INTO filme (id, titulo, sinopse, diretor, duracao, ano_lancamento, classificacao, nota, criado, atualizado, excluido, id_usuario_registro) VALUES
                                                                                                                                                      ('b5d53982-f79d-470d-84b3-dc0dbbd17456', 'Homem-Aranha: Sem Volta Para Casa', 'Peter Parker lida com as consequências de sua identidade secreta ser revelada e busca ajuda para corrigir sua situação.', 'Jon Watts', 148, '2021-12-17', 12, 8.7, NOW(), NULL, NULL, '31e6a6b4-2a4d-4bfc-9f74-bde43312f31c'),
                                                                                                                                                      ('edffb9a3-51ae-4774-b4e5-7798d56c0b0a', 'O Senhor dos Anéis: A Sociedade do Anel', 'Frodo Bolseiro inicia sua jornada para destruir o Um Anel e impedir que ele caia nas mãos de Sauron.', 'Peter Jackson', 178, '2001-12-19', 14, 8.8, NOW(), NULL, NULL, '48571b92-b024-4c70-bb1a-e5a3d64bbdf4'),
                                                                                                                                                      ('2b287d95-c6c3-4e0c-91c3-59cc60e6506f', 'Titanic', 'O romance entre Jack e Rose no infame navio Titanic, enquanto ele afunda nas águas geladas do Atlântico.', 'James Cameron', 195, '1997-12-19', 12, 7.8, NOW(), NULL, NULL, 'ba258ee7-d276-4e56-a6c2-53fc28f0b2db'),
                                                                                                                                                      ('af6f302d-e6d4-4df2-8ff9-d1b1b1cf5c4a', 'Os Incríveis', 'Uma família de super-heróis tenta viver uma vida normal, mas é chamada novamente para a ação.', 'Brad Bird', 115, '2004-11-05', 6, 8.0, NOW(), NULL, NULL, 'cf990b1c-98c9-4e68-b98d-82a0a01d2990'),
                                                                                                                                                      ('67247b73-7f51-4d12-a0cd-e4321241b2f9', 'Jurassic Park', 'Uma visita a um parque de dinossauros geneticamente modificados se transforma em um pesadelo quando as criaturas se libertam.', 'Steven Spielberg', 127, '1993-06-11', 12, 8.1, NOW(), NULL, NULL, '76d424df-e7ed-4935-bb72-d6adf597a079');

-- Relação Filme-Gênero
INSERT INTO re_filme_genero (id, id_filme, id_genero) VALUES
                                                          ('1a99da23-f027-4780-b924-6d1ea573f1fe', 'b5d53982-f79d-470d-84b3-dc0dbbd17456', 'd8b5d320-57f7-4c9f-b4c7-e03a0b0f2f7f'),
                                                          ('53c18d4b-bc2e-4d2e-a5de-b09cd46062d9', 'b5d53982-f79d-470d-84b3-dc0dbbd17456', 'ab243d57-63b6-4f3b-bb5a-08b3f0b58e44'),
                                                          ('607d7ab0-8f9f-4633-872d-57eb6b83f648', 'edffb9a3-51ae-4774-b4e5-7798d56c0b0a', 'feb1d5c7-cf70-4a2d-8e88-d5313d9f9b29'),
                                                          ('ef89c622-ada8-4b19-b09c-3fd9fcd84f2e', 'edffb9a3-51ae-4774-b4e5-7798d56c0b0a', 'd8b5d320-57f7-4c9f-b4c7-e03a0b0f2f7f'),
                                                          ('fbd6798d-b2d5-47a7-a37b-1d8d4d9c5fe3', '2b287d95-c6c3-4e0c-91c3-59cc60e6506f', '0d5c61b4-73c2-4a74-8f92-b6012f7c8b6b'),
                                                          ('b00348b5-8f4a-4593-b1ac-3d8d50ef0f3d', '2b287d95-c6c3-4e0c-91c3-59cc60e6506f', '3d9cf22d-09fc-4b9c-a9a6-c6b12b1d149d'),
                                                          ('c91e9c0b-b772-40b5-99c9-b8baf0b21eab', 'af6f302d-e6d4-4df2-8ff9-d1b1b1cf5c4a', 'feb1d5c7-cf70-4a2d-8e88-d5313d9f9b29'),
                                                          ('3a5f3fc7-7e9b-42a4-a722-748973f8db12', 'af6f302d-e6d4-4df2-8ff9-d1b1b1cf5c4a', 'd8b5d320-57f7-4c9f-b4c7-e03a0b0f2f7f'),
                                                          ('524d6637-d9da-4c5f-bbf4-f9062a6b3f2b', '67247b73-7f51-4d12-a0cd-e4321241b2f9', 'd8b5d320-57f7-4c9f-b4c7-e03a0b0f2f7f'),
                                                          ('0e689f8d-8a7b-4664-9387-d7429084d443', '67247b73-7f51-4d12-a0cd-e4321241b2f9', '3d9cf22d-09fc-4b9c-a9a6-c6b12b1d149d');

-- Salas
INSERT INTO sala (id, nome, fileiras, poltronas) VALUES
                                                     ('941e2d7b-2433-44b4-9f2c-6f287c5f6769', 'Sala 4', 'D', 55),
                                                     ('0fcfa022-87be-47b9-b0f7-bf522c29c1f9', 'Sala 5', 'E', 80),
                                                     ('4b9bb615-1eb3-45fa-905f-45e2d7aee8f1', 'Sala 6', 'F', 90);

-- Sessões
INSERT INTO sessao (id, id_filme, id_sala, data_inicio, data_fim, preco_ingresso, disponibilidade) VALUES
                                                                                                       ('8fd31bcd-bf8c-413d-9b52-3d1f10d3d7ad', 'b5d53982-f79d-470d-84b3-dc0dbbd17456', '941e2d7b-2433-44b4-9f2c-6f287c5f6769', '2025-03-28 19:00:00', '2025-03-28 21:00:00', 32.00, 1),
                                                                                                       ('320c759e-0437-4fe2-b3ab-e5e1c83f9de3', 'edffb9a3-51ae-4774-b4e5-7798d56c0b0a', '0fcfa022-87be-47b9-b0f7-bf522c29c1f9', '2025-03-29 14:30:00', '2025-03-29 16:30:00', 28.00, 1),
                                                                                                       ('b2e3cc2a-f44c-4c9a-a6a2-f5ffbc3d9286', '2b287d95-c6c3-4e0c-91c3-59cc60e6506f', '4b9bb615-1eb3-45fa-905f-45e2d7aee8f1', '2025-03-30 18:00:00', '2025-03-30 20:00:00', 30.00, 1),
                                                                                                       ('af5a32d7-0247-4c61-b5a7-0a15b2d1e1a2', 'af6f302d-e6d4-4df2-8ff9-d1b1b1cf5c4a', '941e2d7b-2433-44b4-9f2c-6f287c5f6769', '2025-03-31 19:30:00', '2025-03-31 21:30:00', 33.00, 1),
                                                                                                       ('76d0210f-243d-4d06-bc64-20f6c1a0c84c', '67247b73-7f51-4d12-a0cd-e4321241b2f9', '0fcfa022-87be-47b9-b0f7-bf522c29c1f9', '2025-04-01 15:00:00', '2025-04-01 17:00:00', 27.00, 1);
