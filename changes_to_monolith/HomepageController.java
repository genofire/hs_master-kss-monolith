package de.mstock.monolith.web;

import java.util.Locale;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;

import de.mstock.monolith.service.ShopService;

@Controller
public class HomepageController {

    private static final String TEMPLATE = "homepage";

    @Autowired
    private ShopService shopService;
    private final String STOCKADMINFRONTENDTEMPLATE = "admin";



    /**
     * Redirect
     *
     * @param model  Template model
     * @return The constant template name fpr the stock admin frontend.
     */
    @RequestMapping(value = "/admin", method = RequestMethod.GET)
    public String redirect(Model model) {
        return this.STOCKADMINFRONTENDTEMPLATE;
    }
    /**
     * Homepage
     *
     * @param model  Template model
     * @param locale Current locale
     * @return The template's name.
     */
    @RequestMapping(value = "/", method = RequestMethod.GET)
    public String homepage(Model model, Locale locale) {
        model.addAttribute("categories", shopService.getCategories(locale));
        return TEMPLATE;
    }

}