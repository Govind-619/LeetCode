class Solution:
    def getMinDistance(self, nums: List[int], target: int, start: int) -> int:
        vals = float("inf")
        for i in range(0,len(nums)):
            if nums[i]==target:
                if abs(i-start)<vals:
                    vals = abs(i-start)
        return vals
        