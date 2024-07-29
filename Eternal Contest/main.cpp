#include <bits/stdc++.h>

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

std::pair<int, int> task7(){

}

std::pair<double, double> task8(){

}

int task9(){

}

double task10(){

}

int task11(){

}

long long task12(){

}

int main()
{
  auto t = task6();
  std::cout << t.first << " " << t.second;
}
