/*
 * @Author: Radish
 * @Date: 2025-02-28 14:10
 * @LastEditors: Radish
 * @LastEditTime: 2025-02-28 14:18
 * @Description: 测试字符串工具类
 * - go test ./... -v 测试所有包
 * Copyright (c) 2025 by EIZ, All Rights Reserved.
 */

package radishutils

import "testing"

func TestCapitalizeFirstLetterUnicode(t *testing.T) {
	t.Log(CapitalizeFirstLetterUnicode("radish"))
	t.Log(CapitalizeFirstLetterUnicode("Radish"))
	t.Log(CapitalizeFirstLetterUnicode("rADISH"))
	t.Log(CapitalizeFirstLetterUnicode("RADISH"))
	t.Log(CapitalizeFirstLetterUnicode("r"))
	t.Log(CapitalizeFirstLetterUnicode("R"))
	t.Log(CapitalizeFirstLetterUnicode(""))
}
