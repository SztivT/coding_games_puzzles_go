#include <iostream>
#include <vector>
#include <cmath>

using namespace std;

int main() {
    int robbers;
    cin >> robbers;
    cin.ignore();
    int vaults;
    cin >> vaults;
    cin.ignore();
    vector<int> robbing_time;
    for (int i = 0; i < robbers; i++) {
        robbing_time.push_back(0);
    }
    for (int i = 0; i < vaults; i++) {
        int char_length;
        int digit_length;
        int time;
        cin >> char_length >> digit_length;
        cin.ignore();
        time = pow(10, digit_length) * pow(5, char_length - digit_length);
        int fastest =  2147483647;
        int f_index;
        for (int j = 0; j < robbing_time.size(); j++) {
            if (fastest > robbing_time[j]) {
                fastest = robbing_time[j];
                f_index = j;
            }
        }
        robbing_time[f_index] += time;
    }
    int slow = 0;
    for (int i : robbing_time) {
        if (i > slow) {
            slow = i;
        }
    }
    cout << slow << endl;
    return 0;
}