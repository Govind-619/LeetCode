class Solution:
    def furthestDistanceFromOrigin(self, moves: str) -> int:
        l=0
        r=0
        n=0
        for v in moves:
            if v=="L":
                l+=1
            elif v=="R":
                r+=1
            else:
                n+=1
        if l>r:
            return l+n-r
        elif r>l:
            return r+n-l
        else:
            return n

        