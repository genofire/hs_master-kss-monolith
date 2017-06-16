package de.mstock.monolith.domain;

import java.io.Serializable;
import java.util.Objects;

import javax.persistence.Embeddable;

@Embeddable
public class CategoryI18nPk implements Serializable {

  private static final long serialVersionUID = 6985884171403931363L;
  private int categoryId;
  private String localeLanguage;

  @Override
  public int hashCode() {
    return Objects.hash(categoryId, localeLanguage);
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
    CategoryI18nPk rhs = (CategoryI18nPk) obj;
    return Objects.equals(categoryId, rhs.categoryId)
        && Objects.equals(localeLanguage, rhs.localeLanguage);
  }

  public int getCategoryId() {
    return categoryId;
  }

  public void setCategoryId(int categoryId) {
    this.categoryId = categoryId;
  }

  public String getLocaleLanguage() {
    return localeLanguage;
  }

  public void setLocaleLanguage(String localeLanguage) {
    this.localeLanguage = localeLanguage;
  }

}
