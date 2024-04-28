CREATE TABLE IF NOT EXISTS allowances (
  id SERIAL PRIMARY KEY,
  personalDeduction DECIMAL NOT NULL,
  kReceipt DECIMAL NOT NULL
);

INSERT INTO allowances (personalDeduction, kReceipt) VALUES (60000.0, 50000.0);