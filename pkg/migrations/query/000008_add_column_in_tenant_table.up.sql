START TRANSACTION;

ALTER TABLE
  tenants ADD COLUMN levelId BIGINT UNSIGNED NOT NULL
AFTER totalProduct;

ALTER TABLE
  tenants
ADD CONSTRAINT FK_tenants_level_tenants_id
FOREIGN KEY (levelId) REFERENCES level_tenants(id);

COMMIT;
