package stack

type FixedCapacityStackOfString struct {
	n int       //capacity
	a []string //stack entries
}

func New(n int) *FixedCapacityStackOfString{
	return &FixedCapacityStackOfString{a:make([]string,n)}
}
// use this function like java constructor func
func (f *FixedCapacityStackOfString) IsEmpty() bool{
	return f.n==0
}

func (f *FixedCapacityStackOfString) Size() int{
	return f.n;
}

func (f *FixedCapacityStackOfString) Push(s string){
	f.a[f.n]=s
	f.n++ //load value first,and go to next index
}
func (f *FixedCapacityStackOfString) Pop() string{
	f.n--
	return f.a[f.n] // move to previous index first,and return value
}