package main

import (
    "fmt"
    "os"
    "flag"
)

var concurrency int
var to int
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
var alReg []string
var azReg []string

func Find(slice []string, val string) (int, bool) {
    for i, item := range slice {
        if item == val {
            return i, true
        }
    }
    return -1, false
}

func init() {
    flag.Usage = func() {
        f := "Usage:\n"
        f += "  cat wordlist.txt | clovery -aws\n"
        f += "  cat wordlist.txt | clovery -ali -ali-reg <region>\n\n"
        f += "Options:\n"
        f += "  -ali            alibaba OSS (-ali-reg required)\n"
        f += "  -ali-reg\n"
        f += `                  Regions: (-ali option only)
    	          =====================================
                  oss-cn-hangzhou       oss-cn-shanghai
                  oss-cn-qingdao        oss-cn-beijing
                  oss-cn-zhangjiakou    oss-cn-huhehaote
                  oss-cn-shenzhen       oss-cn-chengdu
                  oss-cn-hongkong       oss-us-west-1
                  oss-us-east-1         oss-ap-southeast-1
                  oss-ap-southeast-2    oss-ap-southeast-3
                  oss-ap-southeast-5    oss-ap-northeast-1
                  oss-ap-south-1        oss-eu-central-1
                  oss-eu-west-1         oss-me-east-1
        `
        f += "\n  -aws            aws s3 buckets\n"
        f += "  -azb            azure blob storage\n"
        f += "  -azd            azure databases \n"
        f += "  -azv            azure virtual machines (-azv-reg required)\n"
        f += "  -azu-reg\n"
        f += `                  Regions: (-azv option only)
    	          =====================================
                  australiacentral      australiacentral2
                  australiaeast         australiasoutheast
                  brazilsouth           canadacentral
                  canadaeast            centralindia
                  centralus             eastasia
                  eastus                eastus2
                  francecentral         francesouth
                  japaneast             japanwest
                  koreacentral          koreasouth
                  northcentralus        northeurope
                  southafricanorth      southafricawest
                  southcentralus        southeastasia
                  southindia            uksouth
                  ukwest                westcentralus
                  westeurope            westindia
                  westus                westus2                    
        `
        f += "\n  -azw            azure websites\n"
        f += "  -c              concurrency level (default 20)\n"
        f += "  -fe             show 403 forbidden errors\n"
        f += "  -gcp            gcp storage\n"
        f += "  -t              timeout milliseconds (default 10000)\n"
        fmt.Fprintf(os.Stderr, f)
    }
}

func main() {
    flag.IntVar(&concurrency, "c", 20, "concurrency level\n")
    flag.IntVar(&to, "t", 10000, "timeout (milliseconds)\n")
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
    flag.Parse()
    if aws {
        awsCheck()
    }
    if gcp {
        gcpCheck()
    }
    if ali {
        alReg = []string{"oss-cn-hangzhou","oss-cn-shanghai","oss-cn-qingdao","oss-cn-beijing","oss-cn-zhangjiakou","oss-cn-huhehaote","oss-cn-shenzhen","oss-cn-chengdu","oss-cn-hongkong","oss-us-west-1","oss-us-east-1","oss-ap-southeast-1","oss-ap-southeast-2","oss-ap-southeast-3","oss-ap-southeast-5","oss-ap-northeast-1","oss-ap-south-1","oss-eu-central-1","oss-eu-west-1","oss-me-east-1"}
        if alibabaRegion == "" {
            fmt.Println("[\033[31;1m!\033[0m] \033[31;1mNo region supplied\033[0m")
            fmt.Println("[\033[31;1m!\033[0m] \033[31;1mcat <wordlist> | clovery -ali -ali-reg <region>\033[0m")
            fmt.Println("oss-cn-hangzhou\noss-cn-shanghai\noss-cn-qingdao\noss-cn-beijing\noss-cn-zhangjiakou\noss-cn-huhehaote\noss-cn-shenzhen\noss-cn-chengdu\noss-cn-hongkong\noss-us-west-1\noss-us-east-1\noss-ap-southeast-1\noss-ap-southeast-2\noss-ap-southeast-3\noss-ap-southeast-5\noss-ap-northeast-1\noss-ap-south-1\noss-eu-central-1\noss-eu-west-1\noss-me-east-1")
            os.Exit(0)
        }
        _, found := Find(alReg, alibabaRegion)
        if !found {
            fmt.Println("\n[\033[31;1m!\033[0m] \033[31;1m" + alibabaRegion + " not valid Alibaba region, see below:\033[0m")
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
            fmt.Println("[\033[31;1m!\033[0m] \033[31;1mNo region supplied\033[0m")
            fmt.Println("[\033[31;1m!\033[0m] \033[31;1mcat <wordlist> | clovery -azv -azu-reg <region>\033[0m")
            fmt.Println("australiacentral\naustraliacentral2\naustraliaeast\naustraliasoutheast\nbrazilsouth\ncanadacentral\ncanadaeast\ncentralindia\ncentralus\neastasia\neastus\neastus2\nfrancecentral\nfrancesouth\njapaneast\njapanwest\nkoreacentral\nkoreasouth\nnorthcentralus\nnortheurope\nsouthafricanorth\nsouthafricawest\nsouthcentralus\nsoutheastasia\nsouthindia\nuksouth\nukwest\nwestcentralus\nwesteurope\nwestindia\nwestus\nwestus2")
            os.Exit(0)
        }
        _, found := Find(azReg, azureRegion)
        if !found {
            fmt.Println("\n[\033[31;1m!\033[0m] \033[31;1m" + azureRegion + " not valid Azure region, see below:\033[0m")
            fmt.Println("australiacentral\naustraliacentral2\naustraliaeast\naustraliasoutheast\nbrazilsouth\ncanadacentral\ncanadaeast\ncentralindia\ncentralus\neastasia\neastus\neastus2\nfrancecentral\nfrancesouth\njapaneast\njapanwest\nkoreacentral\nkoreasouth\nnorthcentralus\nnortheurope\nsouthafricanorth\nsouthafricawest\nsouthcentralus\nsoutheastasia\nsouthindia\nuksouth\nukwest\nwestcentralus\nwesteurope\nwestindia\nwestus\nwestus2")
            os.Exit(0)
        }
        azureVm()
    }
}