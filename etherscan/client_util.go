package etherscan

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
    "io"
    "log"
    "net/http"
    "os"
    "strings"
    "time"
)

type RequestEntry struct {
    Host      string    `json:"host"`
    Path      string    `json:"path"`
    Params    string    `json:"params"`
    Headers   string    `json:"headers"`
    Method    string    `json:"method"`
    Timestamp time.Time `json:"timestamp"`
    File      string    `json:"file"`
}

func doRequest(req *http.Request, replay bool, basePath string, removeHeader []string, removeQueryParam []string) ([]byte, error) {
    if !replay {
        return performHttpRequest(req)
    }

    host := strings.Replace(req.URL.Host, ".", "_", -1)
    path := req.URL.EscapedPath()
    params := getParams(req, removeQueryParam)
    headers := getHeaders(req, removeHeader)

    basePath = basePath + "/" + host
    basePathResponse := basePath + "/response"

    index := getIndex(basePath)
    log.Println("Index [ " + basePath + "/index.json ]")

    for _, request := range index {

        if request.Method == req.Method &&
            request.Host == host &&
            request.Path == path &&
            request.Params == params &&
            request.Headers == headers {

            name := basePathResponse + path + "/" + request.File

            if m, err := os.Stat(name); err != nil || m.IsDir() {
                continue
            } else {
                log.Println("Returning stored response", name)
                return os.ReadFile(name)
            }
        }
    }

    res, err := performHttpRequest(req)
    if err != nil {
        return nil, err
    }

    fileName := uuid.New().String() + ".json"
    saveFile(basePathResponse+"/"+path, "/"+fileName, res)

    requestEntry := RequestEntry{
        Host:      host,
        Path:      path,
        Params:    params,
        Headers:   headers,
        Method:    req.Method,
        Timestamp: time.Now(),
        File:      fileName,
    }

    index = append(index, requestEntry)

    indexBytes, err := json.Marshal(index)
    logError(err, "Marshalling updated index")

    saveFile(basePath, "/index.json", indexBytes)
    return res, nil
}

func getHeaders(req *http.Request, removeHeader []string) string {
    filteredHeaders := req.Header.Clone()
    for _, y := range removeHeader {
        filteredHeaders.Del(y)
    }
    formattedHeaders := fmt.Sprintf("%v", filteredHeaders)
    return formattedHeaders
}

func getParams(req *http.Request, removeQuery []string) string {
    copiedParams := req.URL.Query()
    for _, y := range removeQuery {
        copiedParams.Del(y)
    }
    params := fmt.Sprintf("%v", copiedParams)
    return params
}

func logError(err error, description string) {
    if err != nil {
        log.Println("ERROR: " + description + " [" + err.Error() + " ]")
    }
}

func getIndex(basePath string) []RequestEntry {
    // Create base path if not exist
    err := os.MkdirAll(basePath+"/response", 0755)
    logError(err, "Creating paths")
    if err != nil {
        return make([]RequestEntry, 0)
    }

    // Create index if not exist
    index := basePath + "/index.json"
    if _, err := os.Stat(index); os.IsNotExist(err) {
        newIndex, err := json.Marshal(make([]RequestEntry, 0))
        logError(err, "Marshalling new index")
        saveFile(basePath, "/index.json", newIndex)
    }

    // Return index from path
    var fileIndex []RequestEntry
    indexByteArray, err := os.ReadFile(index)
    unmarshalError := json.Unmarshal(indexByteArray, &fileIndex)
    logError(unmarshalError, "Unmarshalling index")

    return fileIndex
}

func performHttpRequest(req *http.Request) ([]byte, error) {
    httpClient := http.Client{
        Timeout: 30 * time.Second,
    }

    res, requestError := httpClient.Do(req)
    if requestError != nil {
        return nil, requestError
    }

    resBody, responseError := io.ReadAll(res.Body)
    if responseError != nil {
        return nil, responseError
    }

    return resBody, nil
}

func saveFile(path string, fileName string, byteArray []byte) {

    file := strings.Replace(path+"/"+fileName, "//", "/", -1)

    err := os.MkdirAll(path, 0755)
    logError(err, "Creating path "+path)
    if err != nil {
        return
    }

    log.Println("Writing file", file)
    err = os.WriteFile(file, byteArray, 0666)
    logError(err, "Saving file "+file)
}
