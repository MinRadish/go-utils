/*
 * @Author: Radish
 * @Date: 2025-02-28 11:05
 * @LastEditors: Radish
 * @LastEditTime: 2025-02-28 11:24
 * @Description: 读取execl文件数据
 *
 * Copyright (c) 2025 by Radish, All Rights Reserved.
 */

package radishutils

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// ReadExcel 读取excel文件
//
// p: 文件路径
//
// return: 二维数组
func ReadExcel(p string) [][]string {
	f, err := excelize.OpenFile(p)
	if err != nil {
		panic(err)
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	var all [][]string
	for _, v := range f.GetSheetList() {
		rows, err := f.GetRows(v)
		if err != nil {
			panic(err)
		}
		rows = rows[1:]
		all = append(all, rows...)
	}

	return all
}
