package de.mstock.monolith.web;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;

@Controller
public class HomepageController {

  private static final String TEMPLATE = "homepage";

  @RequestMapping(value = "/", method = RequestMethod.GET)
  public String homepage(Model model) {
    return TEMPLATE;
  }
}
