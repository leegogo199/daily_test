package UnionFind
type Unionfind struct{
	id []int
}
func(u *Unionfind)Create(n int) []int{
	 id:=make([]int,n)
	 for i:=0;i<n;i++{
	 	id[i]=i
	 }
	 return id
}
func (u *Unionfind)Root(i int) int{
	for i!=u.id[i]{
		i=u.id[i]
	}
	return i
}
func (u *Unionfind)Connected(p,q int)bool{
	return u.Root(p)==u.Root(q)
}
func(u *Unionfind)Union(p,q int){
	var i,j=u.Root(p),u.Root(q)
	u.id[i]=j
}


