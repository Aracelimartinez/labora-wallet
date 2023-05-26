CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  user_name VARCHAR(50) NOT NULL,
  document_number VARCHAR(15) NOT NULL,
  document_type VARCHAR(15) NOT NULL,
  country VARCHAR(30) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
);

CREATE TABLE IF NOT EXISTS logs (
  id SERIAL PRIMARY KEY,
  aproved BOOL DEFAULT FALSE,
  user_id INT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_user_log FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS wallets (
  id SERIAL PRIMARY KEY,
  balance DECIMAL(10, 2) DEFAULT 0.00,
  currency VARCHAR(15),
  log_id INT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_log_wallet FOREIGN KEY (log_id)
    REFERENCES logs (id)
    ON DELETE CASCADE
);