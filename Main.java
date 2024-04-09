import java.util.Arrays;

public class Main {
    public static void main(String[] args) {
        int n = 6;
        int m = 5;
        int[] f = new int[] { 1, 26, 20, 27, 19, 16, 6, 9, 30, 24, 20 };
        int[] b = new int[] { 16, 21, 19, 9, 17, 17, 6, 10, 13, 28, 8 };
        
        int minCost = minimizeCost(f, b, n, m);
        System.out.println("Minimum cost of hiring engineers: " + minCost);
    }
    
    public static int minimizeCost(int[] f, int[] b, int n, int m) {
        Arrays.sort(f);
        Arrays.sort(b);
        
        int minCost = 0;
        
        // Sum up the costs of hiring the lowest n frontend developers
        for (int i = 0; i <= n; i++) {
            minCost += f[i];
        }
        
        // Sum up the costs of hiring the lowest m backend developers
        for (int i = 0; i < m; i++) {
            minCost += b[i];
        }
        
        return minCost;
    }
}
