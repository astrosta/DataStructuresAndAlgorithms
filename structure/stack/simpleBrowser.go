package stack

/*
模仿浏览器前进后退
*/

type Browser struct {
	forwardStack  *ArrayStack
	backwardStack *LinkedListStack
}

func NewBrowser() *Browser {
	return &Browser{
		forwardStack:  NewArrayStack(),
		backwardStack: NewLinkedListStack(),
	}
}

func (browser *Browser) CanForward() bool {
	if browser.forwardStack.IsEmpty() {
		return false
	}

	return true
}

func (browser *Browser) CanBack() bool {
	if browser.backwardStack.IsEmpty() {
		return false
	}

	return true
}

func (browser *Browser) Open(addr string) {
	browser.forwardStack.Flush()
}

func (browser *Browser) PushBack(addr string) {
	browser.backwardStack.Push(addr)
}

func (browser *Browser) Forward() {
	if browser.forwardStack.IsEmpty() {
		return
	}

	top := browser.forwardStack.Pop()

	browser.backwardStack.Push(top)
}

func (browser *Browser) Back() {
	if browser.backwardStack.IsEmpty() {
		return
	}

	top := browser.backwardStack.Pop()
	browser.forwardStack.Push(top)
}
