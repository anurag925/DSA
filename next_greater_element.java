import java.util.Stack;

class SolutionN {
  // Function to find the next greater element for each element of the array.
  public static long[] nextLargerElement(long[] arr, int n) {
    long[] res = new long[arr.length];
    Stack<Long> s = new Stack<>();

    for (int i = arr.length-1; i >= 0; i--) {
      while(!s.empty() && s.peek() <= arr[i]){
        s.pop();
      }
      if(s.empty()){
        res[i] = -1;
      }else{
        res[i] = s.peek();
      }
      s.push(arr[i]);
    }
    return res;
  }
}