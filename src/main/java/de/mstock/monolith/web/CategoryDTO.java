package de.mstock.monolith.web;

public class CategoryDTO {

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
