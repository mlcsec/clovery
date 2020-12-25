package cloud

func awsCheck() {
    fmt.Println("\n[\033[36;1m-\033[0m] \033[36;1mAWS Buckets: \033[0m")
    fmt.Println("====================================================")
    timeout := time.Duration(to * 1000000)
    var tr = &http.Transport{
		MaxIdleConns:      30,
		IdleConnTimeout:   time.Second,
		DisableKeepAlives: true,
		DialContext: (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: time.Second,
		}).DialContext,
	}
	client := &http.Client{
		Transport:     tr,
		Timeout:       timeout,
	}
    awsjobs := make(chan string)
    var wg sync.WaitGroup
    for i := 0; i < concurrency; i ++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for domain := range awsjobs {
                url := "http://" + domain + ".s3.amazonaws.com"
                req, err := client.Get(url)
                if err != nil {
                    continue
                }
                defer req.Body.Close()
                if req.StatusCode == 200 {
                    fmt.Println(url + " [\033[32m200\033[0m] \033[0m")
                }
		        if fe {
		            if req.StatusCode == 403 {
			            fmt.Println(url + " [\033[31m403\033[0m] \033[0m")
		           }
		       }
            }
        }()
    }
    sc := bufio.NewScanner(os.Stdin)
    for sc.Scan() {
      domain := sc.Text()
      awsjobs <- domain
    }
    close(awsjobs)
    wg.Wait()
    fmt.Println("====================================================")
}
