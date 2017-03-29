package de.mstock.monolith.web;

import java.util.Locale;

import org.apache.commons.lang3.StringUtils;

public class ReviewDTO {

  private Locale locale;
  private String firstName;
  private String lastName;
  private int ratingStars;
  private String text;

  public String getLanguage() {
    return locale.getLanguage();
  }

  public void setLanguage(String language) {
    this.locale = new Locale(language);
  }

  /**
   * Presentation layer can use this method to access the full name of the reviewer.
   * 
   * @return concatenated name
   */
  public String getDisplayName() {
    if (StringUtils.isNoneBlank(firstName, lastName)) {
      return firstName + " " + lastName.charAt(0) + ".";
    }
    return StringUtils.defaultString(firstName);
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
