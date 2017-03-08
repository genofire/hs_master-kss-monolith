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
public class ProductController {

  private static final String TEMPLATE = "product";

  @Autowired
  private ShopService shopService;

  /**
   * Product page
   * 
   * @param prettyUrlFragment Pretty URL fragment
   * @param model Template model
   * @param locale Current locale
   * @return The template's name.
   */
  @RequestMapping(value = "/products/{prettyUrlFragment:[\\w-]+}", method = RequestMethod.GET)
  public String homepage(@PathVariable String prettyUrlFragment, Model model, Locale locale) {
    model.addAttribute("categories", shopService.getCategories(locale));
    model.addAttribute("product", shopService.getProduct(locale, prettyUrlFragment));
    return TEMPLATE;
  }
}
