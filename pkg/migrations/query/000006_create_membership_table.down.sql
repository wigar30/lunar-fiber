BEGIN TRANSACTION;

ALTER TABLE
  memberships DROP FOREIGN KEY FK_memberships_roles_id;

ALTER TABLE
  memberships DROP FOREIGN KEY FK_memberships_statuses_id;

ALTER TABLE
  memberships DROP FOREIGN KEY FK_memberships_tenants_id;

ALTER TABLE
  memberships DROP FOREIGN KEY FK_memberships_users_id;

ALTER TABLE
  memberships DROP INDEX IDX_memberships_hash;

DROP TABLE IF EXISTS memberships;

COMMIT TRANSACTION;