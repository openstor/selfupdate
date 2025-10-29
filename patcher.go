// SPDX-FileCopyrightText: 2025 openstor contributors
// SPDX-FileCopyrightText: 2015-2025 MinIO, Inc.
// SPDX-License-Identifier: Apache-2.0-or-later

package selfupdate

import (
	"io"

	"github.com/openstor/selfupdate/internal/binarydist"
)

// Patcher defines an interface for applying binary patches to an old item to get an updated item.
type Patcher interface {
	Patch(old io.Reader, new io.Writer, patch io.Reader) error
}

type patchFn func(io.Reader, io.Writer, io.Reader) error

func (fn patchFn) Patch(old io.Reader, new io.Writer, patch io.Reader) error {
	return fn(old, new, patch)
}

// NewBSDifferPatcher returns a new Patcher that applies binary patches using
// the bsdiff algorithm. See http://www.daemonology.net/bsdiff/
func NewBSDiffPatcher() Patcher {
	return patchFn(binarydist.Patch)
}
