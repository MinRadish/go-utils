/*
 * @Author: Radish
 * @Date: 2025-02-28 11:05
 * @LastEditors: Radish
 * @LastEditTime: 2025-02-28 11:39
 * @Description: 文件相关工具类
 *
 * Copyright (c) 2025 by Radish, All Rights Reserved.
 */

package radishutils

import "os"

// PathExists 判断文件或目录是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
