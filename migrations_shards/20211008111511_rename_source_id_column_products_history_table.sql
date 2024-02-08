-- +goose Up
-- +goose StatementBegin
delete
from products_history
where true;

delete
from products
where true;

alter table products_history
    rename column source_id to user_id;
alter table products_history
    alter column user_id drop default;

comment on column products_history.product_id is null;

drop sequence products_history_source_id_seq;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table products_history
    rename column user_id to source_id;
-- +goose StatementEnd
