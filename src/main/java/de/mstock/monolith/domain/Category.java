package de.mstock.monolith.domain;

import java.util.List;
import java.util.Map;

import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.JoinTable;
import javax.persistence.ManyToMany;
import javax.persistence.MapKey;
import javax.persistence.OneToMany;
import javax.persistence.Table;

@Entity
@Table(name = "categories")
public class Category {

  @Id
  private int id;
  private int ordinal;
  @OneToMany(mappedBy = "category")
  @MapKey(name = "categoryI18nPk.localeLanguage")
  private Map<String, CategoryI18n> i18n;
  @ManyToMany
  @JoinTable(name = "category_products", inverseJoinColumns = @JoinColumn(name = "product_id"))
  private List<Product> products;

  public int getId() {
    return id;
  }

  public void setId(int id) {
    this.id = id;
  }

  public int getOrdinal() {
    return ordinal;
  }

  public void setOrdinal(int ordinal) {
    this.ordinal = ordinal;
  }

  public Map<String, CategoryI18n> getI18n() {
    return i18n;
  }

  public void setI18n(Map<String, CategoryI18n> i18n) {
    this.i18n = i18n;
  }

  public List<Product> getProducts() {
    return products;
  }

  public void setProducts(List<Product> products) {
    this.products = products;
  }

}
