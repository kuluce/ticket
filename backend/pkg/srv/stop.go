package srv

func Stop() error {
	return serviceOperation("stop")
}
