#ifndef WRAPPER_H
#define WRAPPER_H

typedef struct args_t args_t;
typedef struct vector_t vector_t;
typedef struct dense_matrix_t dense_matrix_t;
typedef struct meter_t meter_t;
typedef struct fasttext_t fasttext_t;

#ifdef __cplusplus
extern "C" {
#endif

    fasttext_t * fasttext_new();
    
    void fasttext_delete(void *self);
    const char* fasttext_predict(fasttext_t *self, const char* text);
    void fasttext_load_model(fasttext_t *self, const char* path);

#ifdef __cplusplus
}
#endif

#endif /* WRAPPER_H */