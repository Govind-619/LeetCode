const (
    rowLength = 6
    numLetters = 26
)

type Letter int

func NewLetter(ch int) (l Letter) {
    return Letter(ch - 'A')
}

func (l Letter) Distance(o Letter) (distance int) {
    lx := int(l) % rowLength
    ly := int(l) / rowLength
    ox := int(o) % rowLength
    oy := int(o) / rowLength
    if lx > ox {
        distance += lx - ox
    } else {
        distance += ox - lx
    }
    if ly > oy {
        distance += ly - oy
    } else {
        distance += oy - ly
    }
    return distance
}

type Fingers struct {
    left, right Letter;
}

func NewFingers(left, right Letter) (f Fingers) {
    f.left, f.right = left, right
    f.Rectify()
    return f
}

func (f *Fingers) Rectify() {
    if f.left > f.right {
        f.left, f.right = f.right, f.left
    }
}

type Tracker map[Fingers]int

func NewTracker() (t Tracker) {
    return make(map[Fingers]int)
}

func (t *Tracker) Update(f Fingers, newVal int) {
    tVar := (*t)
    origVal, ok := tVar[f]
    if !ok || newVal < origVal{
        tVar[f] = newVal
    }
}

func minimumDistance(word string) int {
    prev := NewTracker()
    cur := NewTracker()

    for left := range numLetters {
        for right := left + 1; right <= numLetters; right++ {
            prev[NewFingers(Letter(left), Letter(right))] = 0
        }
    }

    for _, ch := range word {
        l := NewLetter(int(ch))

        for f, d := range prev {
            if f.left != l {
                newD := d + f.right.Distance(l)
                newF := NewFingers(f.left, l)
                cur.Update(newF, newD)
            }
            if f.right != l {
                newD := d + f.left.Distance(l)
                newF := NewFingers(l, f.right)
                cur.Update(newF, newD)
            }
        }
        prev, cur = cur, NewTracker()
    }

    ans := -1
    for _, d := range prev {
        if ans == -1 || d < ans {
            ans = d
        }
    }
    return ans
}