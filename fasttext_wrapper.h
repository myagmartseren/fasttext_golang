// fasttext_wrapper.h

#ifndef FASTTEXT_WRAPPER_H
#define FASTTEXT_WRAPPER_H

#include <cstdint>
#include <cstddef>

#ifdef __cplusplus
extern "C" {
#endif

// Define your FastText structure
struct FastText
;

// Define your TrainCallback function type
typedef void (*TrainCallback)(float, float, double, double, int64_t);

// Declare your C functions
FastText* fasttext_create_model();
void fasttext_destroy_model(FastText* model);
void fasttext_predict(FastText* model, const char* text, int32_t param);
int32_t fasttext_get_word_id(const FastText* model, const char* word);
int32_t fasttext_get_subword_id(const FastText* model, const char* subword);
int32_t fasttext_get_label_id(const FastText* model, const char* label);
void fasttext_get_word_vector(const FastText* model, float* vec, const char* word);
void fasttext_get_subword_vector(const FastText* model, float* vec, const char* subword);
void fasttext_get_input_vector(const FastText* model, float* vec, int32_t ind);
void fasttext_get_args(const FastText* model, char* args);
int32_t fasttext_get_dimension(const FastText* model);
int32_t fasttext_is_quant(const FastText* model);
void fasttext_save_vectors(const FastText* model, const char* filename);
void fasttext_save_model(const FastText* model, const char* filename);
void fasttext_save_output(const FastText* model, const char* filename);
void fasttext_load_model(FastText* model, const char* filename);
void fasttext_get_sentence_vector(const FastText* model, const char* in, float* vec);
void fasttext_quantize(const FastText* model, const char* qargs, TrainCallback callback);
int64_t fasttext_test(const FastText* model, const char* in, int32_t k, float threshold);
void fasttext_test_with_meter(const FastText* model, const char* in, int32_t k, float threshold, const char* meter);
void fasttext_predict_words(const FastText* model, int32_t k, const int32_t* words, float* predictions, float threshold);
int32_t fasttext_predict_line(const FastText* model, const char* in, int32_t k, float threshold, float* predictions);

#ifdef __cplusplus
}
#endif

#endif  // FASTTEXT_WRAPPER_H
