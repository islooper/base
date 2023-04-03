package util

type Empty struct{}

var empty Empty

//set类型
type Set struct {
	m map[interface{}]Empty
}

func NewSet() *Set {
	return &Set{
		m: map[interface{}]Empty{},
	}
}

//添加元素
func (s *Set) Add(vals ...interface{}) {
	if len(vals) == 0 {
		return
	}
	for _, v := range vals {
		s.m[v] = empty
	}
}

//删除元素
func (s *Set) Remove(val interface{}) {
	delete(s.m, val)
}

//获取长度
func (s *Set) Len() int {
	return len(s.m)
}

//清空set
func (s *Set) Clear() {
	s.m = make(map[interface{}]Empty)
}

type TraverseFunc func(x interface{})

//遍历set
func (s *Set) Traverse(x TraverseFunc) {
	for k, _ := range s.m {
		x(k)
	}
}

// 查询是否包含
func (s *Set) Contains(v interface{}) bool {
	_, ok := s.m[v]
	return ok
}
