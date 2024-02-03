// #include "wrapper.hpp"
#include <fasttext/fasttext.h>
#include "wrapper.h"
#include "wrapper.hpp"

struct fasttext_t
{
    Fasttext *obj;
};

fasttext_t * fasttext_new(){
    fasttext_t *ft = (fasttext_t *) malloc(sizeof(fasttext_t));
    if (!ft)
        return NULL;

    ft->obj = new Fasttext();

    return ft;
}

void fasttext_delete(void *self)
{
    assert(self);

    Fasttext *fasttext = ((fasttext_t *) self)->obj;
    delete fasttext;

    free(self);
}

double point_distance(point_t *p, point_t *q)
{
    assert(p);
    assert(q);

    return Point::distance(p->obj, q->obj);
}
