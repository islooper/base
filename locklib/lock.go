package locklib

type (
	Locker interface {
		Lock()   // 堵塞等待直到获取到锁
		Unlock() // 解锁
	}
	Acquirer interface {
		Acquire() bool // 是否获取到锁
		Release() bool // 是否释放锁
	}
)

// 不做任何锁处理， 空实现Locker和Acquirer的方法
type TodoLock struct{}
func (TodoLock) Lock()         {}
func (TodoLock) Unlock()       {}
func (TodoLock) Acquire() bool { return true }
func (TodoLock) Release() bool { return true }
