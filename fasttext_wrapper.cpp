#include <sstream>
#include <string.h>
#include <stdlib.h>
#include "fasttext_wrapper.h"
#include <fasttext/fasttext.h>

// Args structure
struct Args {
    fasttext::Args args;
};

struct Args* createArgs() {
    return new Args();
}

struct Vector {
    fasttext::Vector vector;
};

struct Vector* createVector(ssize_t dimension) {
    return new Vector{fasttext::Vector(dimension)};
}

struct DenseMatrix {
    fasttext::DenseMatrix matrix;
};

struct DenseMatrix* createDenseMatrix() {
    return new DenseMatrix{fasttext::DenseMatrix()};
}

struct Meter {
    fasttext::Meter meter;
};

struct Meter* createMeter(bool initializeLabels) {
    return new Meter{fasttext::Meter(initializeLabels)};
}

struct FastText {
    fasttext::FastText fasttext;
};

struct FastText* createFastText() {
    return new FastText{fasttext::FastText()};
}

const char* fasttextPredict(struct FastText* fasttext, const char* text) {
    std::stringstream ioss(text);
    std::vector<std::pair<fasttext::real, std::string>> predictions;
    int32_t k;
    fasttext::real threshold;
    fasttext->fasttext.predictLine(ioss, predictions, k, threshold);

    std::stringstream resultStream;
    for (const auto& prediction : predictions) {
        resultStream << prediction.second << " ";
    }

    std::string result = resultStream.str();
    const char* resultStr = strdup(result.c_str());
    return resultStr;
}

void* LoadModel(struct FastText* fasttext, const char* path) {
    fasttext->fasttext.loadModel(path);
}
