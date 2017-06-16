package de.mstock.monolith.web;

import static org.hamcrest.CoreMatchers.containsString;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.content;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

import java.util.Locale;

import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.HttpHeaders;
import org.springframework.test.context.junit4.SpringRunner;
import org.springframework.test.context.web.WebAppConfiguration;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.setup.MockMvcBuilders;
import org.springframework.web.context.WebApplicationContext;

@RunWith(SpringRunner.class)
@SpringBootTest
@WebAppConfiguration
public class HomepageControllerTest {

  @Autowired
  private WebApplicationContext webApplicationContext;

  private MockMvc mockMvc;

  @Before
  public void setup() {
    mockMvc = MockMvcBuilders.webAppContextSetup(webApplicationContext).build();
  }

  @Test
  public void shouldPrintEnglishTexts() throws Exception {
    mockMvc.perform(get("/").header(HttpHeaders.ACCEPT_LANGUAGE, "en").locale(new Locale("en")))
        .andExpect(status().isOk()).andExpect(content().string(containsString("Fruits")))
        .andExpect(content().string(containsString("to top")));
  }

  @Test
  public void shouldPrintGermanTexts() throws Exception {
    mockMvc.perform(get("/").header(HttpHeaders.ACCEPT_LANGUAGE, "de").locale(new Locale("de")))
        .andExpect(status().isOk()).andExpect(status().isOk())
        .andExpect(content().string(containsString("Obst")))
        .andExpect(content().string(containsString("nach oben")));
  }

  @Test
  public void shouldPrintEnglishTextsForAustralia() throws Exception {
    mockMvc
        .perform(
            get("/").header(HttpHeaders.ACCEPT_LANGUAGE, "en-AU").locale(new Locale("en", "AU")))
        .andExpect(status().isOk()).andExpect(content().string(containsString("Fruits")))
        .andExpect(content().string(containsString("to top")));
  }

  @Test
  public void shouldPrintGermanTextsForAustria() throws Exception {
    mockMvc
        .perform(
            get("/").header(HttpHeaders.ACCEPT_LANGUAGE, "de-AT").locale(new Locale("de", "AT")))
        .andExpect(status().isOk()).andExpect(status().isOk())
        .andExpect(content().string(containsString("Obst")))
        .andExpect(content().string(containsString("nach oben")));
  }

  @Test
  public void shouldPrintEnglishTextsForChina() throws Exception {
    mockMvc
        .perform(
            get("/").header(HttpHeaders.ACCEPT_LANGUAGE, "zh-CN").locale(new Locale("zh", "CN")))
        .andExpect(status().isOk()).andExpect(content().string(containsString("Fruits")))
        .andExpect(content().string(containsString("to top")));
  }

}
