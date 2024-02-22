START TRANSACTION;

ALTER TABLE
  tenants DROP FOREIGN KEY FK_tenants_level_tenants_id;

ALTER TABLE
  tenants DROP COLUMN levelId;

COMMIT;
