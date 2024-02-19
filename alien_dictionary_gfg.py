#User function Template for python3
from collections import deque

class Solution:
    def compare(self, a, b, adj):
        la, lb = len(a), len(b)
        for i in range(min(la, lb)):
            if a[i] != b[i]:
                adj[ord(a[i]) - ord('a')].append(ord(b[i]) - ord('a'))
                return

    def sort(self, adj, v):
        indegree = [0] * v
        for values in adj:
            for i in values:
                indegree[i]+=1

        q = deque()
        for i, ind in enumerate(indegree):
            if ind == 0:
                q.append(i)

        res = []
        while len(q) != 0:
            top = q.pop()
            res.append(top)

            for i in adj[top]:
                indegree[i]-=1
                if indegree[i] == 0:
                    q.append(i)

        return res

    def findOrder(self,alien_dict, N, K):
        adj = [[] for i in range(K)]

        for i in range(len(alien_dict) - 1):
            self.compare(alien_dict[i], alien_dict[i+1], adj)

        res = self.sort(adj, K)
        s = ""
        for i in res:
            s += chr(i + ord('a'))
        # print(res, s)
        return s
