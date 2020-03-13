package main

import (
    "fmt"
    "bufio"
    "os"
    "net/http"
    "net"
    "sync"
    "flag"
    "log"
)

var concurrency int
var aws bool
var gcp bool
var azb bool
var azw bool
var azd bool
var azv bool
var ali bool
var fe bool
var azureRegion string
var alibabaRegion string
var w string
var alReg []string
var azReg []string

func main() {
    flag.IntVar(&concurrency, "c", 20, "concurrency level\n")
    flag.BoolVar(&aws, "aws", false, "check s3 buckets\n")
    flag.BoolVar(&gcp, "gcp", false, "check gcp buckets\n")
    flag.BoolVar(&azb, "azb", false, "check azure blobs\n")
    flag.BoolVar(&azw, "azw", false, "check azure websites\n")
    flag.BoolVar(&azd, "azd", false, "check azure databases\n")
    flag.BoolVar(&azv, "azv", false, "check azure virtual machines (-azu-reg required)\n")
    flag.BoolVar(&ali, "ali", false, "check alibaba OSS (-ali-reg required)\n")
    flag.BoolVar(&fe, "fe", false, "show 403 forbidden errors\n")
    flag.StringVar(&alibabaRegion, "ali-reg", "", "\nRegions: (-ali option only)\n===========================\noss-cn-hangzhou\noss-cn-shanghai\noss-cn-qingdao\noss-cn-beijing\noss-cn-zhangjiakou\noss-cn-huhehaote\noss-cn-shenzhen\noss-cn-chengdu\noss-cn-hongkong\noss-us-west-1\noss-us-east-1\noss-ap-southeast-1\noss-ap-southeast-2\noss-ap-southeast-3\noss-ap-southeast-5\noss-ap-northeast-1\noss-ap-south-1\noss-eu-central-1\noss-eu-west-1\noss-me-east-1\n")
    flag.StringVar(&azureRegion, "azu-reg", "", "\nRegions: (-azv option only)\n===========================\naustraliacentral\naustraliacentral2\naustraliaeast\naustraliasoutheast\nbrazilsouth\ncanadacentral\ncanadaeast\ncentralindia\ncentralus\neastasia\neastus\neastus2\nfrancecentral\nfrancesouth\njapaneast\njapanwest\nkoreacentral\nkoreasouth\nnorthcentralus\nnortheurope\nsouthafricanorth\nsouthafricawest\nsouthcentralus\nsoutheastasia\nsouthindia\nuksouth\nukwest\nwestcentralus\nwesteurope\nwestindia\nwestus\nwestus2\n")
    flag.StringVar(&w, "w", "", "wordlist")
    flag.Parse()
    title :=
`
          __
    _____/ /___ _   _____  _______  __
   / ___/ / __ \ | / / _ \/ ___/ / / /
  / /__/ / /_/ / |/ /  __/ /  / /_/ /
  \___/_/\____/|___/\___/_/   \__, /
                             /____/

                        ~ Cloud Discovery ~
`
    fmt.Println("\033[36;1m"+title+"\033[0m")
    if aws {
        awsCheck()
    }
    if gcp {
        gcpCheck()
    }
    if ali {
        alReg = []string{"oss-cn-hangzhou","oss-cn-shanghai","oss-cn-qingdao","oss-cn-beijing","oss-cn-zhangjiakou","oss-cn-huhehaote","oss-cn-shenzhen","oss-cn-chengdu","oss-cn-hongkong","oss-us-west-1","oss-us-east-1","oss-ap-southeast-1","oss-ap-southeast-2","oss-ap-southeast-3","oss-ap-southeast-5","oss-ap-northeast-1","oss-ap-south-1","oss-eu-central-1","oss-eu-west-1","oss-me-east-1"}
        if alibabaRegion == "" {
            fmt.Println("\n[\033[31;1m-\033[0m] \033[31;1mNo region supplied\033[0m")
            os.Exit(0)
        }
        _, found := Find(alReg, alibabaRegion)
        if !found {
            fmt.Println("\n[\033[31;1m-\033[0m] \033[31;1m" + alibabaRegion + " not valid Alibaba region, see below:\033[0m")
            fmt.Println("oss-cn-hangzhou\noss-cn-shanghai\noss-cn-qingdao\noss-cn-beijing\noss-cn-zhangjiakou\noss-cn-huhehaote\noss-cn-shenzhen\noss-cn-chengdu\noss-cn-hongkong\noss-us-west-1\noss-us-east-1\noss-ap-southeast-1\noss-ap-southeast-2\noss-ap-southeast-3\noss-ap-southeast-5\noss-ap-northeast-1\noss-ap-south-1\noss-eu-central-1\noss-eu-west-1\noss-me-east-1")
            os.Exit(0)
        }
        alibabaOss()
    }
    if azb {
        azureBlob()
    }
    if azd {
        azureDb()
    }
    if azw {
        azureWeb()
    }
    if azv {
        azReg = []string{"australiacentral","australiacentral2","australiaeast","australiasoutheast","brazilsouth","canadacentral","canadaeast","centralindia","centralus","eastasia","eastus","eastus2","francecentral","francesouth","japaneast","japanwest","koreacentral","koreasouth","northcentralus","northeurope","southafricanorth","southafricawest","southcentralus","southeastasia","southindia","uksouth","ukwest","westcentralus","westeurope","westindia","westus","westus2"}
        if azureRegion == "" {
            fmt.Println("\n[\033[31;1m-\033[0m] \033[31;1mNo region supplied\033[0m")
            os.Exit(0)
        }
        _, found := Find(azReg, azureRegion)
        if !found {
            fmt.Println("\n[\033[31;1m-\033[0m] \033[31;1m" + azureRegion + " not valid Azure region, see below:\033[0m")
            fmt.Println("australiacentral\naustraliacentral2\naustraliaeast\naustraliasoutheast\nbrazilsouth\ncanadacentral\ncanadaeast\ncentralindia\ncentralus\neastasia\neastus\neastus2\nfrancecentral\nfrancesouth\njapaneast\njapanwest\nkoreacentral\nkoreasouth\nnorthcentralus\nnortheurope\nsouthafricanorth\nsouthafricawest\nsouthcentralus\nsoutheastasia\nsouthindia\nuksouth\nukwest\nwestcentralus\nwesteurope\nwestindia\nwestus\nwestus2")
            os.Exit(0)
        }
        azureVm()
    }
}

