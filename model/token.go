// Copyright 2017 Frédéric Guillot. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package model

// Token represents a CSRF token in the system.
type Token struct {
	ID    string
	Value string
}