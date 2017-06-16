package de.mstock.monolith.domain;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.Table;

@Entity
@Table(name = "reviews")
public class Review {

  @Id
  @GeneratedValue(strategy = GenerationType.IDENTITY)
  private int id;
  private int productId;
  private String localeLanguage;
  private String firstName;
  private String lastName;
  private int ratingStars;
  private String text;

  public int getId() {
    return id;
  }

  public void setId(int id) {
    this.id = id;
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

  public String getFirstName() {
    return firstName;
  }

  public void setFirstName(String firstName) {
    this.firstName = firstName;
  }

  public String getLastName() {
    return lastName;
  }

  public void setLastName(String lastName) {
    this.lastName = lastName;
  }

  public int getRatingStars() {
    return ratingStars;
  }

  public void setRatingStars(int ratingStars) {
    this.ratingStars = ratingStars;
  }

  public String getText() {
    return text;
  }

  public void setText(String text) {
    this.text = text;
  }

}
