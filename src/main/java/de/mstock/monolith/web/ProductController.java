package de.mstock.monolith.web;

import java.util.Locale;
import java.util.List;

import javax.validation.Valid;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.validation.BindingResult;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import de.mstock.monolith.service.ReviewService;
import de.mstock.monolith.service.ShopService;
import de.mstock.monolith.web.form.ReviewForm;

@Controller
public class ProductController {

  private static final String TEMPLATE = "product";

  @Autowired
  private ShopService shopService;

  @Autowired
  private ReviewService reviewService;

  /**
   * Product page
   *
   * @param prettyUrlFragment Pretty URL fragment
   * @param model Template model
   * @param locale Current locale
   * @return The template's name.
   */
  @RequestMapping(value = "/products/{prettyUrlFragment:[\\w-]+}", method = RequestMethod.GET)
  public String product(@PathVariable String prettyUrlFragment, Model model, Locale locale) {
    model.addAttribute("categories", shopService.getCategories(locale));
    model.addAttribute("product", shopService.getProduct(locale, prettyUrlFragment));
    return TEMPLATE;
  }

  /**
   * Post a review
   *
   * @param reviewForm Form data
   * @param bindingResult Form binding result after validation
   * @param prettyUrlFragment Product context
   * @param model Template model
   * @param locale Language context
   * @return The template's name.
   */
  @RequestMapping(value = "/products/{prettyUrlFragment:[\\w-]+}", method = RequestMethod.POST)
  public String post(@Valid ReviewForm reviewForm, BindingResult bindingResult,
      @PathVariable String prettyUrlFragment, Model model, Locale locale) {
    if (bindingResult.hasErrors()) {
      model.addAttribute("success", false);
    } else {
      model.addAttribute("success", true);
      model.addAttribute("reviewPost",
          reviewService.saveReview(reviewForm, locale, prettyUrlFragment));
    }
    model.addAttribute("categories", shopService.getCategories(locale));
    model.addAttribute("product", shopService.getProduct(locale, prettyUrlFragment));
    return TEMPLATE;
  }

  @RequestMapping(value = "/products/{id}.json", method = RequestMethod.GET)
  @ResponseBody
  public ProductDTO productJson(@PathVariable Integer id, Locale locale) {
    return shopService.getProduct(locale, id);
  }

  @RequestMapping(value = "/products.json", method = RequestMethod.GET)
  @ResponseBody
  public List<ProductDTO> allProductJson(Locale locale) {
    return shopService.getAllProducts(locale);
  }

}
