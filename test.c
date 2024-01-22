#include <pybind11/pybind11.h>
#include <pybind11/cast.h>
#include <pybind11/stl.h>

#include <gensim/models/fasttext.h>

int main() {
  pybind11::gil_scoped gil;

  gensim::models::FastText model;
  model.load("model.bin");

  // Access model attributes and methods...

  return 0;
}
