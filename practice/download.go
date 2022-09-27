package practice

import (
	"fmt"
	"github.com/qianlnk/pgbar"
	"io"
	"os"
	"regexp"
	"strconv"
	"sync"
)

func OMADownloadPractice() {
	// 将下载的文件名
	downloadFileName := "./123.zip"
	// 复制文件名
	copyFilename := "./copy.zip"
	storageFilename := "./storage.txt"
	// 打开文件
	downloadFile, err := os.Open(downloadFileName)
	if err != nil {
		panic(err)
	}
	defer downloadFile.Close()
	// 获取文件大小
	info, _ := downloadFile.Stat()
	downloadSize := info.Size()
	var count int64 = 1
	if downloadSize%5 == 0 {
		count *= 5
	} else {
		count *= 10
	}
	// 获取每个协程处理文件大小
	var perG = downloadSize / count
	fmt.Printf("文件总大小: %v,分片数:%v, 每个分片大小:%v", downloadSize, count, perG)
	// open copy file
	copyFile, err := os.OpenFile(copyFilename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer copyFile.Close()
	storageFile, err := os.OpenFile(storageFilename, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// 协程处理
	var currentIndex int64 = 0
	wg := sync.WaitGroup{}
	fmt.Println("协程进度条")
	pgb := pgbar.New("")
	for ; currentIndex < count; currentIndex++ {
		wg.Add(1)
		go func(current int64) {
			p := pgb.NewBar(fmt.Sprint(current+1)+"st", int(perG))
			b := make([]byte, 1024)
			bs := make([]byte, 16)
			currentIndex, _ := storageFile.ReadAt(bs, current*16)
			//取出所有整数
			reg := regexp.MustCompile(`\d+`)
			countStr := reg.FindString(string(bs[:currentIndex]))
			total, _ := strconv.ParseInt(countStr, 10, 0)
			progressBar := 1
			for {
				if total >= perG {
					wg.Done()
					break
				}
				//从指定位置开始读
				n, err := downloadFile.ReadAt(b, current*perG+total)
				if err == io.EOF {
					wg.Done()
					break
				}
				//从指定位置开始写
				copyFile.WriteAt(b, current*perG+total)
				storageFile.WriteAt([]byte(strconv.FormatInt(total, 10)+" "), current*16)
				total += int64(n)
				if total >= perG/10*int64(progressBar) {
					progressBar += 1
					p.Add(int(perG / 10))
				}

			}

		}(currentIndex)
	}
	wg.Wait()
	storageFile.Close()
	os.Remove(storageFilename)
	fmt.Println("下载完成")
}
