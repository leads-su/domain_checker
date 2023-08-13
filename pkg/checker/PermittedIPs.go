package checker

// список разрешенных ip адресов в A записи домена
type PermittedIPs struct {
	ips []string
}

func CreatePermittedIps() *PermittedIPs {
	return &PermittedIPs{}
}

func (pips *PermittedIPs) Add(ip string) *PermittedIPs {
	pips.ips = append(pips.ips, ip)
	return pips
}

func (pips *PermittedIPs) Contains(ip string) bool {
	for _, x := range pips.ips {
		if ip == x {
			return true
		}
	}
	return false
}

func (pips *PermittedIPs) Count() int {
	return len(pips.ips)
}
