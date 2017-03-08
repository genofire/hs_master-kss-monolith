package de.mstock.monolith.domain;

import javax.persistence.EmbeddedId;
import javax.persistence.Entity;
import javax.persistence.ManyToOne;
import javax.persistence.MapsId;
import javax.persistence.Table;

@Entity
@Table(name = "categories_i18n")
public class CategoryI18n {

  @EmbeddedId
  private CategoryI18nPk categoryI18nPk;
  @ManyToOne
  @MapsId("categoryId")
  private Category category;
  private String name;
  private String prettyUrlFragment;

  public String getName() {
    return name;
  }

  public void setName(String name) {
    this.name = name;
  }

  public String getPrettyUrlFragment() {
    return prettyUrlFragment;
  }

  public void setPrettyUrlFragment(String prettyUrlFragment) {
    this.prettyUrlFragment = prettyUrlFragment;
  }

}
