package targetlist

func (l targetList) ResetViewHandlerObject() (err string) {

	return l.reset.ExecuteJsFun(l.Object)
}
