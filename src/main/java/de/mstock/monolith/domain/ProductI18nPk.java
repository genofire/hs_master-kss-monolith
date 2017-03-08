package de.mstock.monolith.domain;

import java.io.Serializable;
import java.util.Objects;

import javax.persistence.Embeddable;

@Embeddable
public class ProductI18nPk implements Serializable {

  private static final long serialVersionUID = 4916705045560501228L;
  private int productId;
  private String localeLanguage;

  @Override
  public int hashCode() {
    return Objects.hash(productId, localeLanguage);
  }

  @Override
  public boolean equals(Object obj) {
    if (obj == null) {
      return false;
    }
    if (obj == this) {
      return true;
    }
    if (getClass() != obj.getClass()) {
      return false;
    }
    ProductI18nPk rhs = (ProductI18nPk) obj;
    return Objects.equals(productId, rhs.productId)
        && Objects.equals(localeLanguage, rhs.localeLanguage);
  }

  public int getProductId() {
    return productId;
  }

  public void setProductId(int productId) {
    this.productId = productId;
  }

  public String getLocaleLanguage() {
    return localeLanguage;
  }

  public void setLocaleLanguage(String localeLanguage) {
    this.localeLanguage = localeLanguage;
  }

}
