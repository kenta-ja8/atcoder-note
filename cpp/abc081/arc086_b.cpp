#include <algorithm>
#include <iostream>
#include <map>
#include <string>
#include <vector>

using std::vector;

using namespace std;

string solve() {
  int N;
  cin >> N;
  int a[N];
  for (int i = 0; i < N; i++) {
    cin >> a[i];
    // std::cout << A[i] << std::endl;
  }

  int min_pos = 0, max_pos = 0;
  for (int i = 0; i < N; ++i) {
    if (a[i] < a[min_pos])
      min_pos = i;
    if (a[i] > a[max_pos])
      max_pos = i;
  }

  bool is_positive = false;
  std::vector<std::pair<int, int>> operations;
  if (std::abs(a[min_pos]) < std::abs(a[max_pos])) {
    is_positive = true;
    for (int i = 0; i < N; ++i) {
      operations.push_back({max_pos + 1, i + 1});
    }

  } else {
    for (int i = 0; i < N; ++i) {
      operations.push_back({min_pos + 1, i + 1});
    }
  }

  if (is_positive) {
    for (int i = 0; i < N - 1; ++i) {
      operations.push_back({i + 1, i + 2});
    }
  } else {
    for (int i = N - 1; i > 0; --i) {
      operations.push_back({i + 1, i});
    }
  }

  string answer;
  answer += to_string(operations.size()) + "\n";
  for (const auto &p : operations) {
    answer += to_string(p.first) + " " + to_string(p.second) + "\n";
  }
  return answer;
}
int main() {
  cout << solve() << endl;
  return 0;
}
