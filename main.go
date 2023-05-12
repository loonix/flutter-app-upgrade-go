package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "path/filepath"
    "strconv"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.Use(corsMiddleware())

    r.GET("/appcast/:version", func(c *gin.Context) {
        version := c.Param("version")
        filePath := filepath.Join("files", version+".json")
        data, err := ioutil.ReadFile(filePath)
        if err != nil {
            c.AbortWithStatus(http.StatusInternalServerError)
            return
        }
        var result interface{}
        err = json.Unmarshal(data, &result)
        if err != nil {
            c.AbortWithStatus(http.StatusInternalServerError)
            return
        }
        c.JSON(http.StatusOK, result)
    })

    r.POST("/upload", func(c *gin.Context) {
        files := c.Request.MultipartForm.File["file"]
        if len(files) == 0 {
            c.AbortWithStatus(http.StatusBadRequest)
            return
        }
        for _, file := range files {
            err := c.SaveUploadedFile(file, filepath.Join("files", file.Filename))
            if err != nil {
                c.AbortWithStatus(http.StatusInternalServerError)
                return
            }
        }
        c.String(http.StatusOK, "Files uploaded")
    })

    r.GET("/download/:filename", func(c *gin.Context) {
        filename := c.Param("filename")
        file, err := os.Open(filepath.Join("files", filename))
        if err != nil {
            c.AbortWithStatus(http.StatusInternalServerError)
            return
        }
        defer file.Close()
        fileInfo, err := file.Stat()
        if err != nil {
            c.AbortWithStatus(http.StatusInternalServerError)
            return
        }
        c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
        c.Header("Content-Type", "application/octet-stream")
        c.Header("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))
        c.File(filepath.Join("files", filename))
    })

    if err := r.Run(":8000"); err != nil {
        panic(err)
    }
}

func corsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
        c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        c.Next()
    }
}
