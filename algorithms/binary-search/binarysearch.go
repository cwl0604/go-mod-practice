package binary

func rank(key int,a []int)int{
	for lo,hi:=0,len(a)-1;lo<=hi;{
		mid:=lo+(hi-lo)/2
		switch {
		case key<a[mid]:
			hi=mid-1
		case key>a[mid]:
			lo=mid+1
		default: 
			return mid
		}
	}
	return -1
}