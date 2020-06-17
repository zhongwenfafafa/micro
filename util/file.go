package util

import (
	"errors"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"micro/defind"
)

const (
	// MergeFileCMD：合并文件
	MergeFileCMD = `
	#!/bin/bash
	# 需要进行合并的分片所在的目录
	chunkDir=$1
	# 合并后的文件的完整路径（目录 + 文件名）
	mergePath=$2
	
	echo "分块合并，输出目录：" $chunkDir
	
	if [! -f $mergePath]; then
		echo "$mergePath not exist"
	else
		rm -f $mergePath
	fi

	for chunk in $(ls $chunkDir | sort -n)
	do
		cat $chunkDir/${chunk} >> ${mergePath}
	done

	echo "合并完成，输出：" mergePath
	`

	// FileSha1CMD：计算文件sha1值
	FileSha1CMD = `
	#!/bin/bash
	sha1sum $1 | awk '{print $1}'
	`

	// FileSizeCMD：计算文件大小
	FileSizeCMD = `
	#!/bin/bash
	ls -l $1 | awk '{print $5}'
	`

	// FileChunkDelCMD：删除文件快
	FileChunkDelCMD = `
	#!/bin/bash
	chunkDir="#CHUNKDIR#"
	targetDir=$1
	# 增加条件判断，避免误删（指定的路包含且不等于chunkDir）
	if [[ $targetDir =~ $chunkDir ]] && [[ $targetDir != $chunkDir ]]; then
		rm -rf $targetDir
	fi
	`
)

// 通过调用shell来删除指定目录
func RemovePathByShell(destPath string) (bool, error) {
	cmdStr := strings.Replace(FileChunkDelCMD, "$1", destPath, 1)
	cmdStr = strings.Replace(cmdStr, "#CHUNKDIR#",
		defind.CHUNK_LOCAL_ROOT_DIR, 1)
	delCmd := exec.Command("bash", "-c", cmdStr)

	if _, err := delCmd.Output(); err != nil {
		return false, err
	}

	return true, nil
}

// 通过调用shell来计算文件大小
func ComputeFileSizeByShell(destPath string) (int, error) {
	cmdStr := strings.Replace(FileSizeCMD, "$1", destPath, 1)
	fSizeCmd := exec.Command("bash", "-c", cmdStr)

	if fSizeStr, err := fSizeCmd.Output(); err != nil {
		return 0, err
	} else {
		reg := regexp.MustCompile("\\s+")
		fSize, err := strconv.Atoi(
			reg.ReplaceAllString(string(fSizeStr), ""))
		if err != nil {
			return 0, err
		}

		return fSize, nil
	}
}

// 通过调用shell来计算文件hash
func ComputeFileSha1ByShell(destPath string) (string, error) {
	cmdStr := strings.Replace(FileSha1CMD, "$1", destPath, 1)
	fHashCmd := exec.Command("bash", "-c", cmdStr)
	if fileHash, err := fHashCmd.Output(); err != nil {
		return "", err
	} else {
		reg := regexp.MustCompile("\\s+")
		return reg.ReplaceAllString(string(fileHash), ""), nil
	}
}

// 通过调用shell来合并分块文件
func MergeChunksByShell(chunkDir, destPath, fileSha1 string) (bool, error) {
	// 合并分块
	cmdStr := strings.Replace(MergeFileCMD, "$1", chunkDir, 1)
	cmdStr = strings.Replace(cmdStr, "$2", destPath, 1)
	mergeCmd := exec.Command("bash", "-c", cmdStr)
	if _, err := mergeCmd.Output(); err != nil {
		return false, err
	}

	// 计算合并后的文件hash
	if fileHash, err := ComputeFileSha1ByShell(destPath); err != nil {
		return false, err
	} else if string(fileHash) != fileSha1 {
		return false, errors.New("merge after file hash not eq dest file hash")
	}

	return true, nil
}
