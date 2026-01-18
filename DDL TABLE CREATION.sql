CREATE TABLE service_requests (
  id CHAR(36) NOT NULL,                 -- UUID
  request_code VARCHAR(50) NULL,

  request_type VARCHAR(100) NOT NULL,

  status VARCHAR(30) NOT NULL DEFAULT 'DRAFT',
  reported_from VARCHAR(50) NULL,

  owner VARCHAR(100) NULL,

  company VARCHAR(100) NULL,
  organization VARCHAR(100) NULL,
  department VARCHAR(100) NULL,

  product_name VARCHAR(150) NULL,
  product_item VARCHAR(150) NULL,
  product_category VARCHAR(150) NULL,
  product_parent_category VARCHAR(150) NULL,

  symptom VARCHAR(150) NULL,             -- Incident
  request_item VARCHAR(150) NULL,        -- Request Fulfillment

  impact VARCHAR(30) NULL,
  urgency VARCHAR(30) NULL,
  priority VARCHAR(30) NULL,

  solver_group VARCHAR(150) NULL,
  solver VARCHAR(150) NULL,

  coordinator_group VARCHAR(150) NULL,
  coordinator VARCHAR(150) NULL,

  sla_type VARCHAR(50) NULL,

  summary TEXT NULL,
  note TEXT NULL,

  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,

  PRIMARY KEY (id),
  UNIQUE KEY uq_request_code (request_code),
  KEY idx_request_type (request_type),
  KEY idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE service_request_attachments (
  id CHAR(36) NOT NULL,
  service_request_id CHAR(36) NOT NULL,

  file_name VARCHAR(255) NOT NULL,
  file_url VARCHAR(500) NOT NULL,
  mime_type VARCHAR(100),
  file_size BIGINT,

  uploaded_at DATETIME NOT NULL,

  PRIMARY KEY (id),
  KEY idx_service_request_id (service_request_id),
  CONSTRAINT fk_attachment_request
    FOREIGN KEY (service_request_id)
    REFERENCES service_requests(id)
    ON DELETE CASCADE
);


CREATE TABLE users (
  id CHAR(36) NOT NULL,
  username VARCHAR(100) NULL,
  password VARCHAR(255) NOT NULL,
  is_active TINYINT(1) NOT NULL DEFAULT 1,

  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  PRIMARY KEY (id),
  UNIQUE KEY uq_users_username (username)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE user_sessions (
   id CHAR(36) NOT NULL,
  user_id CHAR(36) NOT NULL,

  -- random id sesi (boleh dipakai sebagai "sid" claim di JWT access token)
  session_id CHAR(64) NOT NULL,

  status VARCHAR(100) NOT NULL,

  -- refresh token: simpan HASH saja (misal SHA-256 atau bcrypt/argon2)
  refresh_token_hash VARCHAR(255) NOT NULL,
  refresh_expires_at DATETIME NOT NULL,

  -- audit ringan
  ip_address VARCHAR(45) NULL,
  user_agent VARCHAR(255) NULL,

  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  last_seen_at DATETIME NULL,
  revoked_at DATETIME NULL,
  expires_at DATETIME NULL,

  PRIMARY KEY (id),


  UNIQUE KEY uq_user_sessions_user_id (user_id),

  UNIQUE KEY uq_user_sessions_session_id (session_id),

  KEY idx_user_sessions_status (status),
  KEY idx_user_sessions_refresh_expires (refresh_expires_at),

  CONSTRAINT fk_user_sessions_user
    FOREIGN KEY (user_id) REFERENCES users(id)
    ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE r_user_service_request (
  id CHAR(36) NOT NULL,
  user_id CHAR(36) NOT NULL,
  service_request_id CHAR(36) NOT NULL,
  is_active boolean default true,

  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME NULL,

  PRIMARY KEY (id),
   CONSTRAINT fk_usr_user
    FOREIGN KEY (user_id) REFERENCES users(id),

  CONSTRAINT fk_usr_service_request
    FOREIGN KEY (service_request_id) REFERENCES service_requests(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
