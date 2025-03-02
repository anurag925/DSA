#User function Template for python3

from typing import List

class Solution:
    def reverse(self,St): 
        #code here
        if len(St) == 0:
            return
        top = St.pop()
        self.reverse(St)
        self.pushback(St, top)

    def pushback(self, st: List, e):
        if len(st) == 0:
            st.append(e)
        
        top = st.pop()
        self.pushback(st, e)
        st.append(top)
