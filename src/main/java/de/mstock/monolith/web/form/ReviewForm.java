package de.mstock.monolith.web.form;

import javax.validation.constraints.Max;
import javax.validation.constraints.Min;
import javax.validation.constraints.NotNull;
import javax.validation.constraints.Size;

public class ReviewForm {

  @Size(max = 80)
  private String firstName;
  @Size(max = 80)
  private String lastName;
  @Min(1)
  @Max(5)
  @NotNull
  private int ratingStars;
  @Size(max = 1000)
  private String text;

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
