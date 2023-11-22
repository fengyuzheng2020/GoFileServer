package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// Config holds the configuration for the server
type Config struct {
	UploadDir string
	Port      int
}

var config Config

func main() {
	flag.StringVar(&config.UploadDir, "upload-dir", "./uploads/", "Specify the directory to store uploaded files")
	flag.IntVar(&config.Port, "port", 8080, "Specify the port number for the server")
	flag.Parse()

	http.HandleFunc("/", uploadHandler)
	http.HandleFunc("/download/", downloadHandler)
	address := fmt.Sprintf(":%d", config.Port)
	fmt.Printf("Server listening on http://localhost%s\n", address)
	http.ListenAndServe(address, nil)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// 解析表单
		err := r.ParseMultipartForm(10 << 20) // 限制上传文件大小为10 MB
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 获取文件
		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Error retrieving the file:", err)
			return
		}
		defer file.Close()

		// 创建上传目录
		if _, err := os.Stat(config.UploadDir); os.IsNotExist(err) {
			err = os.Mkdir(config.UploadDir, os.ModeDir)
			if err != nil {
				fmt.Println("Error creating the upload directory:", err)
				return
			}
		}

		// 创建上传文件
		filePath := filepath.Join(config.UploadDir, handler.Filename)
		fmt.Println("File:", filePath)
		dst, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Error creating the file:", err)
			return
		}
		defer dst.Close()

		// 将文件内容拷贝到目标文件
		_, err = io.Copy(dst, file)
		if err != nil {
			fmt.Println("Error copying file:", err)
			return
		}

		fmt.Fprintf(w, "File %s uploaded successfully to %s!", handler.Filename, config.UploadDir)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method not allowed. Please use POST.")
	}
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fileName := r.URL.Path[len("/download/"):]
		filePath := filepath.Join(config.UploadDir, fileName)

		file, err := os.Open(filePath)
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		defer file.Close()

		// 获取文件信息
		fileInfo, err := file.Stat()
		if err != nil {
			http.Error(w, "Failed to get file information", http.StatusInternalServerError)
			return
		}
		// 设置响应头，告诉浏览器这是一个文件下载
		w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Length", fmt.Sprint(fileInfo.Size()))

		// 将文件内容拷贝到响应体
		_, err = io.Copy(w, file)
		if err != nil {
			fmt.Println("Error copying file:", err)
			return
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method not allowed. Please use GET.")
	}
}
