from typing import (
    List,
)
from lintcode import (
    Interval,
)

"""
Definition of Interval:
class Interval(object):
    def __init__(self, start, end):
        self.start = start
        self.end = end
"""

class Solution:
    """
    @param intervals: an array of meeting time intervals
    @return: the minimum number of conference rooms required
    """
    def min_meeting_rooms(self, intervals: List[Interval]) -> int:
        start = sorted([i.start for i in intervals])
        end = sorted([i.end for i in intervals])
        
        i = 0
        j = 0
        res, count = 0, 0
        while i < len(start) and j < len(end):
            if start[i] < end[j]:
                i += 1
                count += 1
            else:
                j+=1
                count -= 1
            res = max(res, count)
        return res