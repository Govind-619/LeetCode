class Solution:
    def mirrorDistance(self, n: int) -> int:
        return abs(self.reverse(n)-n)

    def reverse(self, x: int)-> int:
        s=0
        while x>0:
            s=s*10 + x%10
            x //= 10
        return s