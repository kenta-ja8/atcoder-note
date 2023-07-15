#include <iostream>
#include <string>
using namespace std;

int main() {
  string answer;
  int b, c;
  cin >> b >> c;
  if (b * c % 2 == 0) {
    answer = "Even";
  } else {
    answer = "Odd";
  }
  cout << answer << endl;
  return 0;
}
