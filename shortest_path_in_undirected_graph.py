#User function Template for python3
from collections import deque

class Solution:
    def shortestPath(self, edges, n, m, src):
        # code here
        adj = [[] for i in range(n)]
        
        for fr, to in edges:
            adj[fr].append(to)
            adj[to].append(fr)
            
        dis = [1e9] * n
        dis[src] = 0
        q = deque()
        q.append(src)
        
        while len(q) != 0:
            top = q.pop()
            
            for i in adj[top]:
                if dis[i] > dis[top] + 1:
                    dis[i] = dis[top] + 1
                    q.append(i)
        
        for i, v  in enumerate(dis):
            if v >= 1e9:
                dis[i] = -1
                
        return dis
        