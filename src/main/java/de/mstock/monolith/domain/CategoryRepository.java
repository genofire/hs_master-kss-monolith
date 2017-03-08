package de.mstock.monolith.domain;

import java.util.List;

import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.Repository;

public interface CategoryRepository extends Repository<Category, Integer> {

  @Query("select distinct c from Category c join fetch c.i18n i "
      + "where key(i) = ?1 order by c.ordinal")
  List<Category> findAllOrdered(String language);

  @Query("select c from Category c join fetch c.i18n i join fetch c.products "
      + "where key(i) = ?1 and lower(i.prettyUrlFragment) = ?2")
  Category findByPrettyUrlFragment(String language, String prettyUrlFragment);

}
