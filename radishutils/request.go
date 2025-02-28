/*
 * @Author: Radish
 * @Date: 2025-02-28 11:05
 * @LastEditors: Radish
 * @LastEditTime: 2025-02-28 11:38
 * @Description:请求工具
 *
 * Copyright (c) 2025 by Radish, All Rights Reserved.
 */

package radishutils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/MinRadish/go-utils/radishutils/radishlogs"
	"github.com/google/uuid"
)

// 请求方法
type Method string

const (
	GET  Method = "GET"
	POST Method = "POST"
	PUT  Method = "PUT"
)

// 请求头
//
// Key: 请求头key
//
// Val: 请求头value
type Headers struct {
	Key string
	Val string
}

// 请求参数
//
// Path: 请求路径
//
// Method: 请求方法
//
// Debug: 是否开启调试
type Params struct {
	Path   string `url:"path"`
	Method `url:"method"`
	Debug  bool
}

// ExecAsk 执行请求
func ExecAsk(p *Params, b *bytes.Buffer, h []Headers) []byte {
	uid := uuid.New()
	bl := *b
	if p.Debug {
		fmt.Println(p, b)
	}
	req, _ := http.NewRequest(string(p.Method), p.Path, b)
	if len(h) > 0 {
		for _, v := range h {
			req.Header.Set(v.Key, v.Val)
		}
	}
	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		fmt.Println(uid, "ExecAsk-error", err)
		radishlogs.Log.Println(uid.String() + radishlogs.DIVIDING)
		radishlogs.Log.Println(uid, p, bl.String(), h)
		radishlogs.Log.Println(uid.String() + radishlogs.DIVIDING)
	} else {
		defer resp.Body.Close()
		b, _ := io.ReadAll(resp.Body)
		if p.Debug {
			fmt.Println(uid, "resp.Body", time.Now().Format(time.DateTime), string(b))
		}
		return b
	}
	return nil
}
