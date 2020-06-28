// 循环双端队列.
type MyCircularDeque struct {
    leftIndex int
    rightIndex int
    values []int
    count int
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
    return MyCircularDeque{
        leftIndex : 0,
        rightIndex : 0,
        values : make([]int, k + 1),
        count : k + 1,
    }
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsFull() {
        return false
    }
    this.leftIndex--
    if this.leftIndex < 0 {
        this.leftIndex = this.count - 1
    }
    this.values[this.leftIndex] = value
    return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.IsFull() {
        return false
    }
    this.values[this.rightIndex] = value
    this.rightIndex++
    if this.rightIndex >= this.count {
        this.rightIndex = this.rightIndex % this.count
    }
    return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty() {
        return false
    }
    this.values[this.leftIndex] = 0
    this.leftIndex++
    if this.leftIndex >= this.count {
        this.leftIndex = this.leftIndex % this.count
    }
    return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty() {
        return false
    }
    this.rightIndex--
    if this.rightIndex < 0 {
        this.rightIndex = this.count - 1
    }
    this.values[this.rightIndex] = 0
    return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty() {
        return -1
    }

    return this.values[this.leftIndex]
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty() {
        return -1
    }

    index := this.rightIndex - 1
    if index < 0 {
        index = this.count - 1
    }
    return this.values[index]
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
    if this.rightIndex == this.leftIndex {
        return true
    }
    return false
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
     if (this.rightIndex + 1) % this.count ==  this.leftIndex {
        return true
    }
    return false
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
