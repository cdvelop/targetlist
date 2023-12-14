package targetlist

func (l targetList) ResetViewHandlerObject() (err string) {
	_, err = l.reset.CallWithEnableAndQueryParams(l.Object)
	return
}
