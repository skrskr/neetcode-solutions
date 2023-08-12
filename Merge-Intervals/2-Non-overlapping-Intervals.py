class Solution:
    def eraseOverlapIntervals(self, intervals: List[List[int]]) -> int:
        intervals.sort(key=lambda i: i[0])
        PrevEnd = intervals[0][1]
        count = 0
        for start,end in intervals[1:]:
            
            if start < PrevEnd:
                count+=1
                PrevEnd = min(PrevEnd, end)
            else:
                PrevEnd = end
                
        return count
                