#include <iostream>
using namespace std;

int main(void) {
  int num;
  while(cin >> num, num != 42) {
    cout << num << endl;
  }
  return 0;
}
