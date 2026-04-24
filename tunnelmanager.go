package i2pcontrol

type HiddenService struct{}

// ServiceAction performs an action on a tunnel.
func ServiceAction(name, action string, toAll bool) (string, map[string]interface{}, error) {
	retpre, err := Call("TunnelManager", map[string]interface{}{
		"Name":   name,
		"Action": action,
		"All":    toAll,
	})
	if err != nil {
		return "", nil, err
	}
	var tunnelOptions map[string]interface{}
	result := retpre["status"].(string)

	// a get action returns the tunnel options, this is the only action that returns them
	if action == "get" {
		tunnelOptions = retpre["i2p.router.net.tunnels.i2ptunnel.options"].(map[string]interface{})
	}

	return result, tunnelOptions, nil
}

func AddHiddenService(service HiddenService) (string, error) {
	return "", nil
}
