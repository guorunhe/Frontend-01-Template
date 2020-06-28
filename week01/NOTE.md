###### 本周学习到了许多提升效率的方法。之前也自己刷过leetcode，但是总是死磕，没坚持多久就放弃了。现在按照覃超老师讲的“五毒掌”来效果好多了。
### 下面是是关于Java的Queue、PriorityQueue的学习总结：
##### 一、Queue
- Queue 即队列，它的元素遵循先进先出（FIFO）的原则。
- Queue是一个继承于Collection的接口。
- Queue定义了6个方法，分别是：add()、offer()、remove()、poll()、element()、peek()
- add(E e)方法用于在容量允许的情况下插入一个元素e到队尾，成功返回ture，失败抛出IllegalStateException 异常    时间复杂度O(1)
- offer(E e) 方法用于在容量允许的情况下插入一个元素e到队尾，成功返回ture，失败返回null    时间复杂度O(1)
- remove() 该方法返回queue的头部元素，如果queue为空，抛 NoSuchElementException异常    时间复杂度O(1)
- poll()方法与remove()方法的区别就是该方法失败返回null
- element(),返回queue的头部元素但是不做删除操作，如果队列为空的话抛出NoSuchElementException异常    时间复杂度O(1)
- peek(),该方法与element唯一的的区别为queue为空，该方法返回null    时间复杂度O(1)
##### 二、PriorityQueue
- PriorityQueue为优先队列，队中元素按照自然排序、或者通过比较器初始化优先级，优先队列不嫩插入为空的元素，即使依赖自然排序，也不允许插入不可以比较的对象，队列头部为指定排序规则最小的元素。
- PriorityQueue不是线程安全的
- add(E e),插入一个元素在该优先队列中
- offer(E e)，插入一个
- comparator() 返回队列用于排序的比较器，如果是自然排序，则返回null
- contains(Object o) 如果队列中包含o,则返回true
- interator()，返回队列元素的迭代器
- clear()，清空该队列
- toArray(),返回所有元素的queue