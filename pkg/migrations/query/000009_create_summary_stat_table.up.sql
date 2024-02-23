START TRANSACTION;

CREATE TABLE IF NOT EXISTS summary_stats (
  id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  tenantId BIGINT UNSIGNED NOT NULL,
  unprocessed_order INT,
  completed_order INT,
  order_being_sent INT,
  unfinished_complain INT,
  total_complain INT,
  createdAt TIMESTAMP NOT NULL DEFAULT (CURRENT_DATE),
  updatedAt TIMESTAMP NOT NULL DEFAULT (CURRENT_DATE),
  deletedAt TIMESTAMP,
  CONSTRAINT FK_summary_stats_tenants_id
  FOREIGN KEY (tenantId) REFERENCES tenants(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);

COMMIT;