package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func readFile() {
	// ファイルパスを定義
	fileReadPath := "nuc/data001.txt"
	f, err := os.Open(fileReadPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File %s could not be read: %v\n", fileReadPath, err)
	}

	// 関数終了時に閉じる
	defer f.Close()

	// Scannerで読み込む
	// EUCJP形式のテキストを自動でデコードする
	scanner := bufio.NewScanner(transform.NewReader(f, japanese.EUCJP.NewDecoder()))
	fileWritePath := "nuc/destination.txt"
	// 出力先ファイルを書き込みモードでオープンする。存在しなければ新規作成する
	fp, err := os.OpenFile(fileWritePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	// 処理終了時に閉じる
	defer fp.Close()

	writer := bufio.NewWriter(fp)
	for scanner.Scan() {
		_, err2 := writer.WriteString(scanner.Text() + "\n")
		if err2 != nil {
			panic(err2)
		}
	}
	// if werr := writer.Err(); werr != nil {
	// 	fmt.Fprintf(os.Stderr, "FIle %s write error: %v\n", fileWritePath, werr)
	// }
	writer.Flush()

	//
	// 1行ずつコンソールに出力する
	//
	// for scanner.Scan() {
	// 	// 読み込んだテキストを表示
	// 	fmt.Println(scanner.Text())
	// }

	// if serr := scanner.Err(); serr != nil {
	// 	fmt.Fprintf(os.Stderr, "FIle %s scan error: %v\n", filePath, err)
	// }
}

func main() {
	readFile()
}
