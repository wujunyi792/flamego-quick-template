package ip

import (
	"github.com/wujunyi792/flamego-quick-template/pkg/colorful"
	"net"
)

func GetLocalHost() (res []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		println(colorful.Red("net.Interfaces failed, err: " + err.Error()))
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						res = append(res, ipnet.IP.String())
					}
				}
			}
		}

	}
	return
}
