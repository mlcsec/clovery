package cloud

func azureVm() {
    fmt.Println("\n[\033[36;1m-\033[0m] \033[36;1mAzure VMs: \033[0m")
    fmt.Println("====================================================")
    azvjobs := make(chan string)
    var wg sync.WaitGroup
    for i := 0; i < concurrency; i ++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for domain := range azvjobs {
                url := domain + "." + azureRegion + ".cloudapp.azure.com"
                _, err := net.LookupHost(url)
                if err != nil {
                    continue
                }
                fmt.Println(url + " [\033[32mREGISTERED VM\033[0m] \033[0m")
            }
        }()
    }
    sc := bufio.NewScanner(os.Stdin)
    for sc.Scan() {
      domain := sc.Text()
      azvjobs <- domain
    }
    close(azvjobs)
    wg.Wait()
    fmt.Println("====================================================")
}
