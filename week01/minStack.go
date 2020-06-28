// 最小栈.
type MinStack struct {
    mainStack []int
    minValueStack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        mainStack:[]int{},
        minValueStack:[]int{},
    }
}


func (this *MinStack) Push(x int)  {
    this.mainStack = append(this.mainStack, x)
    if len(this.minValueStack) == 0 {
        this.minValueStack = append(this.minValueStack, x)
    } else {
        if this.minValueStack[len(this.minValueStack) - 1] < x {
            this.minValueStack = append(this.minValueStack, this.minValueStack[len(this.minValueStack) - 1])
        } else {
            this.minValueStack = append(this.minValueStack, x)
        }
    }
}


func (this *MinStack) Pop()  {
    this.mainStack = this.mainStack[:len(this.mainStack) - 1]
    this.minValueStack = this.minValueStack[:len(this.minValueStack) - 1]
}


func (this *MinStack) Top() int {
    return this.mainStack[len(this.mainStack) - 1]
}


func (this *MinStack) GetMin() int {
    return this.minValueStack[len(this.minValueStack) - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
