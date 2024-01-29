
#include <stdint.h>
#include <sys/types.h>

#ifndef FASTTEXT_WRAPPER_H
#define FASTTEXT_WRAPPER_H

// Forward declarations for C-compatible structures
struct Args;
struct Vector;
struct DenseMatrix;
struct Meter;
struct FastText;

extern "C" {

// Functions to create and destroy instances
struct Args* createArgs();
struct Vector* createVector(ssize_t dimension);
struct DenseMatrix* createDenseMatrix();
struct Meter* createMeter(bool initializeLabels);
struct FastText* createFastText();

const char* fasttextPredict(struct FastText* fasttext, const char* text);
void* LoadModel(struct FastText* fasttext, const char* path);

}

#endif