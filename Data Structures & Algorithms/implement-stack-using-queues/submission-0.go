type Node struct {
    val  int
    next *Node
}

type MyStack struct {
    q *Node
}

func Constructor() MyStack {
    return MyStack{q: nil}
}

func (this *MyStack) Push(x int) {
    this.q = &Node{val: x, next: this.q}
}

func (this *MyStack) Pop() int {
    if this.q == nil {
        return -1
    }
    top := this.q.val
    this.q = this.q.next
    return top
}

func (this *MyStack) Top() int {
    if this.q == nil {
        return -1
    }
    return this.q.val
}

func (this *MyStack) Empty() bool {
    return this.q == nil
}