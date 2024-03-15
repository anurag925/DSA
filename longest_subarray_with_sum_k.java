import java.util.HashMap;

public class longest_subarray_with_sum_k {
  public static int longestSubarrayWithSumK(int[] a, long k) {
    // Write your code here
    // Write your code here
    int length = 0;
    long[] prefix = new long[a.length];
    HashMap<Long, Integer> sumStore = new HashMap<Long, Integer>();
    sumStore.put((long) 0, -1);
    long curr = 0;
    for (int i = 0; i < a.length; i++) {
      curr += a[i];
      prefix[i] = curr;

      if (curr == k && i > length) {
        // System.out.println("A");
        length = i + 1;
      } else if (sumStore.containsKey(curr - k)) {
        // System.out.println("B");
        int gap = i - sumStore.get(curr - k);
        if (gap > length) {
          length = gap;
        }
      }
      if (a[i] != 0) {
        sumStore.put(curr, i);
      }
      // System.out.println(sumStore);
    }

    return length;

  }
}