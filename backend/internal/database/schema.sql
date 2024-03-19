CREATE TABLE IF NOT EXISTS links (
  id          INT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  redirect    text NOT NULL,
  link        text NOT NULL UNIQUE,
  clicked     INT  DEFAULT 0,
  random      BOOLEAN DEFAULT false,
  created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
