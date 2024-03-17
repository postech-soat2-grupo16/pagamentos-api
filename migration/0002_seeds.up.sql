DELETE FROM pagamentos;

TRUNCATE pagamentos RESTART IDENTITY;

INSERT INTO pagamentos (pedido_id, amount, status)
VALUES ('65619d06-f3fb-4726-b9fa-be597efa0417', 2, 'RECUSADO'),
       ('d06de888-8dd4-457c-8486-2291d5748d48', 2, 'APROVADO'),
       ('e63f2d9e-4407-414d-8bf8-6a21ebf0a8b6', 1, 'APROVADO'),
       ('7d9b08b4-3d1c-4d2b-b0e3-d5a0f0b1d494', 1, 'APROVADO'),
       ('c8c58ca6-8a92-4ce3-8132-5a711a6e900a', 1, 'APROVADO'),
       ('59003cc2-3391-47d1-a93f-7a5e3dc89c9f', 1, 'APROVADO');