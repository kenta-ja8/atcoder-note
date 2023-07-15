#include <algorithm>
#include <iostream>
#include <map>
#include <string>
#include <vector>

using namespace std;

int solve() {
  int N;
  cin >> N;
  int a[2][N];
  for (int i = 0; i < 2; i++) {
    for (int j = 0; j < N; j++) {
      cin >> a[i][j];
    }
  }

  int max_value = 0;
  for (int i = 0; i < N; i++) {
    int val = 0;
    for (int j = 0; j < N; j++) {
      if (j <= i) {
        val = val + a[0][j];
      }
      if (j >= i) {
        val = val + a[1][j];
      }
    }
    if (val > max_value) {
      max_value = val;
    }
  }

  return max_value;
}
int main() {
  cout << solve() << endl;
  return 0;
}
