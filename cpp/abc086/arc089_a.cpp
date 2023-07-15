#include <cmath>
#include <iostream>
#include <string>
using namespace std;

int main() {
  int time;
  cin >> time;

  int travels[time][3];
  for (int i = 0; i < time; i++) {
    cin >> travels[i][0] >> travels[i][1] >> travels[i][2];
  }

  for (int i = 0; i < time; i++) {
    int pre_time = i == 0 ? 0 : travels[i - 1][0];
    int crr_time = travels[i][0];
    int pre_x = i == 0 ? 0 : travels[i - 1][1];
    int crr_x = travels[i][1];
    int pre_y = i == 0 ? 0 : travels[i - 1][2];
    int crr_y = travels[i][2];
    if (crr_time - pre_time < abs(crr_x - pre_x) + abs(crr_y - pre_y)) {
      cout << "No" << endl;
      return 0;
    };
    if ((crr_time - pre_time) % 2 !=
        (abs(crr_x - pre_x) + abs(crr_y - pre_y)) % 2) {
      cout << "No" << endl;
      return 0;
    };
  }
  cout << "Yes" << endl;
  return 0;
}
