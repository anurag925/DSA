#User function Template for python3

from typing import List
from collections import deque

class Solution:
    def eventualSafeNodes(self, V : int, adj : List[List[int]]) -> List[int]:
        adjRev =[[]]*V
        indegree = [0] * V
        visited = [False] * V
        for k,i in enumerate(adj):
            for j in i:
                adjRev[j].append(k)
                indegree[k] += 1

        print(adj, indegree, adjRev)

        q = deque()
        for v, degree in enumerate(indegree):
            if degree == 0:
                q.append(v)

        while len(q) != 0:
            top = q.pop()
            visited[top] = True

            for i in adjRev[top]:
                indegree[i]-=1
                if indegree[i] == 0:
                    q.append(i)

        res = []
        for i, v in enumerate(visited):
            if v == True:
                res.append(i)
        return res

# using dfs for topo sort
class Solution2:
    def dfs(self,adj, vis, rec, i, res):
        vis[i] = True
        rec[i] = True
        for j in adj[i]:
            if ((not vis[j]) and (not self.dfs(adj, vis, rec, j, res))) or rec[j]:
                res[j] = 2
                return False

        res[i] = 1
        rec[i] = False
        return True



    def eventualSafeNodes(self, V : int, adj : List[List[int]]) -> List[int]:
        vis = [False] * V
        rec = [False] * V
        res = [-1] * V

        for i, l in enumerate(adj):
            if not vis[i]:
                self.dfs(adj, vis, rec, i, res)

        s = []

        for i, v in enumerate(res):
            if v == 1:
                s.append(i)
        return s
