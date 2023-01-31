// Copyright (C) 2023 Storj Labs, Inc.
// See LICENSE for copying information.

package server

import (
	"syscall"

	"go.uber.org/zap"
)

const tcpFastOpenServer = 15

func setTCPFastOpen(fd uintptr, queue int) error {
	return syscall.SetsockoptInt(syscall.Handle(fd), syscall.IPPROTO_TCP, tcpFastOpenServer, 1)
}

func tryInitFastOpen(*zap.Logger) {}
