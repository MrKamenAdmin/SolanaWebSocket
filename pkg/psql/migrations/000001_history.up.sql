create table "history"
(
    id           int primary key generated always as identity,
    stake        bigint,
    capture_date timestamp
);

do
$$
    begin
        create index f_history_capture_date on history USING btree (capture_date);
    end
$$;