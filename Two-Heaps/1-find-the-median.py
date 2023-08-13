import heapq

class MedianFinder:

    def __init__(self):
        self.small, self.large = [], []

    def addNum(self, num: int) -> None:
        # python implment min heap only
        # so we will muilple num * -1 and when retive mulplte by -1
        heapq.heappush(self.small, -1 * num)
        
        # make sure every num in small <= every num in larage
        if (self.small and self.large and (self.small[0] * -1 ) > self.large[0]):
            val = -1 * heapq.heappop(self.small)
            heapq.heappush(self.large, val)
            
        # unevent size ?
        # small heap > larage heap or diff greater than 1
        if len(self.small) > len(self.large) + 1:
            val = -1 * heapq.heappop(self.small)
            heapq.heappush(self.large, val)
        
        
        # large heap  > small heap or diff greater than 1
        if len(self.large) > len(self.small) + 1:
            val = heapq.heappop(self.large)
            heapq.heappush(self.small, (-1 * val))

    def findMedian(self) -> float:
        if len(self.small) > len(self.large):
            return -1 * self.small[0]
        
        if len(self.small) < len(self.large):
            return self.large[0]
        
        return ((self.small[0] * -1) + self.large[0]) / 2


# Your MedianFinder object will be instantiated and called as such:
# obj = MedianFinder()
# obj.addNum(num)
# param_2 = obj.findMedian()