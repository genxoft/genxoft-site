-- +migrate Up
create table fcm_settings
(
    token varchar(255) constraint fcm_settings_pk primary key,
    timezone varchar(64) default null,
    ip varchar(64),
    created_at UNSIGNED BIG INT,
    updated_at UNSIGNED BIG INT
);
create index fcm_settings_token_idx on fcm_settings (token);

-- +migrate Down
drop table fcm_settings;
