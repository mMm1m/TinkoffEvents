#include <bits/stdc++.h>
#define FASTIO std::ios::sync_with_stdio(false); std::cin.tie(0); std::cout.tie(0);

using namespace std;

void first() {
  long long n;
  cin >> n;
  bool ans = true;
  for(int i = 0; i < n; ++i){
    long long tmp;
    cin >> tmp;
    ans ^= (tmp & 1);
  }
  if(!ans) cout << "NO";
  else cout << "YES";
}

void second() {
  map<string, int> dayMap;
  vector<string> days = {"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"};
  set<int> positions;
  positions.insert(0);
  positions.insert(29);

  int idx = 1, numOfStrings = 4;
  for (const auto& day : days) {
    dayMap[day] = idx++;
  }

  for (int i = 0; i < numOfStrings; ++i) {
    string line;
    getline(cin, line);
    if (line.empty()) continue;

    stringstream ss(line);
    string day;
    while (ss >> day) {
      positions.insert(i * 7 + dayMap[day]);
    }
  }

  int max = 0, min = 0, maxDays = 0;
  vector<int> posList(positions.begin(), positions.end());
  for (size_t i = 1; i < posList.size(); ++i) {
    if (posList[i] - posList[i - 1] - 1 > maxDays) {
      max = posList[i] - 1;
      min = posList[i - 1] + 1;
      maxDays = posList[i] - posList[i - 1] - 1;
    }
  }
  cout << min << " " << max << endl;
}
long long lcg(long long a, long long e, long long m) {
  return (a * e + 11) % m;
}

vector<int> generateSequence(long long a, long long m, int limit) {
  vector<int> sequence;
  long long seed = 0;
  for (int i = 0; i < limit; ++i) {
    seed = lcg(a, seed, m);
    sequence.push_back((abs(seed % 3 - 1) * 5 + abs(seed % 3) * 2) % 8);
  }
  return sequence;
}

void third(){
  long long n, k, a, m;
  cin >> n >> k >> a >> m;

  vector<int> coins = generateSequence(a, m, 100000);  // Генерируем начальные монеты
  long long totalCandies = 0;
  long long totalCoins = 0;
  long long currentSum = 0;
  long long index = 0;

  while (totalCandies < n) {
    if (index >= coins.size()) {  // Генерируем больше монет при необходимости
      vector<int> moreCoins = generateSequence(a, m, 100000);  // Генерируем еще 100000 монет
      coins.insert(coins.end(), moreCoins.begin(), moreCoins.end());
    }

    int coin = coins[index++];
    currentSum += coin;
    totalCoins++;

    if (currentSum >= 3) {
      long long candies = currentSum / 3;
      if (candies >= k) {
        totalCandies += candies;
        currentSum %= 3;
      }
    }
  }

  cout << totalCoins << endl;
}

void forth() {
  int n, m;
  cin >> n >> m;

  vector<int> starts(m);
  vector<int> ends(m);
  vector<int> centers(m);

  for (int i = 0; i < m; i++) {
    cin >> starts[i] >> ends[i];
    centers[i] = (starts[i] + ends[i]) / 2;
  }
  sort(centers.begin(), centers.end());
  int median = centers[m / 2];
  int minActions = 0;
  for (int i = 0; i < m; i++) {
    minActions += abs(centers[i] - median);
  }
  cout << minActions << endl;
}

vector<int> topologicalSort(const vector<vector<int>>& graph, int n) {
  vector<int> inDegree(n + 1, 0);
  for (int i = 1; i <= n; i++) {
    for (int neighbor : graph[i]) {
      inDegree[neighbor]++;
    }
  }

  queue<int> q;
  for (int i = 1; i <= n; i++) {
    if (inDegree[i] == 0) {
      q.push(i);
    }
  }

  vector<int> result;
  while (!q.empty()) {
    int node = q.front();
    q.pop();
    result.push_back(node);
    for (int neighbor : graph[node]) {
      inDegree[neighbor]--;
      if (inDegree[neighbor] == 0) {
        q.push(neighbor);
      }
    }
  }

  if (result.size() != n) {
    return {};
  }
  return result;
}

void fifth() {
  int n, m;
  cin >> n >> m;

  vector<vector<int>> graph(n + 1);
  for (int i = 0; i < m; i++) {
    int a, b;
    cin >> a >> b;
    if (a != b) {
      graph[a].push_back(b);
    }
  }

  vector<int> result = topologicalSort(graph, n);
  if (result.empty()) {
    cout << "NO" << endl;
  } else {
    cout << "YES" << '\n';
    for (int num : result) {
      cout << num << " ";
    }
    cout << endl;
  }
}

int main() {
  FASTIO
  //first();
  //second();
  // third();
  // forth();
  //fifth();
  return 0;
}
