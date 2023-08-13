import collections
import heapq
class Solution:
    def findLeastNumOfUniqueInts(self, arr: List[int], k: int) -> int:
        m = collections.defaultdict(int)

        # count number of frquescy for each number
        for n in arr:
            m[n] += 1
            

        # using heap sort usin min heap 
        
        h = []
        for key, val in m.items():
            heapq.heappush(h, (val, key))

        while k > 0:
            item = heapq.heappop(h)
            # number repeated muliple time decrease frequency and push again
            if item[0] != 1:
                heapq.heappush(h, (item[0]-1, item[1]))
            k -= 1
        return len(h)
 