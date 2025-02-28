/*
 * @Author: Radish
 * @Date: 2025-02-28 11:05
 * @LastEditors: Radish
 * @LastEditTime: 2025-02-28 11:08
 * @Description: 字符串操作工具
 *
 * Copyright (c) 2025 by Radish, All Rights Reserved.
 */

package radishutils

import (
	"strings"
	"unicode"
)

// 将字符串首字母大写
func CapitalizeFirstLetterUnicode(s string) string {
	if len(s) == 0 {
		return s
	}
	s = strings.ToLower(s)
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
