package de.mstock.monolith.domain;

import java.util.Map;

import javax.persistence.Entity;
import javax.persistence.EnumType;
import javax.persistence.Enumerated;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.MapKey;
import javax.persistence.OneToMany;
import javax.persistence.Table;

@Entity
@Table(name = "products")
public class Product {

  @Id
  @GeneratedValue(strategy = GenerationType.IDENTITY)
  private int id;
  private String itemNumber;
  @Enumerated(EnumType.ORDINAL)
  private ProductWeightUnit unit;
  @OneToMany(mappedBy = "product")
  @MapKey(name = "productI18nPk.localeLanguage")
  private Map<String, ProductI18n> i18n;

  public int getId() {
    return id;
  }

  public void setId(int id) {
    this.id = id;
  }

  public String getItemNumber() {
    return itemNumber;
  }

  public void setItemNumber(String itemNumber) {
    this.itemNumber = itemNumber;
  }

  public ProductWeightUnit getUnit() {
    return unit;
  }

  public void setUnit(ProductWeightUnit unit) {
    this.unit = unit;
  }

  public Map<String, ProductI18n> getI18n() {
    return i18n;
  }

  public void setI18n(Map<String, ProductI18n> i18n) {
    this.i18n = i18n;
  }

}
