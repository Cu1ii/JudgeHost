package util

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
)

// IsFileIn 判断某个目录下的文件是否存在
func IsFileIn(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// IsDirectory 判断某个工作目录是否存在
func IsDirectory(dirPath string) (bool, error) {
	fileInfo, err := os.Stat(dirPath)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}

// ClearFileByFolderName 清空某个文件夹下的所有文件
func ClearFileByFolderName(dirPath string) error {
	dir, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}
	for _, d := range dir {
		os.RemoveAll(path.Join(dirPath, d.Name()))
	}
	return nil
}

// ZipDictionary 压缩某个文件夹，并保存到目标位置
func ZipDictionary(zippedPath, targetPath string) (bool, error) {
	if flag, err := IsDirectory(targetPath); !flag || err != nil {
		return false, err
	}
	zipCmd := exec.Command("zip", "-j", "-r", zippedPath, targetPath)
	if err := zipCmd.Run(); err != nil {
		fmt.Println("Execute failed when Wait:" + err.Error())
		return false, err
	}
	return true, nil
}

// UnZipInDictionary  解压缩 zip 至当前文件夹
func UnZipInDictionary(zippedPath, targetPath string) (bool, error) {
	fmt.Println(zippedPath)
	zipCmd := exec.Command("unzip", "-d", targetPath, zippedPath)
	if err := zipCmd.Run(); err != nil {
		fmt.Println("Execute failed when unzip:" + err.Error())
		return false, err
	}
	return true, nil
}

// ReadFileByLines 讲目标文件中的数据按行读取
func ReadFileByLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	br := bufio.NewReader(file)
	var stringList []string
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		stringList = append(stringList, string(line))
	}
	return stringList, nil
}

// ReadFileByLines 讲目标文件中的数据按行读取
func ReadFileByByte(filePath string, byteNumber int64) (string, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return "", err
	}
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return "", err
	}
	if fileInfo.Size() <= byteNumber {
		byteNumber = fileInfo.Size()
	}
	buffer := make([]byte, byteNumber)

	if _, err := file.Read(buffer); err != nil && err != io.EOF {
		logrus.Errorf("read %s error %v", filePath, err)
		return "", err
	}
	if byteNumber > 300 {
		return string(buffer) + "......", err
	}
	return string(buffer), err
}

// DeleteFile 根据 deleteSelf 来判断是否删除自身, 如果是文件夹则判断是清空文件夹还是连带文件夹一起删除
func DeleteFile(filePath string, deleteSelf bool) (bool, error) {
	file, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		log.Println(err.Error())
		return false, err
	}
	if deleteSelf {
		if err := os.RemoveAll(filePath); err != nil {
			return false, err
		}
	} else if file.IsDir() {
		if err := ClearFileByFolderName(filePath); err != nil {
			return false, err
		}
	}
	return true, nil
}

func WriteDataToFilePath(data, filePath string) (bool, error) {
	targetFile, err := os.OpenFile(filePath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return false, err
	}
	defer targetFile.Close()
	if _, err := targetFile.Write([]byte(data)); err != nil {
		fmt.Println("[file.go] = 103: ", err.Error())
		return false, err
	}
	return true, nil
}
