#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main() {
    int length;
    int height;
    string text;
    vector<string> ascii_abc;
    vector<int> ascii_text;
    cin >> length;
    cin.ignore();
    cin >> height;
    cin.ignore();
    getline(cin, text);
    for (int i = 0; i < height; i++) {
        string row;
        getline(cin, row);
        ascii_abc.push_back(row);
    }
    for (char i : text) {
        int ascii_value = (unsigned char) i;
        if (ascii_value > 96 && ascii_value < 123) {
            ascii_text.push_back(ascii_value - 97);
            continue;
        }
        if (ascii_value < 65 || ascii_value > 90) {
            ascii_text.push_back(26);
            continue;
        }
        ascii_text.push_back(ascii_value - 65);
    }
    for (int h = 0; h < height; h++) {
        for (int letter : ascii_text) {
            cout << ascii_abc[h].substr(letter * length, length);
        }
        cout << endl;
    }
}