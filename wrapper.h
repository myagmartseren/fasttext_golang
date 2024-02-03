#ifndef WRAPPER_H
#define WRAPPER_H

typedef struct fasttext_t fasttext_t;

#ifdef __cplusplus
extern "C" {
#endif

fasttext_t * fasttext_new();
void fasttext_delete(void *self);
static char* fasttext_predict(fasttext_t *self, const char* text);
void fasttext_load(fasttext_t *self, const char* path);

#ifdef __cplusplus
}
#endif

#endif /* WRAPPER_H */