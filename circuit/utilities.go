package circuit

import (
	"os/exec"
	"torctl/collection"
)

func GetExits(t *collection.Connection) string {
	var ip_url = "https://icanhazip.com"
	socks_url := "socks5://" + t.Auth.AuthHost + ":9050"
	b4, _ := exec.Command("curl", "-4", "-x", socks_url, ip_url).Output()
	b6, _ := exec.Command("curl", "-6", "-x", socks_url, ip_url).Output()
	response := "IPv4: " + string(b4) + "\n" + "IPv6: " + string(b6) + "\n"
	t.Exits.V4 = string(b4)
	t.Exits.V6 = string(b6)
	return response
}
