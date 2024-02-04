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

void fasttext_load_model(fasttext_t *self,  const char* path){
    self->obj->load_model(path);
}

const char* fasttext_predict(fasttext_t *self, const char* text){
    return self->obj->predict(text);
}