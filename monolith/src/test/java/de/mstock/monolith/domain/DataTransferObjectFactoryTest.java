package de.mstock.monolith.domain;

import static org.hamcrest.CoreMatchers.equalTo;
import static org.hamcrest.CoreMatchers.is;
import static org.hamcrest.Matchers.empty;
import static org.hamcrest.Matchers.not;
import static org.hamcrest.Matchers.nullValue;
import static org.junit.Assert.assertThat;
import static org.mockito.Matchers.any;
import static org.mockito.Matchers.anyString;
import static org.mockito.Mockito.RETURNS_DEEP_STUBS;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.never;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

import java.math.BigDecimal;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import java.util.Locale;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.InjectMocks;
import org.mockito.runners.MockitoJUnitRunner;

import de.mstock.monolith.web.ProductDTO;

@RunWith(MockitoJUnitRunner.class)
public class DataTransferObjectFactoryTest {

  @InjectMocks
  private DataTransferObjectFactory dtoFactory;

  @Test
  public void shouldCreateProductDTOWithoutReviews() {
    Product product = mock(Product.class, RETURNS_DEEP_STUBS);
    ProductI18n productI18n = mock(ProductI18n.class);
    when(product.getI18n().get(any())).thenReturn(productI18n);
    when(productI18n.getPrice()).thenReturn(BigDecimal.valueOf(1.23));
    verify(productI18n, never()).getReviews();
    List<ProductDTO> productDTOs =
        dtoFactory.createProductWithoutReviewsDTOs(Arrays.asList(product), new Locale("de"));
    assertThat("The review list is not present", productDTOs.get(0).getReviews(), is(nullValue()));
  }

  @Test
  public void shouldCreateProductDTOWithReviews() {
    Product product = mock(Product.class, RETURNS_DEEP_STUBS);
    Review review = mock(Review.class);
    when(review.getLocaleLanguage()).thenReturn("de");
    when(product.getI18n().get(any()).getPrice()).thenReturn(BigDecimal.valueOf(1.23));
    when(product.getI18n().get(any()).getReviews()).thenReturn(Arrays.asList(review));
    List<ProductDTO> productDTOs =
        dtoFactory.createProductDTOs(Arrays.asList(product), new Locale("de"));
    assertThat("The review list is filled", productDTOs.get(0).getReviews(), is(not(empty())));
  }

  @Test
  public void shouldCreateProductDTOWithEmptyReviews() {
    Product product = mock(Product.class, RETURNS_DEEP_STUBS);
    when(product.getI18n().get(any()).getPrice()).thenReturn(BigDecimal.valueOf(1.23));
    when(product.getI18n().get(any()).getReviews()).thenReturn(Collections.emptyList());
    List<ProductDTO> productDTOs =
        dtoFactory.createProductDTOs(Arrays.asList(product), new Locale("de"));
    assertThat("The review list is present, but empty", productDTOs.get(0).getReviews(),
        is(empty()));
  }

  @Test
  public void shouldFormatPrice() {
    Locale locale = new Locale("en", "US");
    Product product = mock(Product.class, RETURNS_DEEP_STUBS);
    when(product.getI18n().get(anyString()).getPrice()).thenReturn(BigDecimal.valueOf(1.47));
    ProductDTO productDTO = dtoFactory.createProductDTO(product, locale);
    assertThat("Product has a formatted price", productDTO.getPrice(), is(equalTo("$1.47")));
  }
}
