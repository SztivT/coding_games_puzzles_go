#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

using namespace std;

int main() {
    int n;
    int closest = 5527;
    cin >> n;
    cin.ignore();
    if (n == 0) {
        cout << 0 << endl;
        return 0;
    }
    for (int i = 0; i < n; i++) {
        int temperature;
        cin >> temperature;
        cin.ignore();
        if (abs(temperature) < abs(closest)) {
            closest = temperature;
            continue;
        }
        if (temperature == abs(closest)) {
            closest = temperature;
            continue;
        }
    }
    cout << closest;
    return 0;
}