func Find(slice []string, val string) (int, bool) {
    for i, item := range slice {
        if item == val {
            return i, true
        }
    }
    return -1, false
}

func awsCheck() {
    fmt.Println("\n[\033[36;1m-\033[0m] \033[36;1mAWS Buckets: \033[0m")
    fmt.Println("====================================================")
    awsjobs := make(chan string)
    var wg sync.WaitGroup
    for i := 0; i < concurrency; i ++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for domain := range awsjobs {
                url := "http://" + domain + ".s3.amazonaws.com"
                req, err := http.Get(url)
                if err != nil {
                    continue
                }
                if req.StatusCode == 200 {
                    fmt.Println("[\033[32m200\033[0m] \033[0m" + url + "\033[0m")
                }
		        if fe {
		            if req.StatusCode == 403 {
			            fmt.Println("[\033[31m403\033[0m] \033[0m" + url + "\033[0m")
		           }
		       }
                defer req.Body.Close()
            }
        }()
    }
    file, err := os.Open(w)
    if err != nil {
        log.Fatalf("failed opening file: %s", err)
    }
    sc := bufio.NewScanner(file)
    sc.Split(bufio.ScanLines)
    for sc.Scan() {
      domain := sc.Text()
      awsjobs <- domain
    }
    file.Close()
    close(awsjobs)
    wg.Wait()
    fmt.Println("====================================================")
}

func gcpCheck() {
    fmt.Println("\n[\033[36;1m-\033[0m] \033[36;1mGCP Buckets: \033[0m")
    fmt.Println("====================================================")
    gcpjobs := make(chan string)
    var wg sync.WaitGroup
    for i := 0; i < concurrency; i ++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for domain := range gcpjobs {
                url := "http://storage.googleapis.com/" + domain
                req, err := http.Get(url)
                if err != nil {
                    continue
                }
                if req.StatusCode == 200 {
                    fmt.Println("[\033[32m200\033[0m] \033[0m" + url + "\033[0m")
                }
		        if fe {
		            if req.StatusCode == 403 {
			            fmt.Println("[\033[31m403\033[0m] \033[0m" + url + "\033[0m")
		           }
		        }
                defer req.Body.Close()
            }
        }()
    }
    file, err := os.Open(w)
    if err != nil {
        log.Fatalf("failed opening file: %s", err)
    }
    sc := bufio.NewScanner(file)
    sc.Split(bufio.ScanLines)
    for sc.Scan() {
      domain := sc.Text()
      gcpjobs <- domain
    }
    file.Close()
    close(gcpjobs)
    wg.Wait()
    fmt.Println("====================================================")
}

