#ifndef WRAPPER_HPP
#define WRAPPER_HPP

#include<fasttext/fasttext.h>

class Args{
private:
    fasttext::Args a;
public:
    Args();
};

class Vector{
private:
    fasttext::Vector v;
public:
    Vector();
};

class DenseMatrix{
private:
    fasttext::DenseMatrix m;
public:
    DenseMatrix();
};

class Meter{
private:
    fasttext::Meter m;
public:
    Meter();
};

class FastText{
private:
    fasttext::FastText* f;
public:
    FastText();
    const char* predict(const char* text);
    void load_model(const char* path);
    // ~Fasttext();
};

#endif