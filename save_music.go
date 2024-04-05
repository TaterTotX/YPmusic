package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// SaveInfo 存储信息到用户根目录下的文件中
func SaveInfo(data string) error {
	// 获取用户的根目录
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// 设置文件路径
	filePath := filepath.Join(homeDir, "music_info.txt")

	// 将信息写入文件
	err = ioutil.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		return err
	}

	return nil
}

// ReadInfo 从用户根目录下的文件中读取信息
func ReadInfo() (string, error) {
	// 获取用户的根目录
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// 设置文件路径
	filePath := filepath.Join(homeDir, "music_info.txt")

	// 读取信息
	dataBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(dataBytes), nil
}
