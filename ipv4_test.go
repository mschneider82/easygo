package easygo

import "testing"

func TestIsIPv4(t *testing.T) {
	var input = "123.123.123.123"
	if IsIPv4(input) {
		if IsIPv4("no valid ipv4 string") {
			t.Errorf("expected not a valid Ipv4 address")
		}
	} else {
		t.Errorf("expected valid Ipv4 address")
	}
}

func TestIPv4AddressInString(t *testing.T) {
	var input = "sent to [1.2.3.4]:25. example.com: 250 Message"
	if !IPv4AddressInString(input) {
		t.Errorf("expected IP in string %v", IPv4AddressInString(input))
	}
}

func TestIPv4ExtractFromString(t *testing.T) {
	var input = "sent to [1.2.3.4]:25. example.com: 250 Message"
	if IPv4ExtractFromString(input) != "1.2.3.4" {
		t.Errorf("expected IP in string %v", IPv4ExtractFromString(input))
	}
}
