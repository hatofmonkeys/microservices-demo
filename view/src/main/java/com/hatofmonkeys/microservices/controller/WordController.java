package com.hatofmonkeys.microservices.controller;

import com.hatofmonkeys.microservices.model.Word;
import com.hatofmonkeys.microservices.service.WordService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.ModelMap;
import org.springframework.util.StringUtils;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.servlet.View;
import org.springframework.web.servlet.view.RedirectView;

@Controller
public class WordController {

  @Autowired
  private WordService wordService;

  @RequestMapping(value = "/word", method = RequestMethod.GET)
  public String getWordList(ModelMap model) {
    model.addAttribute("wordList", wordService.listWord());
    return "output";
  }
}