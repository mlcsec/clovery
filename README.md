# clovery
Cloud Discovery - check for open AWS, GCP, Alibaba, and Azure cloud services.  Check out the [demo](https://mlcsec.com/release-clovery/).

<br>

## Feedback
Just started learning Go, any feedback on usage/code optimisation would be appreciated!

<br>

## Info
* Cloud providers will RATE LIMIT you if usage too frequent/heavy or wordlist too big
* May have to configure max file descriptors with `ulimit -n <max>` depending on wordlist size
* 403 errors aid in further discovery, if company/keyword returns a lot of 403 errors - consider mutating the wordlist

<br>

## Installation
```
▶ go get github.com/mlcsec/clovery
```

<br>

## Options
```
▶ clovery -h
Usage of clovery:
  -ali
    	check alibaba OSS (-ali-reg required)
    	
  -ali-reg string
    	
    	Regions: (-ali option only)
    	===========================
    	oss-cn-hangzhou
    	oss-cn-shanghai
    	oss-cn-qingdao
    	oss-cn-beijing
    	oss-cn-zhangjiakou
    	oss-cn-huhehaote
    	oss-cn-shenzhen
    	oss-cn-chengdu
    	oss-cn-hongkong
    	oss-us-west-1
    	oss-us-east-1
    	oss-ap-southeast-1
    	oss-ap-southeast-2
    	oss-ap-southeast-3
    	oss-ap-southeast-5
    	oss-ap-northeast-1
    	oss-ap-south-1
    	oss-eu-central-1
    	oss-eu-west-1
    	oss-me-east-1
    	
  -all
    	run all checks (-aws, -gcp, -azb, -azd, -azw)
    	
  -aws
    	check s3 buckets
    	
  -azb
    	check azure blobs
    	
  -azd
    	check azure databases
    	
  -azu-reg string
    	
    	Regions: (-azv option only)
    	===========================
    	australiacentral
    	australiacentral2
    	australiaeast
    	australiasoutheast
    	brazilsouth
    	canadacentral
    	canadaeast
    	centralindia
    	centralus
    	eastasia
    	eastus
    	eastus2
    	francecentral
    	francesouth
    	japaneast
    	japanwest
    	koreacentral
    	koreasouth
    	northcentralus
    	northeurope
    	southafricanorth
    	southafricawest
    	southcentralus
    	southeastasia
    	southindia
    	uksouth
    	ukwest
    	westcentralus
    	westeurope
    	westindia
    	westus
    	westus2
    	
  -azv
    	check azure virtual machines (-azu-reg required)
    	
  -azw
    	check azure websites
    	
  -c int
    	concurrency level
    	 (default 20)
  -fe
    	show 403 forbidden errors
    	
  -gcp
    	check gcp buckets
    	
  -w string
    	wordlist
```

<br>

## Usage
* `createWordlists.sh` uses `sed` to append given keyword to front & back of the supplied wordlist
* `sed` patterns based on common bucket name formats
* supplied `common.txt` is a good starting point 
* can tailor any wordlist based on given keyword

<br>

## Credit
Tool inspired by [AWSBucketDump](https://github.com/jordanpotti/AWSBucketDump), [GCPBucketBrute](https://github.com/RhinoSecurityLabs/GCPBucketBrute), and [cloud_enum](https://github.com/initstring/cloud_enum).  Shoutout and thanks to them.
