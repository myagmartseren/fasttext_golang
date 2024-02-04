#ifndef WRAPPER_HPP
#define WRAPPER_HPP

#include<fasttext/fasttext.h>

class FastText{
private:
    fasttext::FastText* f;
public:
    FastText();
    const char* predict(const char* text);
    void load_model(const char* path);
    // ~Fasttext();
};

class DenseMatrix{
private:
    fasttext::DenseMatrix m;
public:
    DenseMatrix();
};

#endif