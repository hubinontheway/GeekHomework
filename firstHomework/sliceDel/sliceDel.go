package sliceDel

// 封装一下对于slice的操作，使用泛型并且进行约束
type MySlice[T any] struct {
	Slice []T
}

// 返回一个初始容量为0的泛型切片
func Constructor[T any](elements ...T) *MySlice[T] {
	s := make([]T, 0)
	s = append(s, elements...)
	return &MySlice[T]{
		Slice: s,
	}
}

// 删除操作
func (m *MySlice[T]) DeleteAtIndex(index int) {
	if index < 0 || index > len(m.Slice) {
		return
	}

	// 重新组成为一个slice
	m.Slice = append(m.Slice[:index], m.Slice[index+1:]...)

	// 检查是否需要缩小容量
	if cap(m.Slice) > 4*len(m.Slice) {
		newSlice := make([]T, len(m.Slice))
		copy(newSlice, m.Slice)
		m.Slice = newSlice
	}
}
