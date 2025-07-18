package file

import (
	"blogx_server/global"
	"blogx_server/utils"
	"github.com/pkg/errors"
	"strings"
)

func ImageSuffixJudge(filename string) (suffix string, err error) {
	list := strings.Split(filename, ".")
	if len(list) == 1 {
		err = errors.New("文件名错误")
		return
	}
	suffix = list[len(list)-1]
	l := global.Config.Upload.WhiteList
	if !utils.InList(suffix, l) {
		err = errors.New("文件非法")
		return
	}
	return
}
