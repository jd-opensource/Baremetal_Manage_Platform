package cron

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

const REGEX_FILE_NAME string = `^.*\.202[0-9]{7}$`
const LOG_DIR = "./log"

func doCleanLogCron() error {

	files := []string{}
	infos := []os.FileInfo{}

	remainDays, _ := beego.AppConfig.Int("logs.remainDays")
	if remainDays == 0 {
		remainDays = 3 //测试时默认3天
	}

	if err := filepath.Walk(LOG_DIR, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		infos = append(infos, info)
		return nil
	}); err != nil {
		fmt.Println("doCleanLogCron.Walk error:", err.Error())
		return err
	}

	for _, info := range infos {
		if info.IsDir() {
			continue
		}
		if match, _ := regexp.MatchString(REGEX_FILE_NAME, info.Name()); match {
			modifySince := time.Now().Unix() - info.ModTime().Unix()
			if modifySince > int64(remainDays)*24*3600 {
				if err := os.Remove(LOG_DIR + string(os.PathSeparator) + info.Name()); err != nil {
					fmt.Printf("os.Remove error, file:%s, error:%s\n", info.Name(), err.Error())
				}
			}
		}
	}
	return nil
}
