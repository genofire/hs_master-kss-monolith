package de.mstock.monolith.web;

public class CategoryDTO {

  private final String name;
  private final String prettyUrlFragment;

  public CategoryDTO(String name, String prettyUrlFragment) {
    this.name = name;
    this.prettyUrlFragment = prettyUrlFragment;
  }

  public String getName() {
    return name;
  }

  public String getPrettyUrlFragment() {
    return prettyUrlFragment;
  }

}
