#User function Template for python3

from typing import List
from collections import deque

class Solution:

    def toposort(self, adj, vis, s, v):
        vis[v] = True
        for i, w in adj[v]:
            if not vis[i]:
                self.toposort(adj, vis, s, i)
        s.append(v)



    def shortestPath(self, n : int, m : int, edges : List[List[int]]) -> List[int]:
        path = [1e9] * n
        vis = [False] * n
        s = []
        adj = [[] for i in range(n)]

        for i in range(m):
            fr = edges[i][0]
            to = edges[i][1]
            wt = edges[i][2]
            adj[fr].append([to, wt])

        for i in range(len(adj)):
            if not vis[i]:
                self.toposort(adj, vis, s, i)
        
        path[0] = 0
        for i, idis in enumerate(s):
            for j, wt in adj[i]:
                path[j] = min(path[j], path[i] + wt)

        for i, v in enumerate(path):
            if v >= 1e9:
                path[i] = -1 
        
        return path
