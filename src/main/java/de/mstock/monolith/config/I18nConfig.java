package de.mstock.monolith.config;

import java.util.Arrays;
import java.util.List;
import java.util.Locale;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.LocaleResolver;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurerAdapter;
import org.springframework.web.servlet.i18n.AcceptHeaderLocaleResolver;

@Configuration
public class I18nConfig extends WebMvcConfigurerAdapter {

  private static final List<Locale> SUPPORTED_LOCALES =
      Arrays.asList(new Locale("de_DE"), new Locale("en_US"));

  /**
   * Creates a Bean, managed by Spring.
   * 
   * @return A configured LocaleResolver
   */
  @Bean
  public LocaleResolver localeResolver() {
    AcceptHeaderLocaleResolver localeResolver = new AcceptHeaderLocaleResolver();
    localeResolver.setSupportedLocales(SUPPORTED_LOCALES);
    localeResolver.setDefaultLocale(Locale.US);
    return localeResolver;
  }

}
