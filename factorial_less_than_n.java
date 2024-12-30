import java.util.ArrayList;

class Solution {
  static ArrayList<Long> factorialNumbers(long n) {
    ArrayList<Long> ans = new ArrayList<>();
    Long num = (long)1;
    int i = 1;
    while (num <= n) {
      ans.add(num);
      i++;
      num *= i;
    }
    return ans;
  }
}