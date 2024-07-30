#include <bits/stdc++.h>

const int thousand = 1e3+1;
const int hundred = 100;

int task1(){
  int a,b,c,d; std::cin >> a >> b >> c >> d;
  return d > b ? (a+(d-b)*c) : a;
}

long long task2(){
  long long n, pieces = 0, amount = 1; std::cin >> n;
  while(amount < n){
    ++pieces;
    amount *= 2;
  }
  return pieces;
}

int task3(){
  int n,t; std::cin >> n >> t;
  std::vector<int> v(n);
  for(int i = 0; i < n; ++i) std::cin >> v[i];
  int k; std::cin >> k; k--;
  int ans = 1e9;
  for(int i = 0; i < n; ++i){
    bool ok = true;
    int dist = 0;
    // bottom to up
    for(int j = i; j < n; ++j){
      if(j != i) dist += v[j]-v[j-1];
      if(k == j && dist > k){
        ok = false;
        break;
      }
    }
    // up to bottom
    for(int j = n-2; j >= 0; --j){
      if(i == 0) break;
      dist += v[j+1]-v[j];
      if(k == j && dist > k && k < i) {
        ok = false;
        break;
      }
    }
    std::cout << dist << " ";
    if(ok) {
      ans = std::min(ans, dist);
    }

    ok = true;
    dist = 0;
    // bottom to up
    for(int j = i; j >= 0; --j){
      if(j != i) dist += v[j+1]-v[j];
      if(k == j && dist > k){
        ok = false;
        break;
      }
    }
    // up to bottom
    for(int j = 1; j < n; ++j){
      if(i == n-1) break;
      dist += v[j]-v[j-1];
      if(k == j && dist > k && k > i) {
        ok = false;
        break;
      }
    }
    if(ok) {
      ans = std::min(ans, dist);
    }
  }
  return ans;
}

int task4(){
  int n, k;
  std::vector<int>a(1e3+1);
  std::cin >> n >> k;

  for(int i = 1; i <= n; ++i) {
    std::cin >> a[i];
  }

  std::vector<int> b;
  for(int i = 1; i <= n; ++i){
    int x=1;
    while(a[i]>0){
      int ch = a[i]%10;
      a[i] /= 10;
      b.push_back((9-ch)*x);
      x *= 10;
    }
  }
  std::sort(b.rbegin(), b.rend());
  int ans = 0;
  for(int kt = 0; kt < k;  kt++) ans += b[kt];

  return ans;
}

long long task5(){
  long long l,r, t = 0; std::cin >> l >> r;
  for(int k = 1;k < 19;++k){
    for(int u = 1; u <= 9; ++u){
      if((long long)(u*(((long long)pow(10,k)-1)/9)) <= r && (long long)(u*(((long long)pow(10,k)-1)/9)) >= l){
        ++t;
      }
    }
  }
  return t;
}

std::pair<long , long> task6(){
  int n;
  std::vector<int>ab(1e3+1);
  std::cin >> n;
  for(int i = 1; i <= n; ++i) {
    std::cin >> ab[i];
  }
  int cntC = 0, cntM = 0, c = n/2,m=(n+1)/2,a=0,b=0, ansI, ansJ;
  for(int i = 1; i <= n; ++i) {
    cntC += (ab[i] % 2 == 0);
    cntM += (ab[i] % 2 == 1);
    if (i % 2 == 1 && ab[i] % 2 == 0) a++, ansI = i;
    if (i % 2 == 0 && ab[i] % 2 == 1) b++, ansJ = i;
  }
    if(c == cntC && m == cntM && a == 1 && b == 1){
      return std::make_pair(ansI, ansJ);
    }
    return std::make_pair(-1,-1);
}

bool isNotCycle(std::vector<int>& b, int n){
  std::vector<bool> used(n+1, false);
  int now = 1;
  used[now] = true;
  while(!used[b[now]]){
    now = b[now];
    used[now] = true;
  }
  for(int i = 1; i <= n; ++i)
    if(!used[i]) return false;
  return true;
}

std::pair<int, int> task7(){
  int n; std::cin >> n;
  std::vector<int> a(thousand), b(thousand);
  for(int i = 1; i <= n; ++i){
    std::cin >> a[i];
  }
  std::vector<int> cnt(n+1);
  for(int i = 1; i <= n; ++i) {
    cnt[a[i]]++;
    if(cnt[a[i]] > 2) return std::make_pair(-1,-1);
  }
  int zero = 0, two = 0;
  for(int i = 1; i <= n; ++i){
    if(cnt[i] == 0) zero = i;
    else if(cnt[i] == 2) two = i;
  }
  if(zero == 0) return std::make_pair(-1,-1);
  for(int i = 1; i <= n; ++i){
    if(a[i] == two){
      for(int j = 1; j <= n; ++j) b[j] = a[j];
      b[i] = zero;
      if(isNotCycle(b,n)){
        return std::make_pair(i, zero);
      }
    }
  }
  return std::make_pair(-1,-1);
}

std::pair<double, double> task8(){

}

int task9(){
  int n;
  std::cin >> n;
  std::vector<int> a(n+1);
  for (int i = 1; i <= n; ++i) {
    std::cin >> a[i];
  }
  const int INF = 2e9;
  std::vector<std::vector<int>> dp(n+1, std::vector<int>(n+1, INF));

  dp[0][0] = 0;
  for (int i = 1; i <= n; ++i) {
    for (int j = 0; j <= i; ++j) {
      dp[i][j] = std::min(dp[i][j], dp[i-1][j] + a[i]);
      if (a[i] > 100 && j + 1 <= i) {
        dp[i][j+1] = std::min(dp[i][j+1], dp[i-1][j] + a[i]);
      }
      if (j >= 1) {
        dp[i][j-1] = std::min(dp[i][j-1], dp[i-1][j]);
      }
    }
  }

  int ans = INF;
  for (int i = 0; i <= n; ++i) {
    ans = std::min(ans, dp[n][i]);
  }
  std::cout << ans << std::endl;
  return ans;
}

double task10(){

}

long long pow(long long a, long long n, long long p){
  long long res = 1;
  while(n > 0){
    if(n%2) res = (res*a)%p;
    a = (a*a)%p;
    n /=2;
  }
  return res;
}

long long task11(){
  long long l,r,p; std::cin >> l >> r >> p;
  long long ans = 0;
  for(long long i = l; i <= r; ++i){
    ans = (ans + pow(i, p-2, p))%p;
  }
  return ans;
}

long long task12(){

}

int main()
{
  std::cout << task11();
}
