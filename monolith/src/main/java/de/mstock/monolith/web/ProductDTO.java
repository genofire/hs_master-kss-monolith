package de.mstock.monolith.web;

import java.util.List;

import de.mstock.monolith.domain.ProductWeightUnit;

public class ProductDTO {

  private int id;
  private String itemNumber;
  private ProductWeightUnit unit;
  private String name;
  private String prettyUrlFragment;
  private String price;
  private String description;
  private List<ReviewDTO> reviews;

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

  public String getPrice() {
    return price;
  }

  public void setPrice(String price) {
    this.price = price;
  }

  public String getDescription() {
    return description;
  }

  public void setDescription(String description) {
    this.description = description;
  }

  public List<ReviewDTO> getReviews() {
    return reviews;
  }

  public void setReviews(List<ReviewDTO> reviews) {
    this.reviews = reviews;
  }

}
