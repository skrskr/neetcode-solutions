# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseBetween(self, head: Optional[ListNode], left: int, right: int) -> Optional[ListNode]:
        # insert dumy node
        dummy = ListNode(0, head)
        
        # 1- reach node in left
        leftPrev, current = dummy, head
        for i in range(left - 1):
            leftPrev = current 
            current = current.next
            
        # 2- reverse linked list from l to r
        prev = None
        for i in range((right-left) + 1):
            tmpNode = current.next
            current.next = prev
            prev = current
            current = tmpNode
            
        
        # 3- update pointers
        leftPrev.next.next = current
        leftPrev.next = prev
        
        return dummy.next
        