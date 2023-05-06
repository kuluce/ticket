package srv

func Uninstall() error {
	return serviceOperation("uninstall")
}
