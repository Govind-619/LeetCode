class Solution:
    def maxDistance(self, nums1: List[int], nums2: List[int]) -> int:
        maxd = 0
        i= 0
        j=0
        while i<len(nums1) and j<len(nums2):
            if nums1[i]<=nums2[j]:
                if j-i>maxd:
                    maxd= j-i
                j += 1
            else:
                i += 1
                if i>j:
                    j=i
        return maxd
        