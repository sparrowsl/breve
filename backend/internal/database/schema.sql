CREATE TABLE IF NOT EXISTS links (
  id          SERIAL NOT NULL PRIMARY KEY,
  redirect    text NOT NULL,
  url         text NOT NULL,
  clicked     INTEGER NOT NULL DEFAULT 0,
  random      BOOLEAN NOT NULL DEFAULT false,
  created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
