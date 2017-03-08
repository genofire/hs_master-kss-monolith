package de.mstock.monolith.web;

import java.util.Locale;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;

import de.mstock.monolith.service.ShopService;

@Controller
public class CategoryController {

  private static final String TEMPLATE = "category";

  @Autowired
  private ShopService shopService;

  /**
   * Category page
   * 
   * @param prettyUrlFragment Pretty URL fragment
   * @param model Template model
   * @param locale Current locale
   * @return The template's name.
   */
  @RequestMapping(value = "/categories/{prettyUrlFragment:[\\w-]+}", method = RequestMethod.GET)
  public String category(@PathVariable String prettyUrlFragment, Model model, Locale locale) {
    model.addAttribute("categories", shopService.getCategories(locale));
    model.addAttribute("products", shopService.getProductsForCategory(locale, prettyUrlFragment));
    model.addAttribute("prettyUrlFragment", prettyUrlFragment);
    return TEMPLATE;
  }
}
