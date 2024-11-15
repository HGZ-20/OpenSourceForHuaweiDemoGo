//go:build windows
// +build windows

/*
 * Copyright (c) Huawei Technologies Co., Ltd. 2024-2030. All rights reserved.
 * window tcp connect
 */

package pq

import (
	"fmt"
	"syscall"
)

/* TCP_SYNCNT parameter is used to control the number of retransmission times */
const TCP_SYNCNT = 0x7

func setSocketParameter(retrytimes int64) func(string, string, syscall.RawConn) error {
	return func(network, address string, c syscall.RawConn) error {
		var syscallErr error
		err := c.Control(func(fd uintptr) {
			syscallErr = syscall.SetsockoptInt(syscall.Handle(int(fd)), syscall.IPPROTO_TCP, TCP_SYNCNT, int(retrytimes))
			if syscallErr != nil {
				syscallErr = fmt.Errorf("setsockopt TCP_SYNCTNT : %s)", retrytimes)
				return
			}
		})
		if err != nil {
			return err
		}
		return syscallErr
	}
}
