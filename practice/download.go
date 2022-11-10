package practice

import (
	"encoding/base64"
	"fmt"
	"giao/util"
	"github.com/pkg/errors"
	"github.com/prometheus/common/log"
	"github.com/qianlnk/pgbar"
	"github.com/schollz/progressbar/v3"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
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

// 直接下载文件
func directDownload(url string, header map[string]string, writer io.Writer) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	bar := progressbar.DefaultBytes(resp.ContentLength, "downloading"+fmt.Sprintf("%v", header))
	defer resp.Body.Close()
	_, err = io.Copy(io.MultiWriter(writer, bar), resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func DownWithPiece(url string, filePath string, pieceNum int64, pieceSize int64, mutex *sync.Mutex) error {
	startRange := pieceNum * pieceSize
	endRange := (pieceNum+1)*pieceSize - 1
	byteRanges := fmt.Sprintf("bytes=%v-%v", startRange, endRange)
	fmt.Println("开始下载", pieceNum, byteRanges)
	var headers = map[string]string{}
	headers["Range"] = byteRanges
	fw, err := os.OpenFile(filePath, os.O_WRONLY, 0660)
	if err != nil {
		return err
	}
	defer fw.Close()
	// 打开文件偏移位置写入
	if _, err = fw.Seek(startRange, 0); err != nil {
		return err
	}
	if err = directDownload(url, headers, fw); err != nil {
		return err
	}
	return nil
}

func DownWithMultiThreadV1(url string, filePath string, threads int, size int64) error {
	pieceSize := size / int64(threads)
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	// 创建目标文件
	f, _ := os.Create(filePath)
	defer f.Close()
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func(pieceNum int64, pieceSize int64) {
			defer wg.Done()
			err := DownWithPiece(url, filePath, pieceNum, pieceSize, &mutex)
			if err != nil {
				fmt.Println("goroutine err", err)
			}
		}(int64(i), pieceSize)
	}
	wg.Wait()

	fmt.Println("下载结束")
	return nil
}

// GetOriginFileSize 获取文件大小 若无法分片则抛出错误
func GetOriginFileSize(url string) (string, int64, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", 0, err
	}
	req.Header.Set("Range", "bytes=0-1")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	contentRange := resp.Header.Get("Content-Range")
	if contentRange == "" {
		return "", 0, errors.New("该文件不支持分片下载")
	}
	// 获取前缀
	suffix := GetSuffix(resp.Header.Get("Content-Type"))
	s := strings.Split(contentRange, "/")[1]
	fileSize, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return "", 0, err
	}
	return suffix, fileSize, nil
}

var contentTypes = map[string]string{
	"image/gif":                    "gif",
	"image/jpeg":                   "jpg",
	"application/x-img":            "img",
	"image/png":                    "png",
	"image/webp":                   "png",
	"application/json":             "json",
	"application/pdf":              "pdf",
	"application/msword":           "word",
	"application/octet-stream":     "rar",
	"application/x-zip-compressed": "zip",
	"application/x-msdownLoad":     "exe",
	"video/mpeg4":                  "mp4",
	"video/avi":                    "avi",
	"audio/mp3":                    "mp3",
	"text/css":                     "css",
	"application/x-javascript":     "js",
	"application/vnd.android.package-archive": "apk",
}

func GetSuffix(t string) string {
	return contentTypes[t]
}

func TestMultiThread() {
	urlArr := [...]string{
		"https://www.ucg.ac.me/skladiste/blog_44233/objava_64433/fajlovi/Computer%20Networking%20_%20A%20Top%20Down%20Approach,%207th,%20converted.pdf",
		"https://dldir1.qq.com/qqfile/qq/PCQQ9.6.2/QQ9.6.2.28756.exe",
		"https://img1.baidu.com/it/u=3076513868,1671318234&fm=253&fmt=auto&app=138&f=PNG?w=1025&h=449",
	}
	url := urlArr[0]
	suffix, size, err := GetOriginFileSize(url)
	log.Info("size", size)
	if err != nil {
		fmt.Println("getOriginFileSize", err)
		return
	}
	baseFileName := base64.StdEncoding.EncodeToString([]byte(url[len(url)-10:]))
	dirName := "./download/"
	filePath := dirName + baseFileName + "." + suffix

	// 判断文件是否存在
	if fileExists, _ := util.IsExists(filePath); !fileExists {
		_ = os.Mkdir(dirName, 0644)
		f, _ := os.OpenFile(filePath, os.O_CREATE, os.ModePerm)
		f.Close()
	}
	err = DownWithMultiThreadV1(url, filePath, 5, size)
	if err != nil {
		fmt.Println("DownloadError", err)
		return
	}
}
