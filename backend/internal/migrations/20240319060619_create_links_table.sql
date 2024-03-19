-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS links (
  id          INT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  redirect    VARCHAR(255) NOT NULL,
  url         VARCHAR(255) NOT NULL UNIQUE,
  clicked     INT  DEFAULT 0,
  random      BOOLEAN DEFAULT false,
  created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS links;
-- +goose StatementEnd
