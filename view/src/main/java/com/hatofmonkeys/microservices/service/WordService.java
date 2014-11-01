package com.hatofmonkeys.microservices.service;

import java.util.List;
import java.util.UUID;

import com.hatofmonkeys.microservices.model.Word;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.stereotype.Repository;

@Repository
public class WordService {

  @Autowired
  private MongoTemplate mongoTemplate;

  public static final String COLLECTION_NAME = "words";

  public List<Word> listWord() {
    return mongoTemplate.findAll(Word.class, COLLECTION_NAME);
  }
}