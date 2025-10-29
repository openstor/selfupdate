// SPDX-FileCopyrightText: 2025 openstor contributors
// SPDX-FileCopyrightText: 2015-2025 MinIO, Inc.
// SPDX-License-Identifier: Apache-2.0-or-later

// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package osext

import (
	"os"
	"strconv"
	"syscall"
)

func executable() (string, error) {
	f, err := os.Open("/proc/" + strconv.Itoa(os.Getpid()) + "/text")
	if err != nil {
		return "", err
	}
	defer f.Close()
	return syscall.Fd2path(int(f.Fd()))
}
