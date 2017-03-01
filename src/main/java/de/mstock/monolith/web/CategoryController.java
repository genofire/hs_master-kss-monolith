package de.mstock.monolith.web;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;

@Controller
public class CategoryController {

  private static final String TEMPLATE = "category";

  @RequestMapping(value = "/categories/{name:[\\w-]+}", method = RequestMethod.GET)
  public String homepage(@PathVariable String name, Model model) {
    model.addAttribute("name", name);
    model.addAttribute("activeCategory", name);
    return TEMPLATE;
  }
}
