package helper

import "github.com/sirupsen/logrus"

func CheckError(keyword string, err error) {
	if err != nil {
		logrus.Error(keyword, ":", err)
	}

}
