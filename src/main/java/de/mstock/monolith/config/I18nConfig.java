package de.mstock.monolith.config;

import java.util.Arrays;
import java.util.List;
import java.util.Locale;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.LocaleResolver;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurerAdapter;
import org.springframework.web.servlet.i18n.AcceptHeaderLocaleResolver;

import de.mstock.monolith.web.i18n.AcceptHeaderLookupLocaleResolver;

@Configuration
public class I18nConfig extends WebMvcConfigurerAdapter {

  private static final List<Locale> SUPPORTED_LOCALES =
      Arrays.asList(new Locale("en", "US"), new Locale("de", "DE"));

  /**
   * Creates a Bean, managed by Spring.
   * 
   * @return A configured LocaleResolver
   */
  @Bean
  public LocaleResolver localeResolver() {
    Locale defaultLocale = SUPPORTED_LOCALES.get(0);
    AcceptHeaderLocaleResolver localeResolver = new AcceptHeaderLookupLocaleResolver(defaultLocale);
    localeResolver.setSupportedLocales(SUPPORTED_LOCALES);
    localeResolver.setDefaultLocale(defaultLocale);
    return localeResolver;
  }

}
