package internal

import "os/exec"

func DownloadLists(asset string) {
	var cmd string
	switch asset {
	case "resolvers":
		cmd = "wget https://raw.githubusercontent.com/trickest/resolvers/refs/heads/main/resolvers-trusted.txt -O lists/resolvers.txt"
	case "subdomains":
		cmd = "wget https://gist.github.com/six2dez/a307a04a222fab5a57466c51e1569acf/raw -O lists/subdomains.txt"
	}

	if err := exec.Command("bash", "-c", cmd).Run(); err != nil {
		panic(err)
	}
}
