INSERT INTO "accounts" (
  document_number
) VALUES (1234568900);

INSERT INTO "operation_types" (
   description, multiplier
) VALUES ('COMPRA A VISTA', -1), ('COMPRA PARCELADA', -1), ('SAQUE', -1), ('PAGAMENTO', 1);

INSERT INTO "transactions" (
  account_id, operation_type_id, amount, event_date
) VALUES (1,1, -50, '2020-01-01T10:32:07.7199222'), (1,1, -23.5, '2020-01-01T10:48:12.213585'),
(1,1, -18.7, '2020-01-02T19:01:23.1458543')
