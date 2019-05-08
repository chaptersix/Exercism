#include "reverse_string.h"
#include <algorithm>

std::string reverse_string::reverse_string(std::string a) {
    reverse(a.begin(), a.end());
    return a;
}