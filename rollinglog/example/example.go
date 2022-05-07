/*
 * Copyright 2021 Kristian Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"github.com/kristianhuang/go-cmp/rollinglog"
)

func main() {
	opts := &rollinglog.Options{
		Level:            "debug",
		Format:           "json",
		EnableColor:      false,
		DisableCaller:    false,
		OutputPaths:      []string{"test.log", "stdout"},
		ErrorOutputPaths: []string{"error.log"},
		Rolling:          true,
		RollingMaxSize:   1,
	}

	// 初始化全局logger
	rollinglog.Init(opts)
	defer rollinglog.Flush()

	for i := 0; i < 30000; i++ {
		rollinglog.Warnf("This is a formatted %s message", "hello")
		rollinglog.V(rollinglog.InfoLevel).Info("nice to meet you.")
	}

}