func alibabaOss() {
    fmt.Println("\n[\033[36;1m-\033[0m] \033[36;1mAlibaba OSS: \033[0m")
    fmt.Println("====================================================")
    alijobs := make(chan string)
    var wg sync.WaitGroup
    for i := 0; i < concurrency; i ++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for domain := range alijobs {
                url := "http://" + domain + "." + alibabaRegion + ".aliyuncs.com"
                req, err := http.Head(url)
                if err != nil {
                    continue
                }
                if req.StatusCode == 200 {
                    fmt.Println("[\033[32m200\033[0m] \033[0m" + url + "\033[0m")
                }
		        if fe {
		            if req.StatusCode == 403 {
			            fmt.Println("[\033[31m403\033[0m] \033[0m" + url + "\033[0m")
		           }
		        }
                defer req.Body.Close()
            }
        }()
    }
    file, err := os.Open(w)
    if err != nil {
        log.Fatalf("failed opening file: %s", err)
    }
    sc := bufio.NewScanner(file)
    sc.Split(bufio.ScanLines)
    for sc.Scan() {
      domain := sc.Text()
      alijobs <- domain
    }
    file.Close()
    close(alijobs)
    wg.Wait()
    fmt.Println("====================================================")
}

func azureBlob() {
    fmt.Println("\n[\033[36;1m-\033[0m] \033[36;1mAzure Blobs: \033[0m")
    fmt.Println("====================================================")
    azujobs := make(chan string)
    var wg sync.WaitGroup
    for i := 0; i < concurrency; i ++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for domain := range azujobs {
                url := "http://" + domain + ".blob.core.windows.net"
                req, err := http.Get(url)
                if err != nil {
                    continue
                }
                if req.StatusCode == 200 {
                    fmt.Println("[\033[32m200\033[0m] \033[0m" + url + "\033[0m")
                }
		        if fe {
		            if req.StatusCode == 403 {
			            fmt.Println("[\033[31m403\033[0m] \033[0m" + url + "\033[0m")
		           }
		        }
                defer req.Body.Close()
            }
        }()
    }
    file, err := os.Open(w)
    if err != nil {
        log.Fatalf("failed opening file: %s", err)
    }
    sc := bufio.NewScanner(file)
    sc.Split(bufio.ScanLines)
    for sc.Scan() {
      domain := sc.Text()
      azujobs <- domain
    }
    file.Close()
    close(azujobs)
    wg.Wait()
    fmt.Println("====================================================")
}

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
                fmt.Println("[\033[32mRegistered DNS\033[0m] \033[0m" + url + "\033[0m")
            }
        }()
    }
    file, err := os.Open(w)
    if err != nil {
        log.Fatalf("failed opening file: %s", err)
    }
    sc := bufio.NewScanner(file)
    sc.Split(bufio.ScanLines)
    for sc.Scan() {
      domain := sc.Text()
      azdjobs <- domain
    }
    file.Close()
    close(azdjobs)
    wg.Wait()
    fmt.Println("====================================================")
}

func azureWeb() {
    fmt.Println("\n[\033[36;1m-\033[0m] \033[36;1mAzure Websites: \033[0m")
    fmt.Println("====================================================")
    azwjobs := make(chan string)
    var wg sync.WaitGroup
    for i := 0; i < concurrency; i ++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for domain := range azwjobs {
                url := "http://" + domain + ".azurewebsites.net"
                req, err := http.Head(url)
                if err != nil {
                    continue
                }
                if req.StatusCode == 200 {
                    fmt.Println("[\033[32m200\033[0m] \033[0m" + url + "\033[0m")
                }
		        if fe {
		            if req.StatusCode == 403 {
			            fmt.Println("[\033[31m403\033[0m] \033[0m" + url + "\033[0m")
		           }
		        }
                defer req.Body.Close()
            }
        }()
    }
    file, err := os.Open(w)
    if err != nil {
        log.Fatalf("failed opening file: %s", err)
    }
    sc := bufio.NewScanner(file)
    sc.Split(bufio.ScanLines)
    for sc.Scan() {
      domain := sc.Text()
      azwjobs <- domain
    }
    file.Close()
    close(azwjobs)
    wg.Wait()
    fmt.Println("====================================================")
}

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
                fmt.Println("[\033[32mRegistered VM\033[0m] \033[0m" + url + "\033[0m")
            }
        }()
    }
    file, err := os.Open(w)
    if err != nil {
        log.Fatalf("failed opening file: %s", err)
    }
    sc := bufio.NewScanner(file)
    sc.Split(bufio.ScanLines)
    for sc.Scan() {
      domain := sc.Text()
      azvjobs <- domain
    }
    file.Close()
    close(azvjobs)
    wg.Wait()
    fmt.Println("====================================================")
}
