import java.util.Stack;

class Solution {
  public static int[] help_classmate(int arr[], int n) {
    int[] res = new int[arr.length];
    Stack<Integer> s = new Stack<>();

    for (int i = arr.length - 1; i >= 0; i--) {
      while (!s.empty() && s.peek() >= arr[i]) {
        s.pop();
      }
      if (s.empty()) {
        res[i] = -1;
      } else {
        res[i] = s.peek();
      }
      s.push(arr[i]);
    }
    return res;
  }
}
