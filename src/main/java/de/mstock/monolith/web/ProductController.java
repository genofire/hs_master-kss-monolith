package de.mstock.monolith.web;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;

@Controller
public class ProductController {

  private static final String TEMPLATE = "product";

  @RequestMapping(value = "/products/{name:[\\w-]+}", method = RequestMethod.GET)
  public String homepage(@PathVariable String name, Model model) {
    model.addAttribute("name", name);
    return TEMPLATE;
  }
}
