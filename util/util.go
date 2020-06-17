package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
)

// 生成md5
func MD5(data []byte) string {
	_md5 := md5.New()
	_md5.Write(data)

	return hex.EncodeToString(_md5.Sum([]byte("")))
}

// 判断某个元素是否在slice、array、map中
func Contain(target interface{}, obj interface{}) (bool, error) {
	targetVal := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice,reflect.Array:
		for i := 0; i < targetVal.Len(); i++ {
			if targetVal.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetVal.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	default:
		// TODO 记录Warn日志
	}

	return false, fmt.Errorf("%+v not in this array/slice/map", obj)
}
