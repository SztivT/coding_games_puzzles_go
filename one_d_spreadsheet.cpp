#include <iostream>
#include <vector>

using namespace std;

vector<string> operation;
vector<string> arg1;
vector<string> arg2;
vector<int> evaluated;

void print_sheet() {
    for (int i : evaluated) {
        cout << i << endl;
    }
}

int reference(int cell, bool *success) {
    if (operation[cell] != "C") {
        *success = false;
        return 0;
    }
    *success = true;
    return evaluated[cell];
}

int value(const string &arg, bool *success) {
    if (arg[0] != '$') {
        *success = true;
        return stoi(arg);
    }
    return reference(stoi(arg.substr(1)), success);


}

int main() {
    int n;
    int completed = 0;
    cin >> n;
    cin.ignore();
    for (int i = 0; i < n; i++) {
        evaluated.push_back(0);
        string o;
        string a1;
        string a2;
        cin >> o >> a1 >> a2;
        cin.ignore();
        operation.push_back(o);
        arg1.push_back(a1);
        arg2.push_back(a2);
    }
    while (completed != n) {
        for (int i = 0; i < evaluated.size(); i++) {
            bool isSuccess = false;
            bool num1convert = false;
            bool num2convert = false;
            int num1;
            int num2;
            int val;
            if (operation[i] == "C") {
                continue;
            }
            if (operation[i] == "VALUE") {
                val = value(arg1[i], &isSuccess);
            } else {
                num1 = value(arg1[i], &num1convert);
                num2 = value(arg2[i], &num2convert);
                if (num1convert && num2convert) {
                    isSuccess = true;
                }
                val = (operation[i] == "ADD") * (num1 + num2) +
                        (operation[i] == "SUB") * (num1 - num2) +
                        (operation[i] == "MULT") * (num1 * num2);
            }
            if (!isSuccess) {
                continue;
            }
            evaluated[i] = val;
            operation[i] = "C";
            completed++;
        }
    }
    print_sheet();
    return 0;
}