package dsa

type MyQueue struct {
	s1 []int
	s2 []int
}

func Constructor() MyQueue {
	return MyQueue{
		s1: make([]int, 0),
		s2: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	for len(this.s1) != 0 {
		this.s2 = append(this.s2, this.s1[0])
		this.s1 = this.s1[1:]
	}
	this.s2 = append(this.s2, x)
	for len(this.s2) != 0 {
		this.s1 = append(this.s1, this.s2[0])
		this.s2 = this.s2[1:]
	}
}

func (this *MyQueue) Pop() int {
	res := this.s1[0]
	this.s1 = this.s1[1:]
	return res
}

func (this *MyQueue) Peek() int {
	return this.s1[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.s1) == 0
}
