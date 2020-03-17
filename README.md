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
| Flag | Type |Description |
|:---  |:--- |:---         |
|-ali  |bool|check alibaba OSS|
|-ali-reg|string|Regions: (-ali option only)<br>===========================<br>oss-cn-hangzhou<br>oss-cn-shanghai<br>oss-cn-qingdao<br>oss-cn-beijing<br>oss-cn-zhangjiakou<br>oss-cn-huhehaote<br>oss-cn-shenzhen<br>oss-cn-chengdu<br>oss-cn-hongkong<br>oss-us-west-1<br>oss-us-east-1<br>oss-ap-southeast-1<br>oss-ap-southeast-2<br>oss-ap-southeast-3<br>oss-ap-southeast-5<br>oss-ap-northeast-1<br>oss-ap-south-1<br>oss-eu-central-1<br>oss-eu-west-1<br>oss-me-east-1|
|-aws  |bool|check s3 buckets|
|-azb  |bool|check azure blobs|
|-azd  |bool|check azure databases|
|-azu-reg|string|Regions: (-azv option only)<br>===========================<br>australiacentral<br>australiacentral2<br>australiaeast<br>australiasoutheast<br>brazilsouth<br>canadacentral<br>canadaeast<br>centralindia<br>centralus<br>eastasia<br>eastus<br>eastus2<br>francecentral<br>francesouth<br>japaneast<br>japanwest<br>koreacentral<br>koreasouth<br>northcentralus<br>northeurope<br>southafricanorth<br>southafricawest<br>southcentralus<br>southeastasia<br>southindia<br>uksouth<br>ukwest<br>westcentralus<br>westeurope<br>westindia<br>westus<br>westus2|
|-azw  |bool|check azure websites|
|-c    |int|concurrency level (default 20)|
|-fe   |bool|show 403 forbidden errors|
|-gcp  |bool|check gcp buckets|
|-w    |string|wordlist|


<br>

## Usage
* `createWordlists.sh` uses `sed` to append given keyword to front & back of the supplied wordlist
* `sed` patterns based on common bucket name formats
* supplied `common.txt` is a good starting point 
* can tailor any wordlist based on given keyword

<br>

## Credit
Tool inspired by [AWSBucketDump](https://github.com/jordanpotti/AWSBucketDump), [GCPBucketBrute](https://github.com/RhinoSecurityLabs/GCPBucketBrute), and [cloud_enum](https://github.com/initstring/cloud_enum).  Shoutout and thanks to them.
