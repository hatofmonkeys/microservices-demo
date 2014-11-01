package com.hatofmonkeys.microservices.model;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

@Document
public class Word {

  @Id
  private String id;
  private String word;
  private int count;

  public String getId() {
    return id;
  }
  public String getWord() {
    return word;
  }
  public int getCount() {
    return count;
  }
}