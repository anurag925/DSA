import java.util.Arrays;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

/**
 * CoriolisRohitTest
 */
public class CoriolisRohitTest {

  public static int solveRec(int i, int sumTillNow, int[][] sumArray){
    if (i >= sumArray.length){
      return sumTillNow;
    }

    return Math.min(solveRec(i+1, sumTillNow + sumArray[i][1], sumArray), solveRec(i+1, sumTillNow + sumArray[i][2], sumArray));
  }

  public static int solve(int N, int M, int[] F, int[] B) {
    // int result = 0;
    // Set<Integer> taken = new HashSet<>();
    // int[][] sumArray = new int[M + N][3];
    // for (int i = 0; i < M + N; i++) {
    //     sumArray[i] = new int[] { i, F[i], B[i] };
    // }
    // if (N > M){
    //   Arrays.sort(sumArray, (o1, o2) -> o1[1] - o2[1]);
    //   for (int j = 0; j < N; j++) {
    //     taken.add(sumArray[j][0]);
    //     result+=sumArray[j][1];
    //   }
    //   Arrays.sort(sumArray, (o1, o2) -> o1[2] - o2[2]);
    //   int j = 0;
    //   while (j < M+N){
    //     int[] val = sumArray[j];
    //     if (!taken.contains(val[0])){
    //       result+=sumArray[j][2];
    //     }
    //     j++;
    //   }

    // }
    // return result;

    int[][] sumArray = new int[M + N][3];
    for (int i = 0; i < M + N; i++) {
        sumArray[i] = new int[] { i, F[i], B[i] };
    }
    System.out.println(Arrays.deepToString(sumArray));
    return solveRec(0, 0, sumArray);
  }

  public static void main(String[] args) {
    // int n = 6;
    // int m = 5;
    // int[] f = new int[] { 1, 26, 20, 27, 19, 16, 6, 9, 30, 24, 20 };
    // int[] b = new int[] { 16, 21, 19, 9, 17, 17, 6, 10, 13, 28, 8 };
    int n = 1;
    int m = 2;
    int[] f = new int[] { 3, 4, 7 };
    int[] b = new int[] { 1, 5, 6 };
    int res = solve(n, m, f, b);
    System.out.println(res);
  }
}