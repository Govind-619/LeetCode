class Solution(object):
    def getMinDistance(self, nums, target, start):
        vals = float("inf")
        for i in range(0,len(nums)):
            if nums[i]==target:
                if abs(i-start)<vals:
                    vals = abs(i-start)
        return vals
        