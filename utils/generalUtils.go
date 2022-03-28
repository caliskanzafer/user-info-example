package utils

import "log"

func checkError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
