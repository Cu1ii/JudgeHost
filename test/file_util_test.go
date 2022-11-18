package test_test

import (
	"JudgeHost/src/util"
	"log"
	"testing"
)

func TestDeleteFile(t *testing.T) {

	if _, err := util.DeleteFile("/home/cu1/test/aa", true); err != nil {
		log.Println(err.Error())
	}

	if _, err := util.DeleteFile("/home/cu1/test", false); err != nil {
		log.Println(err.Error())
	}
}

func TestWriteDataToFilePath(t *testing.T) {
	if _, err := util.WriteDataToFilePath("this is a test\n", "/home/cu1/test/a.in"); err != nil {
		log.Println(err.Error())
	}

	if _, err := util.WriteDataToFilePath("append this is a test", "/home/cu1/test/a.in"); err != nil {
		log.Println(err.Error())
	}

}
