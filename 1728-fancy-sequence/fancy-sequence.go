// operation, either addition or multiplication
type op struct {
    plus, times int
}

// number in the sequence, updated by the first cnt operations
type num struct {
    value, cnt int
}

// the sequence and the operations
type Fancy struct {
    seq []num
    ops []op
}

const mod = 1000000007

func Constructor() Fancy {
    return Fancy{[]num{}, []op{}}
}

func (this *Fancy) Append(val int)  {
    this.seq = append(this.seq, num{val,len(this.ops)})
}

func (this *Fancy) AddAll(inc int)  {
    this.ops = append(this.ops, op{inc, 1})
}

func (this *Fancy) MultAll(m int)  {
    this.ops = append(this.ops, op{0, m})
}

func (this *Fancy) GetIndex(idx int) int {
    if len(this.seq) <= idx { return -1 }
    // actualize the idx'th element
    for i := this.seq[idx].cnt; i < len(this.ops); i++ {
        o := this.ops[i]
        this.seq[idx] = num{
            (this.seq[idx].value * o.times + o.plus) % mod, 
            this.seq[idx].cnt + 1,
        }
    }
    return this.seq[idx].value
}

/**
 * Your Fancy object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Append(val);
 * obj.AddAll(inc);
 * obj.MultAll(m);
 * param_4 := obj.GetIndex(idx);
 */