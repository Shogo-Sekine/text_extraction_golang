package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func process() {
	// ファイルパスを定義
	fileReadPath := "nuc/"

	// 読み出し対象のファイルリストを読み込む
	fileList := readDir(fileReadPath)

	// 出力先ファイルを書き込みモードでオープンする。存在しなければ新規作成する
	fileWritePath := "nuc/destination.txt"
	fp, err := os.OpenFile(fileWritePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	// 処理終了時に閉じる
	defer fp.Close()

	writer := bufio.NewWriter(fp)

	for _, file := range fileList {
		// 読み込むファイルのパスを生成する
		filePath := fileReadPath + file.Name()
		f, err := os.Open(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "File %s could not be read: %v\n", fileReadPath, err)
		}
		// Scannerで読み込む
		// EUCJP形式のテキストを自動でデコードする
		scanner := bufio.NewScanner(transform.NewReader(f, japanese.EUCJP.NewDecoder()))
		// 読み込んだファイルのテキストを1行ずつスキャンして出力先のファイルに追記する
		for scanner.Scan() {
			_, err2 := writer.WriteString(scanner.Text() + "\n")
			if err2 != nil {
				panic(err2)
			}
		}
		writer.Flush()
	}
	fmt.Println("finished...")
}

func readDir(fileReadPath string) []os.FileInfo {
	fileList, _ := ioutil.ReadDir(fileReadPath)
	return fileList
}

func main() {
	process()
}
