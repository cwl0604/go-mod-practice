package stack

type FixedCapacityStackOfString struct {
	n int       //capacity
	a []interface{} //stack entries
}
// use this function like java constructor func
func New(n int) *FixedCapacityStackOfString{
	return &FixedCapacityStackOfString{a:make([] interface{},n)}
}

func (f *FixedCapacityStackOfString) IsEmpty() bool{
	return f.n==0
}

func (f *FixedCapacityStackOfString) Size() int{
	return f.n;
}

func (f *FixedCapacityStackOfString) Push(s interface{}){
	f.a[f.n]=s
	f.n++ //load value first,and go to next index
}
func (f *FixedCapacityStackOfString) Pop() interface{}{
	f.n--
	return f.a[f.n] // move to previous index first,and return value
}