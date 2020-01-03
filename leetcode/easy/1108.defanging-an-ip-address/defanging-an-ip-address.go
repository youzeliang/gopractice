package _108_defanging_an_ip_address

import "strings"

// https://leetcode-cn.com/problems/defanging-an-ip-address/

func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}
