insert into category (name, code, is_active)
values ('Телефоны', 'phone', true);

insert into product (name, code, updated_at, category_id, category_code, is_active)
values ('Redmi 9C NFC', 'redmi_9c_nfc', NOW(), 1, 'phone', true),
        ('Samsung Galaxy A12', 'samsung_galaxy_a12', NOW(), 1, 'phone', true),
        ('Apple iPhone 11 64Gb Фиолетовый', 'apple_iphone_11_64gb_purple', NOW(), 1, 'phone', true);

