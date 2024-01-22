#include "fasttext_wrapper.h"

extern "C" {
    FastTextModel* fasttext_load_model(const char* path) {
        FastTextModel* model = new FastTextModel();
        model->model = new fasttext::FastText();
        model->model->loadModel(path);
        return model;
    }

    void fasttext_predict(FastTextModel* model, const char* text, int k, float* predictions) {
        std::vector<std::pair<fasttext::real, fasttext::entry_type>> results;
        model->model->predict(text, k, results);

        for (int i = 0; i < k; i++) {
            predictions[i] = results[i].first;
        }
    }

    int fasttext_getDimension(FastTextModel* model) {
        return model->model->getDimension();
    }

    void fasttext_getWordVector(FastTextModel* model, const char* word, float* vector) {
        model->model->getWordVector(word, vector);
    }
}
