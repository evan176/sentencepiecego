// sentencepiece_wrapper.h
#ifdef __cplusplus
extern "C" {
#endif
  typedef void* SentencePieceProcessorGo;
  SentencePieceProcessorGo loadSentencePieceProcessor(char *path);
  int Encode(SentencePieceProcessorGo sp, char *text, int *tokenIDs, int maxTokens);
  void Free(SentencePieceProcessorGo sp);
#ifdef __cplusplus
}
#endif
