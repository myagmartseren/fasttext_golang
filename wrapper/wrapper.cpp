#include "wrapper.hpp"
#include <sstream>
#include <string.h>
#include <stdlib.h>
#include <fasttext/fasttext.h>

FastText::FastText() {
    this->f= new fasttext::FastText();
}

void FastText::load_model(const char* path){
    f->loadModel(path);
}

const char* FastText::predict(const char* text){
    std::stringstream ioss(text);
    std::vector<std::pair<fasttext::real, std::string>> predictions;
    int32_t k=1;
    fasttext::real threshold;
    f->predictLine(ioss, predictions, k, threshold);

    std::stringstream resultStream;
    for (const auto& prediction : predictions) {
        resultStream << prediction.second << " ";
    }

    std::string result = resultStream.str();
    const char* resultStr = strdup(result.c_str());
    return resultStr;
}