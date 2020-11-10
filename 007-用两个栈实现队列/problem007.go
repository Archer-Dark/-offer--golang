package problem007

import "container/list"

type CQueue struct {
	stact1,stact2 *list.List
}


func Constructor() CQueue {
	return CQueue{
		stact1: list.New(),
		stact2: list.New(),
	}
}


func (this *CQueue) AppendTail(value int)  {
	this.stact1.PushBack(value)
}


func (this *CQueue) DeleteHead() int {
	if this.stact2.Len() == 0{
		for this.stact1.Len() > 0 {
			this.stact2.PushBack(this.stact1.Remove(this.stact1.Back()))
		}
	}
	if this.stact2.Len() > 0 {
		e := this.stact2.Back()
		this.stact2.Remove(e)
		return e.Value.(int)
	}
	return -1

}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
