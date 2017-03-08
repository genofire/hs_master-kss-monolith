package de.mstock.monolith.web;

import de.mstock.monolith.domain.ProductWeightUnit;

public class ProductDTO {

  private final String itemNumber;
  private final ProductWeightUnit unit;
  private final String name;
  private final String prettyUrlFragment;
  private final String price;
  private final String description;

  /**
   * A simplified representation of a product.
   * 
   * @param itemNumber A product's unique item number.
   * @param unit The unit of a product that relates to its price.
   * @param name The localized name.
   * @param price The product's price per unit.
   * @param description The description of the product.
   */
  public ProductDTO(String itemNumber, ProductWeightUnit unit, String name,
      String prettyUrlFragment, String price, String description) {
    this.itemNumber = itemNumber;
    this.unit = unit;
    this.name = name;
    this.prettyUrlFragment = prettyUrlFragment;
    this.price = price;
    this.description = description;
  }

  public String getItemNumber() {
    return itemNumber;
  }

  public ProductWeightUnit getUnit() {
    return unit;
  }

  public String getName() {
    return name;
  }

  public String getPrice() {
    return price;
  }

  public String getDescription() {
    return description;
  }

  public String getPrettyUrlFragment() {
    return prettyUrlFragment;
  }

}
