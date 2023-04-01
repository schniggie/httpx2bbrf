package main

import (
        "encoding/json"
        "fmt"
        "os"
        "os/exec"
        "strings"
)

func main() {
        decoder := json.NewDecoder(os.Stdin)

        for {
                json := make(map[string]interface{})

                err := decoder.Decode(&json)

                if err != nil {
                    fmt.Fprintln(os.Stderr, "error:", err)
                    break
                }

                urlUrl, ok := json["url"].(string)
                if !ok {
                    urlUrl = "<nil>"
                }
                urlTitle, _ := json["title"].(string)
                urlWebserver, _ := json["webserver"].(string)
                urlContentType, _ := json["content_type"].(string)
                urlContentLength, _ := json["content_length"].(float64)
                urlStatusCode, _ := json["status_code"].(float64)
                urlServerResponse, _ := json["body"].(string)

                if urlUrl != "<nil>" {
                        // Remove the port part of the URL if it's the default http or https port
                        if strings.HasSuffix(urlUrl, ":80") {
                            urlUrl = strings.TrimSuffix(urlUrl, ":80")
                        } else if strings.HasSuffix(urlUrl, ":443") {
                            urlUrl = strings.TrimSuffix(urlUrl, ":443")
                        }

                        urlUrlString := fmt.Sprintf("%v", urlUrl)
                        urlTitleString := fmt.Sprintf("%v", urlTitle)
                        urlWebserverString := fmt.Sprintf("%v", urlWebserver)
                        urlContentTypeString := fmt.Sprintf("%v", urlContentType)
                        urlContentLengthString := fmt.Sprintf("%v", int64(urlContentLength))
                        urlStatusCodeString := fmt.Sprintf("%v", int64(urlStatusCode))
                        urlServerResponseString := fmt.Sprintf("%v", urlServerResponse)

                        fullCommand := "bbrf url add '" + urlUrlString + " " + urlStatusCodeString + " " + urlContentLengthString + "' -t 'title:" + urlTitleString + "'" + " -t 'webserver:" + urlWebserverString + "'" + " -t 'contenttype:" + urlContentTypeString + "'" + " -t 'contentlength:" + urlContentLengthString + "'" + " -t 'statuscode:" + urlStatusCodeString + "'" + " -t 'serverresponse:" + urlServerResponseString + "'" + " -p @INFER"

                        // Append any command line arguments to the fullCommand
                        for _, arg := range os.Args[1:] {
                            fullCommand += " " + arg
                        }

                        out, _ :=exec.Command("sh","-c",fullCommand).Output()
                        fmt.Printf("%s", out)

                } else {
                        fmt.Println("url is <nil>")
                        break
                }
        }
}
