package de.mstock.monolith.service;

import java.util.Locale;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import de.mstock.monolith.domain.DataTransferObjectFactory;
import de.mstock.monolith.domain.Product;
import de.mstock.monolith.domain.ProductRepository;
import de.mstock.monolith.domain.Review;
import de.mstock.monolith.domain.ReviewRepository;
import de.mstock.monolith.web.ReviewDTO;
import de.mstock.monolith.web.form.ReviewForm;

@Service
public class ReviewService {

  @Autowired
  private ReviewRepository reviewRepository;

  @Autowired
  private ProductRepository productRepository;

  @Autowired
  private DataTransferObjectFactory dtoFactory;

  /**
   * Stores a review from a posted form.
   * 
   * @param reviewForm Post data
   * @param locale Language context
   * @param prettyUrlFragment Used to get the product context
   * @return DTO
   */
  public ReviewDTO saveReview(ReviewForm reviewForm, Locale locale, String prettyUrlFragment) {
    Product product = productRepository.findByI18nName(locale.getLanguage(), prettyUrlFragment);
    Review review = new Review();
    review.setProductId(product.getId());
    review.setLocaleLanguage(locale.getLanguage());
    review.setFirstName(reviewForm.getFirstName());
    review.setLastName(reviewForm.getLastName());
    review.setRatingStars(reviewForm.getRatingStars());
    review.setText(reviewForm.getText());
    return dtoFactory.createReviewDTO(reviewRepository.save(review));
  }
}
