package i2pcontrol

type HiddenService struct{}

func GetHiddenService(name string) (map[string]interface{}, error) {
	return nil, nil
}

func DeleteHiddenService(name string) (string, error) {
	return "", nil
}

func StopAllServices() (string, error) {
	return "", nil
}

func StartAllServices() (string, error) {
	return "", nil
}

func RestartAllServices() (string, error) {
	return "", nil
}

func AddHiddenService(service HiddenService) (string, error) {
	return "", nil
}
