// SPDX-FileCopyrightText: 2025 openstor contributors
// SPDX-FileCopyrightText: 2015-2025 MinIO, Inc.
// SPDX-License-Identifier: Apache-2.0-or-later

//go:build !windows
// +build !windows

package selfupdate

func hideFile(path string) error {
	return nil
}
