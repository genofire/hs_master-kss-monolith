package de.mstock.monolith.web.i18n;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Locale;
import java.util.Locale.LanguageRange;

import javax.servlet.http.HttpServletRequest;

import org.springframework.web.servlet.i18n.AcceptHeaderLocaleResolver;

public class AcceptHeaderLookupLocaleResolver extends AcceptHeaderLocaleResolver {

  private final Locale fallback;
  private final ArrayList<LanguageRange> ranges = new ArrayList<>();

  public AcceptHeaderLookupLocaleResolver(Locale fallback) {
    this.fallback = fallback;
  }

  @Override
  public void setSupportedLocales(List<Locale> locales) {
    super.setSupportedLocales(locales);
    ranges.clear();
    for (Locale supportedLocale : getSupportedLocales()) {
      ranges.add(new LanguageRange(supportedLocale.getLanguage() + "-*"));
    }
  }

  @Override
  public Locale resolveLocale(HttpServletRequest request) {
    Locale resolvedLocale = super.resolveLocale(request);
    if (getSupportedLocales().contains(resolvedLocale)) {
      return resolvedLocale;
    }
    Locale lookup = Locale.lookup(ranges, Arrays.asList(resolvedLocale));
    if (lookup != null) {
      return lookup;
    }
    return fallback;
  }

}
