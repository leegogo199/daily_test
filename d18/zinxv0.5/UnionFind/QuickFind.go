package UnionFind
type UnionFind struct{
	id []int
}
func (u *UnionFind)Create( n int) ([]int,error){
	 var id =make([]int,n)
	for i:=0;i<n;i++{
		id[i]=i
	}
	return id,nil
}
func (u *UnionFind)Connected(p,q int)bool{
	return u.id[p]==u.id[q]
}
func (u *UnionFind)Union(p,q int){
	var pid,qid=u.id[p],u.id[q]
	for i:=0;i<len(u.id);i++{
		if u.id[i]==pid{
			u.id[i]=qid
		}
	}
}

