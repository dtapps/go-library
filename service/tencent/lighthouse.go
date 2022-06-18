package tencent

import "log"

type lighthouse struct {
}

func (l lighthouse) DescribeFirewallRules() {
	log.Println("DescribeFirewallRules")
}
