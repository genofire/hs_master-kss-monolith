alter table categories_i18n
  add constraint categories_i18n_category_id_fkey foreign key (category_id)
  references categories
  on delete cascade;

alter table products_i18n
  add constraint products_i18n_product_id_fkey foreign key (product_id)
  references products
  on delete cascade;
