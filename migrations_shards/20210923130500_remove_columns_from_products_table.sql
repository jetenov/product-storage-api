-- +goose Up
-- +goose StatementBegin
alter table products
    drop column raw_data;
alter table products
    drop column competitor_count;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table products
    add column raw_data jsonb;
alter table products
    add column competitor_count int;
-- +goose StatementEnd
