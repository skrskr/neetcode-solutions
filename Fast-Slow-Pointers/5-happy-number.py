class Solution:
    def isHappy(self, n: int) -> bool:
        ## Not effcient solution since use Memory O(N)
        # visited = set()
        
        # while n not in visited:
        #     visited.add(n)
            
        #     n = self.squared(n)
            
        #     if n == 1:
        #         return True
            
        # return False

        ## Effecient solution with memory O(1)
        #20 -> 4 -> 16 -> 37 -> 58 -> 89 -> 145 > 42 -> 20
        
        slow = self.squared(n)
        fast = self.squared(self.squared(n))
        
        while slow != fast and fast != 1:
            slow = self.squared(slow)
            fast = self.squared(self.squared(fast))
        
        return fast == 1
        
    def squared(self, n):
        output = 0
        
        while n:
            digit = n % 10
            output += (digit ** 2)
            n = n // 10
    
        return output