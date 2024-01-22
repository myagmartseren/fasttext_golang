// fasttext_wrapper.h

#ifndef FASTTEXT_WRAPPER_H
#define FASTTEXT_WRAPPER_H

#include <fasttext/fasttext.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    fasttext::FastText* model;
} FastTextModel;

FastTextModel* fasttext_load_model(const char*);
void fasttext_predict(FastTextModel*, const char*, int, float*);
int fasttext_getDimension(FastTextModel*);
void fasttext_getWordVector(FastTextModel*, const char*, float*);

#ifdef __cplusplus
}
#endif

#endif // FASTTEXT_WRAPPER_H
