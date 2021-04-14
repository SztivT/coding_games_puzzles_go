#include <iostream>

typedef unsigned long long ull;

using namespace std;

ull convert_string_to_dec(const string& binary) {
    ull letter_length = 0;
    for (char i : binary) {
        if (i == '1') {
            letter_length += 3;
            continue;
        }
        letter_length += 4;
    }
    return letter_length;
}

void convert_dec_bin(ull number, string *binary) {
    if (number / 2 != 0) {
        convert_dec_bin(number / 2, binary);
    }
    char binary_digit[2];
    sprintf(binary_digit, "%llu", number % 2);
    *binary += binary_digit[0];
}

ull word_length_converter(ull number) {
    string binary_string;

    convert_dec_bin(number, &binary_string);
    return convert_string_to_dec(binary_string);
}

int main() {
    ull start;
    ull temp;
    ull steps;
    cin >> start >> steps;
    cin.ignore();
    for (int i = 0; i < steps; i++) {
        temp = word_length_converter(start);
        if (temp == start) {
            break;
        }
        start = temp;
    }
    printf("%llu", start);
}
