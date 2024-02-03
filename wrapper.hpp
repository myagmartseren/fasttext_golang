#ifndef WRAPPER_HPP
#define WRAPPER_HPP
class Fasttext
{
private:
public:
    Fasttext();
    char* predict(Fasttext *self, const char* text);
    void load(Fasttext *self, const char* path);
    ~Fasttext();
};

#endif