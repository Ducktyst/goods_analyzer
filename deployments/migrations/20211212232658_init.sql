-- +goose Up
-- +goose StatementBegin
CREATE TABLE category
(
    id   int primary key,
    name varchar(250)
);

CREATE TABLE product
(
    id          int primary key,
    name        varchar(250),
    code        varchar(250),
    updated_at  timestamp,
    category_id int,
    is_active   boolean,
    constraint product_category_id_fk foreign key (category_id) references category (id) ON DELETE RESTRICT
);

create table user
(
    id       int primary key,
    name     varchar(250),
    login    varchar(16),
    password varchar(255)
);

create table filter
(
    filter_name  varchar(255),
    field_name   varchar(255),
    filter_value varchar(255)
);

CREATE TABLE user_filter
(
    user_id int not null,
    constraint user_filter_user_id_fk foreign key (user_id) references user (id) on delete restrict
);

create table favorite_product
(
    user_id int not null,
    product_id int not null,
    constraint favorite_product_user_id_fk foreign key (user_id) references user (id) on delete restrict,
    constraint favorite_product_product_id_fk foreign key (product_id) references product (id) on delete restrict
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
