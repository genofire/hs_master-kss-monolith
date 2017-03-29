package de.mstock.monolith.service;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Locale;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import de.mstock.monolith.domain.Category;
import de.mstock.monolith.domain.CategoryRepository;
import de.mstock.monolith.domain.DataTransferObjectFactory;
import de.mstock.monolith.domain.Product;
import de.mstock.monolith.domain.ProductRepository;
import de.mstock.monolith.web.CategoryDTO;
import de.mstock.monolith.web.ProductDTO;

@Service
public class ShopService {

  @Autowired
  private CategoryRepository categoryRepository;

  @Autowired
  private ProductRepository productRepository;

  @Autowired
  private DataTransferObjectFactory dtoFactory;

  /**
   * Gets all categories of the current language.
   * 
   * @return A simplified Data Transfer Object.
   */
  public List<CategoryDTO> getCategories(Locale locale) {
    String language = locale.getLanguage();
    List<CategoryDTO> categories = new ArrayList<>();
    for (Category category : categoryRepository.findAllOrdered(language)) {
      categories.add(dtoFactory.createCategoryDTO(category, locale));
    }
    return Collections.unmodifiableList(categories);
  }

  /**
   * Gets all products for a category in the current language.
   * 
   * @return A simplified Data Transfer Object.
   */
  public List<ProductDTO> getProductsForCategory(Locale locale, String prettyUrlFragment) {
    String language = locale.getLanguage();
    Category category = categoryRepository.findByPrettyUrlFragment(language, prettyUrlFragment);
    if (category == null) {
      throw new NotFoundException();
    }
    List<ProductDTO> products =
        dtoFactory.createProductWithoutReviewsDTOs(category.getProducts(), locale);
    return Collections.unmodifiableList(products);
  }

  /**
   * Gets a product in the current language.
   * 
   * @return A simplified Data Transfer Object.
   */
  public ProductDTO getProduct(Locale locale, String prettyUrlFragment) {
    Product product = productRepository.findByI18nName(locale.getLanguage(), prettyUrlFragment);
    if (product == null) {
      throw new NotFoundException();
    }
    return dtoFactory.createProductDTO(product, locale);
  }

}
