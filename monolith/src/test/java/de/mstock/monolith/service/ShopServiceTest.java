package de.mstock.monolith.service;

import static org.hamcrest.CoreMatchers.equalTo;
import static org.hamcrest.CoreMatchers.is;
import static org.junit.Assert.assertThat;
import static org.mockito.Matchers.any;
import static org.mockito.Matchers.eq;
import static org.mockito.Mockito.RETURNS_DEEP_STUBS;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

import java.util.Arrays;
import java.util.List;
import java.util.Locale;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.runners.MockitoJUnitRunner;

import de.mstock.monolith.domain.Category;
import de.mstock.monolith.domain.CategoryRepository;
import de.mstock.monolith.domain.DataTransferObjectFactory;
import de.mstock.monolith.domain.ProductRepository;
import de.mstock.monolith.web.CategoryDTO;

@RunWith(MockitoJUnitRunner.class)
public class ShopServiceTest {

  @Mock
  private CategoryRepository categoryRepository;

  @Mock
  private ProductRepository productRepository;

  @Mock
  private DataTransferObjectFactory dataTransferObjectFactory;

  @InjectMocks
  private ShopService shopService;

  @Test
  public void shouldGetDataTransferObjectForEveryEntity() {
    Locale locale = new Locale("de");
    Category category = mock(Category.class, RETURNS_DEEP_STUBS);
    List<Category> categoryEntities = Arrays.asList(category, category, category);
    when(categoryRepository.findAllOrdered(eq(locale.getLanguage()))).thenReturn(categoryEntities);
    when(dataTransferObjectFactory.createCategoryDTO(any(), any()))
        .thenReturn(mock(CategoryDTO.class));
    List<CategoryDTO> categories = shopService.getCategories(locale);
    assertThat("Same amount of categories", categories.size(),
        is(equalTo(categoryEntities.size())));
  }

}
