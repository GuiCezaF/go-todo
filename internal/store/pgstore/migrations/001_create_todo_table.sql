CREATE TABLE IF NOT EXISTS to_dos(
  "id"            uuid                PRIMARY KEY     NOT NULL    DEFAULT gen_random_uuid(),
  "title"         VARCHAR(255)                        NOT NULL,
  "description"   TEXT                                NOT NULL,
  "completed"     BOOLEAN                             NOT NULL    DEFAULT FALSE  
)

---- create above / drop below ----

DROP TABLE IF EXISTS to_dos;
