DELETE FROM pagamentos;

TRUNCATE pagamentos RESTART IDENTITY;

INSERT INTO pagamentos (pedido_id, amount, status)
VALUES (3, 2, 'NEGADO'),
       (3, 2, 'APROVADO'),
       (5, 1, 'APROVADO'),
       (6, 1, 'APROVADO'),
       (7, 1, 'APROVADO'),
       (8, 1, 'APROVADO');