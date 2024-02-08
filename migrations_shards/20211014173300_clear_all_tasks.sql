-- +goose Up
-- +goose StatementBegin
delete from tasks_history where true;
delete from task_user where true;
delete from tasks where true;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
select 'down SQL query';
-- +goose StatementEnd
