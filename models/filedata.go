package models

import (
	"fmt"
	"os"
)

func GetFileModifiedTime(filename string) ( err error,modifiedtime int64) {
	file, err := os.Stat(filename)
	if err == nil {
		modifiedtime = file.ModTime().Unix()
	}
	return err,modifiedtime
}

func GetFileUrlWithModifiedTime(filename string) ( err error,path string){
	err,modifiedtime:=GetFileModifiedTime(filename)
	if err ==nil{
		path="/"+filename+"?_"+fmt.Sprint(modifiedtime)
	}
	return err,path
}