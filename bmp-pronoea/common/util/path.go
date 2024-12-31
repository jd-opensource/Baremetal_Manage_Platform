package util

import (
"errors"
"fmt"
"io"
"io/ioutil"
"os"
"path"
"path/filepath"
"syscall"
)

// GetParentPath 获取父级目录
// path format is /path/to/file, return /path/to
func GetParentPath(pathStr string) string {
	return path.Dir(pathStr)
}

// CheckIsDir 是目录吗？
func CheckIsDir(filename string) error {
	fi, err := os.Stat(filename)
	if err != nil {
		return err
	}
	if !fi.IsDir() {
		return errors.New("The given path is not directory.")
	}

	return nil
}

// PathIsExist 目录存在吗
func PathIsExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}


func PathIsEmpty(path string) (bool,error) {
	if !PathIsExist(path) {
		return true, nil             //目录不存在，认为目录为空
	}

	dirlist, err := ioutil.ReadDir(path)
	if err != nil {
		return true, err
	}
	if len(dirlist) > 0 {
		return false, nil
	}
	return true, nil
}

// CheckIsFile 是文件吗
func CheckIsFile(filename string) error {
	fi, err := os.Stat(filename)
	if err != nil {
		return err
	}
	if !fi.Mode().IsRegular() {
		return errors.New("The given path is not file.")
	}

	return nil
}

// FileIsExist 文件存在吗
func FileIsExist(path string) bool {
	if !PathIsExist(path) {
		return false
	}

	err := CheckIsFile(path)
	if err != nil {
		return false
	}

	return true
}

func RemoveFile(path string) error {
	if !FileIsExist(path) {    //文件不存在认为是已经删除了
		return fmt.Errorf("file not exist")
	}

	err := os.Remove(path)
	return err
}

// CheckParent make sure parents dir exist, not exist then create
func CheckParent(aph string) error {
	pph := path.Dir(aph)
	err := os.MkdirAll(pph, os.ModeDir|0755)
	if nil != err {
		return err
	}
	return nil
}

// WriteFile 写文件
func WriteFile(filename string, content string) error {
	err := CheckParent(filename)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	ret, err := f.WriteString(content)
	if err != nil || ret != len(content) {
		msg := fmt.Sprintf("WriteFile[%v] err.", filename)
		return errors.New(msg)
	}

	return nil
}

func DiffDir(excludePath, spath, dpath string) (bool, string, error) {
	return diffByDiffCmd(excludePath, spath, dpath)
}

func RsyncDir(spath, dpath string) error {
	isdir, err := isDir(spath)
	if err != nil {
		return err
	}

	command := ""
	if isdir {
		command = fmt.Sprintf("mkdir -p %v %v && rsync  -rtopg --delete  %v/  %v/", spath, dpath, spath, dpath)
	} else {
		command = fmt.Sprintf("rsync  -avrtopg --delete  %v  %v", spath, dpath)
	}

	_, err = CmdRunWithOutput(command)
	return err
}

func RsyncDirWithOutPut(spath, dpath string) (string, error) {
	isdir, err := isDir(spath)
	if err != nil {
		return "", err
	}

	command := ""
	if isdir {
		command = fmt.Sprintf("mkdir -p %v %v && rsync  -rtopg --delete  %v/  %v/", spath, dpath, spath, dpath)
	} else {
		command = fmt.Sprintf("rsync  -avrtopg --delete  %v  %v", spath, dpath)
	}

	outPut, err := CmdRunWithOutput(command)
	if err != nil {
		return outPut, fmt.Errorf("RsyncDir error! command[%v], output:[%v], error:[%v]", command, outPut, err)
	}
	return outPut, err
}

func diffByDiffCmd(excludePath, spath, dpath string) (bool, string, error) {
	if !PathIsExist(spath) {
		msg := fmt.Sprintf("spath[%v] not exist.", spath)
		return false, "", errors.New(msg)
	}
	if !PathIsExist(dpath) {
		// dpath is not exist
		return false, "", nil
	}

	command := fmt.Sprintf("diff -aHr -x '%v' %v %v", excludePath, spath, dpath)
	diffStr, err := CmdRunWithOutput(command)
	if err != nil {
		if err.Error() == "exit status 1" {
			return false, diffStr, nil
		}

		msg := fmt.Sprintf("Exec diff  %v %v %v error  %v.\n", spath, dpath, command, err)
		return true, diffStr, errors.New(msg)
	}

	return true, diffStr, nil
}

func isDir(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	if fi.IsDir() {
		return true, nil
	}
	if fi.Mode().IsRegular() {
		return false, nil
	}
	return false, errors.New("File is not regular file")
}

func diffByRsyncCmd(spath, dpath string) bool {
	return true
}

func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)

	return err
}


func CopyDirectory(scrDir, dest string) error {
	entries, err := ioutil.ReadDir(scrDir)
	if err != nil {
		return err
	}
	if !PathIsExist(dest) {
		os.MkdirAll(dest, os.ModeDir|0755)
	}
	for _, entry := range entries {
		sourcePath := filepath.Join(scrDir, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			return err
		}

		stat, ok := fileInfo.Sys().(*syscall.Stat_t)
		if !ok {
			return fmt.Errorf("failed to get raw syscall.Stat_t data for '%s'", sourcePath)
		}

		switch fileInfo.Mode() & os.ModeType{
		case os.ModeDir:
			if err := CreateIfNotExists(destPath, 0755); err != nil {
				return err
			}
			if err := CopyDirectory(sourcePath, destPath); err != nil {
				return err
			}
		case os.ModeSymlink:
			if err := CopySymLink(sourcePath, destPath); err != nil {
				return err
			}
		default:
			if err := Copy(sourcePath, destPath); err != nil {
				return err
			}
		}

		if err := os.Lchown(destPath, int(stat.Uid), int(stat.Gid)); err != nil {
			return err
		}

		isSymlink := entry.Mode()&os.ModeSymlink != 0
		if !isSymlink {
			if err := os.Chmod(destPath, entry.Mode()); err != nil {
				return err
			}
		}
	}
	return nil
}

func Copy(srcFile, dstFile string) error {
	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}

	defer out.Close()

	in, err := os.Open(srcFile)
	defer in.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}

func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

func CreateIfNotExists(dir string, perm os.FileMode) error {
	if Exists(dir) {
		return nil
	}

	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}

	return nil
}

func CopySymLink(source, dest string) error {
	link, err := os.Readlink(source)
	if err != nil {
		return err
	}
	return os.Symlink(link, dest)
}