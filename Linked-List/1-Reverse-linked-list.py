# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # 1|next-2 -> 2|next-3 -> 3|next-null
        # 1|next-null <- 2|next-1 <- 3|next-2
        
        if head is None:
            return head

        previous, current, nextNode = None, head, None
        
        while current:
            nextNode = current.next
            current.next = previous
            previous = current
            current = nextNode
        
        return previous