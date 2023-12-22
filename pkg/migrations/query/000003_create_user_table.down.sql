START TRANSACTION;

ALTER TABLE
  users DROP FOREIGN KEY FK_users_statuses_id;

ALTER TABLE
  users DROP FOREIGN KEY FK_users_roles_id;

ALTER TABLE
  users DROP INDEX IDX_users_hash;

DROP TABLE users;

COMMIT;
