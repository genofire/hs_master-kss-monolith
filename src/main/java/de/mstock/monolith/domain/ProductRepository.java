package de.mstock.monolith.domain;

import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.Repository;

public interface ProductRepository extends Repository<Product, Integer> {

  @Query("select p from Product p join fetch p.i18n i left join fetch i.reviews r "
      + "where key(i) = ?1 and lower(i.name) = ?2")
  Product findByI18nName(String language, String name);

}
