#include <iostream>
#include "math_utils.h"

int main() {
    std::cout << "Math Calculator using Takwin Go!" << std::endl;
    
    int a = 5, b = 3;
    std::cout << "a = " << a << ", b = " << b << std::endl;
    std::cout << "add(a, b) = " << add(a, b) << std::endl;
    std::cout << "multiply(a, b) = " << multiply(a, b) << std::endl;
    std::cout << "power(a, 3) = " << power(a, 3) << std::endl;
    
    return 0;
}