package main

import (
    "fmt"
    "bufio"
    "os"
    "net/http"
    "net"
    "sync"
    "time"
)

func azureDb() {
    fmt.Println("\n[\033[36;1m-\033[0m] \033[36;1mAzure Databases: \033[0m")
    fmt.Println("====================================================")
    azdjobs := make(chan string)
    var wg sync.WaitGroup
    for i := 0; i < concurrency; i ++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for domain := range azdjobs {
                url := domain + ".database.windows.net"
                _, err := net.LookupHost(url)
                if err != nil {
                    continue
                }
                fmt.Println(url + " [\033[32mREGISTERED DNS\033[0m] \033[0m")
            }
        }()
    }
    sc := bufio.NewScanner(os.Stdin)
    for sc.Scan() {
      domain := sc.Text()
      azdjobs <- domain
    }
    close(azdjobs)
    wg.Wait()
    fmt.Println("====================================================")
}
