/*
 * @Author: Radish
 * @Date: 2025-02-28 11:05
 * @LastEditors: Radish
 * @LastEditTime: 2025-02-28 11:07
 * @Description: csv工具类
 *
 * Copyright (c) 2025 by Radish, All Rights Reserved.
 */

package radishutils

import (
	"encoding/csv"
	"os"
)

// Read 读取csv文件
func Read(fPath string) [][]string {
	f, err := os.Open(fPath)
	if err != nil {
		panic(err)
	}
	defer func() {
		f.Close()
	}()
	r := csv.NewReader(f)
	c, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	return c
}
