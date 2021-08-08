// sentencepiece_wrapper.cc
#include <iostream>
#include "sentencepiece_wrapper.h"
#include "sentencepiece_processor.h"

SentencePieceProcessorGo loadSentencePieceProcessor(char *path) {
  sentencepiece::SentencePieceProcessor *sp = new sentencepiece::SentencePieceProcessor();
  auto status = sp->Load(std::string(path));
  if (!status.ok()) {
    return NULL;
  }
  return (void*)sp;
}

int Encode(SentencePieceProcessorGo sp, char *text, int *tokenIDs, int maxTokens) {
  std::vector<int> ids;
  auto status = ((sentencepiece::SentencePieceProcessor*)sp)->Encode(std::string(text), &ids);
  // Return non ok
  if (!status.ok()) {
    return -1;
  }
  int size = int(ids.size());
  // If ids size is greater than token_ids, just return size
  if (size > maxTokens) {
    return size;
  }
  // copy value to tokenIDs one by one
  for (int i=0;i < size; i++) {
    *(tokenIDs+i) = (int)ids[i];
  }
  return size;
}

void Free(SentencePieceProcessorGo sp) {
   sentencepiece::SentencePieceProcessor* ptr = (sentencepiece::SentencePieceProcessor*) sp;
   delete ptr;
}
