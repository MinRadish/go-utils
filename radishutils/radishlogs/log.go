/*
 * @Author: Radish
 * @Date: 2025-02-28 11:05
 * @LastEditors: Radish
 * @LastEditTime: 2025-02-28 13:26
 * @Description: 日志
 *
 * Copyright (c) 2025 by Radish, All Rights Reserved.
 */

package radishlogs

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

var (
	Log *loger = &loger{}
)

type loger struct {
	Loger *log.Logger
	Path  string
}

// 初始化
func (l *loger) init() {
	if l.Loger == nil {
		s := l.Path
		if s == "" {
			s = "runtime/logs/"
		}
		l.Loger = initLoger(s)
	}
}

// 记录输出信息到日志文件，并在控制台输出
func (l *loger) Printlns(args ...any) {
	l.init()
	fmt.Println(args...)
	l.Loger.Println(args...)
}

// 记录输出信息到日志文件，
func (l *loger) Println(args ...any) {
	l.init()
	l.Loger.Println(args...)
}

// 设置日志前缀
func (l *loger) SetPrefix(s string) {
	l.init()
	l.Loger.SetPrefix(s)
}

// 分界线
const DIVIDING = "--------------------------------------------------------------------------------------------"

func initLoger(p string) *log.Logger {
	file := getLogDirAndCreate(p)
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	return log.New(logFile, "[Radish]", log.LstdFlags|log.Lshortfile|log.LUTC)
}

// 判断文件是否存在不存在则创建
func getLogDirAndCreate(p string) string {
	n := time.Now()
	today := n.Format("200601")
	pa := p + today
	if ok, _ := PathExists(pa); !ok {
		_ = os.MkdirAll(pa, os.ModePerm)
	}
	fn := n.Format("02") + ".log"

	return pa + "/" + fn
}

// 判断文件是否存在不存在则创建
func getSqlLogDirAndCreate() string {
	n := time.Now()
	today := n.Format("200601")
	pa := "runtime/sql-logs/" + today
	if ok, _ := PathExists(pa); !ok {
		_ = os.MkdirAll(pa, os.ModePerm)
	}
	fn := n.Format("02") + ".log"

	return pa + "/" + fn
}

func getGormLogWriter() logger.Writer {
	file := getSqlLogDirAndCreate()
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	return log.New(logFile, "[SQL]", log.LstdFlags|log.Lshortfile|log.LUTC)
}

// 获取gorm日志
func GetGormLogger() logger.Interface {
	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold:             0,           // 慢 SQL 阈值
		LogLevel:                  logger.Info, // 日志级别
		IgnoreRecordNotFoundError: false,       // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  false,       // 禁用彩色打印
	})
}

// 检测文件是否存在
// 防止嵌套import - 在directory中也有
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
