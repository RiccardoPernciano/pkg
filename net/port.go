// Copyright (c) 2015-2021 MinIO, Inc.
//
// This file is part of MinIO Object Storage stack
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package net

import (
	"errors"
	"net"
	"strconv"
)

// Port - network port
type Port uint16

// String - returns string representation of port.
func (p Port) String() string {
	return strconv.Itoa(int(p))
}

// GetFreePort asks the kernel for a free open port that is ready to use.
func GetFreePort() (Port, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return Port(l.Addr().(*net.TCPAddr).Port), nil
}

// GetNextFreePort - returns the immediate next port if it's available. The maximum allowed input is 65534, since we increment the value by 1.
func GetNextFreePort(port string) (Port, error) {
	if port == "" || port == "0" {
		return 0, errors.New("invalid starting port")
	}
	p, err := strconv.Atoi(port)
	if err != nil || p <= 0 || p >= 65535 {
		return 0, errors.New("invalid port number (must be between 1 and 65534)")
	}
	nextPort := p + 1
	addr := &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: nextPort}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return Port(l.Addr().(*net.TCPAddr).Port), nil
}

// ParsePort - parses string into Port
func ParsePort(s string) (p Port, err error) {
	switch s {
	case "https":
		return Port(443), nil
	case "http":
		return Port(80), nil
	}

	var i int
	if i, err = strconv.Atoi(s); err != nil {
		return p, errors.New("invalid port number")
	}

	if i < 0 || i > 65535 {
		return p, errors.New("port must be between 0 to 65535")
	}

	return Port(i), nil
}
