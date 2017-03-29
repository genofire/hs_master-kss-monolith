create table reviews (
  id serial primary key,
  product_id integer,
  locale_language varchar(2) not null,
  first_name varchar(80),
  last_name varchar(80),
  rating_stars integer not null check (rating_stars between 1 and 5),
  text varchar(1000),
  foreign key (product_id, locale_language) references products_i18n(product_id, locale_language)
);

insert into reviews (id, product_id, locale_language, first_name, last_name, rating_stars, text) values
  (1, 1, 'en', 'John', 'Doe', 5, 'Absolutely perfect!'),
  (2, 1, 'de', 'Max', 'Mustermann', 3, null),
  (3, 1, 'de', 'Erika', 'Mustermann', 4, 'Ich liebe dieses Produkt! Leider ist ab und zu auch eine matschige Kiwi dabei.'),
  (4, 4, 'de', 'Otto', 'Normalverbraucher', 4, 'Genau das was ich gesucht habe!'),
  (5, 5, 'de', null, null, 1, 'Schmeckt nicht!'),
  (6, 5, 'en', 'John', null, 2, 'What''s the country of origin of these tomatoes?');
