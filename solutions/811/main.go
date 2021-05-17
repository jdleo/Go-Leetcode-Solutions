package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	res := []string{}
	// map of each subdomain to its count
	counts := map[string]int{}

	// go through domains
	for _, domain := range cpdomains {
		// get visits, domain using helper func
		visits, d := parse(domain)
		// get subdomains
		subs := strings.Split(d, ".")
		// get d
		d = subs[len(subs)-1]
		// go through length
		for i := len(subs) - 2; i > -1; i-- {
			counts[d] += visits
			d = subs[i] + "." + d
		}
		counts[d] += visits
	}

	// go through counts
	for domain, count := range counts {
		// add formatted string to res
		res = append(res, fmt.Sprintf("%d %s", count, domain))
	}

	return res
}

func parse(domain string) (int, string) {
	res := strings.Split(domain, " ")
	visits, _ := strconv.Atoi(res[0])
	return visits, res[1]
}
