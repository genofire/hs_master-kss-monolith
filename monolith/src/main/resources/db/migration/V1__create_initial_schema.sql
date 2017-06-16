create table categories (
  id serial primary key,
  ordinal integer not null
);

create table categories_i18n (
  category_id integer,
  locale_language varchar(2) not null,
  name varchar(80) not null,
  pretty_url_fragment varchar(31) not null,
  primary key (category_id, locale_language)
);

create table products (
  id serial primary key,
  item_number varchar(8) not null,
  unit integer not null
);

create table products_i18n (
  product_id integer,
  locale_language varchar(2) not null,
  name varchar(80) not null,
  pretty_url_fragment varchar(31) not null,
  price numeric(10, 2) not null check (price > 0),
  description varchar(1000),
  primary key (product_id, locale_language)
);

create table category_products (
  category_id integer references categories on delete cascade,
  product_id integer references products on delete cascade,
  ordinal integer,
  primary key (category_id, product_id)
);

create unique index i_product_cascade on category_products(product_id, category_id);
