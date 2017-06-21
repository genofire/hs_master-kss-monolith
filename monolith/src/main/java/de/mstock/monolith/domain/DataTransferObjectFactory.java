package de.mstock.monolith.domain;

import java.text.NumberFormat;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Locale;

import org.springframework.stereotype.Component;

import de.mstock.monolith.web.CategoryDTO;
import de.mstock.monolith.web.ProductDTO;
import de.mstock.monolith.web.ReviewDTO;

@Component
public class DataTransferObjectFactory {


    /**
     * Creates a Data Transfer Object (DTO).
     *
     * @param category database entity
     * @param locale   the requested locale
     * @return DTO
     */
    public CategoryDTO createCategoryDTO(Category category, Locale locale) {
        CategoryI18n i18n = category.getI18n().get(locale.getLanguage());
        CategoryDTO categoryDTO = new CategoryDTO();
        categoryDTO.setName(i18n.getName());
        categoryDTO.setPrettyUrlFragment(i18n.getPrettyUrlFragment());
        return categoryDTO;
    }

    /**
     * Creates a Data Transfer Object (DTO).
     *
     * @param product database entity
     * @param locale  the requested locale
     * @return DTO
     */
    public ProductDTO createProductDTO(Product product, Locale locale) {
        return createProductDTO(product, locale, NumberFormat.getCurrencyInstance(locale));
    }

    private ProductDTO createProductDTO(Product product, Locale locale, NumberFormat numberFormat) {
        ProductDTO productDTO = createProductWithoutReviewsDTO(product, locale, numberFormat);
        ProductI18n i18n = product.getI18n().get(locale.getLanguage());
        productDTO.setReviews(createReviewDTOs(i18n.getReviews()));
        return productDTO;
    }

    /**
     * Creates Data Transfer Objects (DTOs).
     *
     * @param products database entities
     * @param locale   the requested locale
     * @return DTOs
     */
    public List<ProductDTO> createProductDTOs(List<Product> products, Locale locale) {
        List<ProductDTO> productDTOs = new ArrayList<>(products.size());
        NumberFormat numberFormat = NumberFormat.getCurrencyInstance(locale);
        for (Product product : products) {
            productDTOs.add(createProductDTO(product, locale, numberFormat));
        }
        return productDTOs;
    }

    /**
     * Creates Data Transfer Objects (DTOs) without loading their reviews.
     *
     * @param products database entities
     * @param locale   the requested locale
     * @return DTOs
     */
    public List<ProductDTO> createProductWithoutReviewsDTOs(List<Product> products, Locale locale) {
        List<ProductDTO> productDTOs = new ArrayList<>(products.size());
        NumberFormat numberFormat = NumberFormat.getCurrencyInstance(locale);
        for (Product product : products) {
            productDTOs.add(createProductWithoutReviewsDTO(product, locale, numberFormat));
        }
        return productDTOs;
    }

    private ProductDTO createProductWithoutReviewsDTO(Product product, Locale locale,
                                                      NumberFormat numberFormat) {
        ProductI18n i18n = product.getI18n().get(locale.getLanguage());
        String price = numberFormat.format(i18n.getPrice());
        ProductDTO productDTO = new ProductDTO();
        // Addition: productDTO.setID()
        productDTO.setId(product.getId());
        productDTO.setItemNumber(product.getItemNumber());
        productDTO.setUnit(product.getUnit());
        productDTO.setName(i18n.getName());
        productDTO.setPrettyUrlFragment(i18n.getPrettyUrlFragment());
        productDTO.setPrice(price);
        productDTO.setDescription(i18n.getDescription());
        return productDTO;
    }

    /**
     * Creates a Data Transfer Object (DTO).
     *
     * @param review database entity
     * @return DTO
     */
    public ReviewDTO createReviewDTO(Review review) {
        ReviewDTO dto = new ReviewDTO();
        dto.setLanguage(review.getLocaleLanguage());
        dto.setRatingStars(review.getRatingStars());
        dto.setFirstName(review.getFirstName());
        dto.setLastName(review.getLastName());
        dto.setText(review.getText());
        return dto;
    }

    private List<ReviewDTO> createReviewDTOs(List<Review> reviews) {
        List<ReviewDTO> ratingDTOs = new ArrayList<>(reviews.size());
        for (Review review : reviews) {
            ratingDTOs.add(createReviewDTO(review));
        }
        return Collections.unmodifiableList(ratingDTOs);
    }

}
