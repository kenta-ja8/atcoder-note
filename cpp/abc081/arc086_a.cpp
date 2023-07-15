#include <algorithm>
#include <iostream>
#include <map>
#include <string>
#include <vector>

using std::vector;

using namespace std;

bool valueComp(const std::pair<int, int> &a, const std::pair<int, int> &b) {
  return a.second < b.second;
}
int solve() {
  int N, K;
  cin >> N >> K;
  int A[N];
  for (int i = 0; i < N; i++) {
    cin >> A[i];
    // std::cout << A[i] << std::endl;
  }

  std::map<int, int> count_map;
  for (int i = 0; i < N; i++) {
    count_map[A[i]]++;
  }
  int count_map_size = count_map.size();

  std::vector<std::pair<int, int> > pairs(count_map.begin(), count_map.end());
  std::sort(pairs.begin(), pairs.end(), valueComp);

  int answer = 0;
  for (std::vector<std::pair<int, int>>::iterator it = pairs.begin();
       it != pairs.end(); ++it) {
    if (count_map_size <= K) {
      return answer;
    }
    answer = answer + it->second;
    count_map_size--;
  }
  return answer;
}
int main() {
  cout << solve() << endl;
  return 0;
}
