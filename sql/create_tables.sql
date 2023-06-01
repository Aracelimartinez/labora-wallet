CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  user_name VARCHAR(50) NOT NULL,
  document_number VARCHAR(15) NOT NULL,
  document_type VARCHAR(15) NOT NULL,
  country VARCHAR(30) NOT NULL,
  date_of_birth DATE NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT unique_user_identity UNIQUE (document_number, document_type, country )
);

CREATE TABLE IF NOT EXISTS logs (
  id SERIAL PRIMARY KEY,
  approved BOOL DEFAULT FALSE,
  user_id INT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_user_log FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS wallets (
  id SERIAL PRIMARY KEY,
  account_number INT NOT NULL,
  balance DECIMAL(10, 2) DEFAULT 0.00,
  currency VARCHAR(15),
  log_id INT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_log_wallet FOREIGN KEY (log_id)
    REFERENCES logs (id)
    ON DELETE CASCADE,
  CONSTRAINT unique_account_number UNIQUE (account_number)
);

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    sender_wallet_id INT NOT NULL,
    receiver_wallet_id INT,
    type VARCHAR(20) NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    date DATE NOT NULL,
    approved BOOL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT sender_wallet_id FOREIGN KEY (sender_wallet_id)
        REFERENCES wallets (id)
        ON DELETE CASCADE
);
