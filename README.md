# clovery
Cloud Discovery - brute force public AWS, GCP, Alibaba, and Azure cloud services.  

<a href="https://asciinema.org/a/381917"><img src="https://asciinema.org/a/381917.png" alt="asciicast" /></a>
<br>

## Info
* Cloud providers will rate limit you if usage too frequent/heavy or wordlist too big (no output)
* `403` errors aid in further discovery - consider mutating wordlist

<br>

## Installation
```
go get github.com/mlcsec/clovery
```
<br>

## Help
```
$ clovery -h
Usage:
  cat wordlist.txt | clovery -aws
  cat wordlist.txt | clovery -ali -ali-reg <region>

Options:
  -ali            alibaba OSS (-ali-reg required)
  -ali-reg
                  Regions: (-ali option only)
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
        
  -aws            aws s3 buckets
  -azb            azure blob storage
  -azd            azure databases 
  -azv            azure virtual machines (-azv-reg required)
  -azu-reg
                  Regions: (-azv option only)
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
        
  -azw            azure websites
  -c              concurrency level (default 20)
  -fe             show 403 forbidden errors
  -gcp            gcp storage
  -t              timeout milliseconds (default 10000)
```
<br>

## Wordlists
* `createWordlists.sh` uses `sed` to append given keyword to front & back of the supplied wordlist
* Tailor any wordlist based on given keyword
