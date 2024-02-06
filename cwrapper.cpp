#include <cassert>
#include <cstdlib>
#include <stdio.h>
#include "wrapper.h"
#include "wrapper.hpp"

struct fasttext_t{
    FastText *obj;
};

fasttext_t * fasttext_new(){
    fasttext_t *ft = (fasttext_t *) malloc(sizeof(fasttext_t));
    if (!ft)
        return NULL;

    ft->obj = new FastText();

    return ft;
}

void fasttext_delete(void *self){
    assert(self);

    FastText *f = ((fasttext_t *) self)->obj;
    delete f;

    free(self);
}

void fasttext_load_model(fasttext_t *self, const char* path) {
    try {
        self->obj->load_model(path);
    } catch (const std::exception& ex) {
        errno = EIO;  // Example: Set to input/output error
        perror("fasttext_load_model: ");  // Print error message
        std::cerr << "Exception: " << ex.what() << std::endl;
    }
}

const char* fasttext_predict(fasttext_t *self, const char* text){
    try {
        return self->obj->predict(text);
    } catch (const std::exception& ex) {
        errno = EIO;  // Example: Set to input/output error
        perror("fasttext_predict: ");  // Print error message
        std::cerr << "Exception: " << ex.what() << std::endl;
    }
    return NULL;
}