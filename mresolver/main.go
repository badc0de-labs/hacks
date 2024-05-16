package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)
func main() {
	showUnresolved := flag.Bool("u", false, "Show unresolved names")
	flag.Parse()
	var unresolved []string
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan(){
		domain := sc.Text()
		ips, err := net.LookupHost(domain)
		if err != nil {
			unresolved = append(unresolved, domain)
			continue
		}
		fmt.Println(domain, ips)
	}
	if *showUnresolved {
		fmt.Println("\nDomains unable to resolve: ",unresolved)
	}

}
