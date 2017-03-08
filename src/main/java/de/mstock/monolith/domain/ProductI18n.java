package de.mstock.monolith.domain;

import java.math.BigDecimal;

import javax.persistence.EmbeddedId;
import javax.persistence.Entity;
import javax.persistence.ManyToOne;
import javax.persistence.MapsId;
import javax.persistence.Table;

@Entity
@Table(name = "products_i18n")
public class ProductI18n {

  @EmbeddedId
  private ProductI18nPk productI18nPk;
  @ManyToOne
  @MapsId("productId")
  private Product product;
  private String name;
  private String prettyUrlFragment;
  private BigDecimal price;
  private String description;

  public ProductI18nPk getProductI18nPk() {
    return productI18nPk;
  }

  public void setProductI18nPk(ProductI18nPk productI18nPk) {
    this.productI18nPk = productI18nPk;
  }

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

  public BigDecimal getPrice() {
    return price;
  }

  public void setPrice(BigDecimal price) {
    this.price = price;
  }

  public String getDescription() {
    return description;
  }

  public void setDescription(String description) {
    this.description = description;
  }
}
