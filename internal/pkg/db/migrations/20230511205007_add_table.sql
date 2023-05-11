-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS urls (
    id          BIGSERIAL   PRIMARY KEY NOT NULL,
    long_url    TEXT        UNIQUE NOT NULL,
    short_url   VARCHAR(10) UNIQUE NOT NULL
);

CREATE INDEX index_short_url
    ON urls (short_url);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS urls;
-- +goose StatementEnd
