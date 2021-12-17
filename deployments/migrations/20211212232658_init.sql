-- +goose Up
-- +goose StatementBegin
CREATE TABLE category
(
    id        serial primary key,
    name      varchar(250),
    code      varchar(100),
    is_active boolean default true,
    constraint category_code_uniq unique (code)
);

CREATE TABLE product
(
    id            serial primary key,
    name          varchar(250),
    code          varchar(100),
    updated_at    timestamp default now(),
    category_id   int,
    category_code varchar(250),
    is_active     boolean   default true,
    constraint product_category_id_fk foreign key (category_id) references category (id) ON DELETE RESTRICT,
    constraint product_category_code_fk foreign key (category_code) references category (code) ON DELETE RESTRICT
);

CREATE TABLE app_user
(
    id       serial primary key,
    name     varchar(250),
    login    varchar(16),
    password varchar(255)
);

CREATE TABLE filter
(
    code         varchar(100) primary key,
    filter_name  varchar(255),
    field_name   varchar(255),
    filter_value varchar(255)
);
-- filter min with filter_value = min_value?
-- filter max value with filter_value = max value for selected field
-- or add diapason filter: diapason_filter flag, optional min, max_value fields

CREATE TABLE category_filter
(
    category_code varchar(100),
    filter_code   varchar(100),
    constraint category_filter_category_code_fk foreign key (category_code) references category (code) on delete restrict,
    constraint category_filter_filter_code_fk foreign key (filter_code) references filter (code) on delete restrict
);

CREATE TABLE user_filter
(
    user_id int not null,
    constraint user_filter_user_id_fk foreign key (user_id) references app_user (id) on delete restrict
);

CREATE TABLE favorite_product
(
    user_id    int not null,
    product_id int not null,
    constraint favorite_product_user_id_fk foreign key (user_id) references app_user (id) on delete restrict,
    constraint favorite_product_product_id_fk foreign key (product_id) references product (id) on delete restrict
);

CREATE TABLE shop
(
    id   serial primary key,
    name varchar(250),
    url  varchar(250)
);

CREATE TABLE product_price
(
    shop_id int,
    constraint product_price_shop_id_fk foreign key (shop_id) references shop (id) on delete restrict
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table product_price;
drop table shop;
drop table favorite_product;
drop table user_filter;
drop table category_filter;
drop table filter;
drop table app_user;
drop table product;
drop table category;
-- +goose StatementEnd
