package infra

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// Unzip 解压 zip 文件到指定目录
// zipFile: 待解压的 zip 文件路径
// destDir: 解压目标目录
func Unzip(zipFile, destDir string) error {
	// 1. 打开 zip 文件
	r, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer r.Close()

	// 2. 遍历 zip 中的所有文件/目录
	for _, f := range r.File {
		err := unzipFile(f, destDir)
		if err != nil {
			return err
		}
	}
	return nil
}

// 解压单个文件到目标目录
func unzipFile(f *zip.File, destDir string) error {
	// 构建目标文件路径
	destPath := filepath.Join(destDir, f.Name)

	// 处理目录：如果是目录，创建对应的文件夹
	if f.FileInfo().IsDir() {
		// 递归创建目录（包括父目录），权限 0755
		return os.MkdirAll(destPath, 0755)
	}

	// 确保父目录存在（如果文件在子目录中）
	parentDir := filepath.Dir(destPath)
	if err := os.MkdirAll(parentDir, 0755); err != nil {
		return err
	}

	// 打开 zip 中的文件
	srcFile, err := f.Open()
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// 创建目标文件（权限继承自 zip 中的文件，或默认 0644）
	dstFile, err := os.OpenFile(
		destPath,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		f.Mode().Perm(), // 使用 zip 中文件的权限
	)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// 复制文件内容
	_, err = io.Copy(dstFile, srcFile)
	return err
}
