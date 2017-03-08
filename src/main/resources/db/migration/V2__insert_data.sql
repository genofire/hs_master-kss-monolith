insert into categories (id, ordinal) VALUES
  (1, 10),
  (2, 20);

insert into categories_i18n (category_id, locale_language, name, pretty_url_fragment) values
  (1, 'en', 'Fruits', 'fruits'),
  (1, 'de', 'Obst', 'obst'),
  (2, 'en', 'Vegetables', 'vegetables'),
  (2, 'de', 'Gemüse', 'gemuese');

insert into products (id, item_number, unit) values
  (1, 'AS-62653', 1),
  (2, 'KP-77763', 2),
  (3, 'KL-58727', 2),
  (4, 'RP-87973', 3),
  (5, 'LC-52364', 2),
  (6, 'WL-59573', 1);

insert into products_i18n (product_id, locale_language, name, pretty_url_fragment, price, description) values
  (1, 'en', 'Kiwis', 'kiwis', 0.21, 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut.'),
  (1, 'de', 'Kiwis', 'kiwis', 0.19, 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut.'),
  (2, 'en', 'Blueberries', 'blueberries', 2.68, 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut.'),
  (2, 'de', 'Heidelbeeren', 'heidelbeeren', 2.52, 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut.'),
  (3, 'en', 'Cherries', 'cherries', 1.67, 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut.'),
  (3, 'de', 'Kirschen', 'kirschen', 1.57, 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut.'),
  (4, 'en', 'Potatoes', 'potatoes', 1.75, 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut.'),
  (4, 'de', 'Kartoffeln', 'kartoffeln', 1.65, 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut.'),
  (5, 'en', 'Tomatoes', 'tomatoes', 1.18, 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut.'),
  (5, 'de', 'Tomaten', 'tomaten', 1.11, 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut.'),
  (6, 'en', 'Rhubarb', 'rhubarb', 1.52, 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut.'),
  (6, 'de', 'Rhabarber', 'rhabarber', 1.43, 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut.');

insert into category_products (category_id, product_id, ordinal) values
  (1, 1, null),
  (1, 2, null),
  (1, 3, 10),
  (1, 5, null),
  (1, 6, null),
  (2, 4, 10),
  (2, 5, 20),
  (2, 6, null);
