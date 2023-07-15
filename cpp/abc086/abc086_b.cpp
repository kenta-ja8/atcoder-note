#include <cmath>
#include <iostream>
#include <string>
using namespace std;

int main() {
  string a, b;
  cin >> a >> b;

  int num = stoll(a + b);
  long long root = sqrt(num);
  if (root * root == num) {
    cout << "Yes" << endl;
    return 0;
  }
  cout << "No" << endl;
  return 0;
}
