package oss

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	uuid "github.com/satori/go.uuid"
	"github.com/wujunyi792/flamego-quick-template/config"
	"github.com/wujunyi792/flamego-quick-template/internal/core/logx"
	"hash"
	"io"
	"path"
	"time"
)

var client *oss.Client
var bucket *oss.Bucket
var log = logx.NameSpace("oss")

func InitOSS() {
	// 创建OSSClient实例。
	var err error
	conf := &config.GetConfig().OSS
	client, err = oss.New(conf.EndPoint, conf.AccessKeyId, conf.AccessKeySecret)
	// 获取存储空间。
	if err != nil {
		log.Fatalln(err)
	}
	bucket, err = client.Bucket(conf.BucketName)
	if err != nil {
		log.Fatalln("阿里云OSS连接失败: ", err)
	}
}

func UploadFileToOss(filename string, fd io.Reader) string {
	conf := &config.GetConfig().OSS
	fname := uuid.NewV4().String() + path.Ext(filename)
	err := bucket.PutObject(conf.Path+fname, fd)
	pictureUrl := conf.BaseURL + conf.Path + fname
	if err != nil {
		log.Errorln("File upload to OSS fail，fileName：", pictureUrl, ", err: :", err)
		return ""
	}
	return pictureUrl
}

func getGmtIso8601(expireEnd int64) string {
	var tokenExpire = time.Unix(expireEnd, 0).UTC().Format("2006-01-02T15:04:05Z")
	return tokenExpire
}

// GetPolicyToken 客户端直接上传OSS需要的配置 https://help.aliyun.com/document_detail/91818.htm?spm=a2c4g.11186623.0.0.1607566anAGeY2#concept-mhj-zzt-2fb
func GetPolicyToken() interface{} {
	conf := &config.GetConfig().OSS
	now := time.Now().Unix()
	expireEnd := now + conf.ExpireTime
	var tokenExpire = getGmtIso8601(expireEnd)

	//create post policy json
	var policyConfig ConfigStruct
	policyConfig.Expiration = tokenExpire
	var condition []string
	condition = append(condition, "starts-with")
	condition = append(condition, "$key")
	condition = append(condition, conf.Path)
	policyConfig.Conditions = append(policyConfig.Conditions, condition)

	//calucate signature
	result, _ := json.Marshal(policyConfig)
	debyte := base64.StdEncoding.EncodeToString(result)
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(conf.AccessKeySecret))
	io.WriteString(h, debyte)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	var callbackParam CallbackParam
	callbackParam.CallbackUrl = conf.CallbackUrl
	callbackParam.CallbackBody = "filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}"
	callbackParam.CallbackBodyType = "application/x-www-form-urlencoded"
	callbackStr, err := json.Marshal(callbackParam)
	if err != nil {
		log.Errorln("callback json err:", err)
	}
	callbackBase64 := base64.StdEncoding.EncodeToString(callbackStr)

	var policyToken PolicyToken
	policyToken.AccessKeyId = conf.AccessKeyId
	policyToken.Host = conf.BaseURL
	policyToken.Expire = expireEnd
	policyToken.Signature = signedStr
	policyToken.Directory = conf.Path
	policyToken.Policy = debyte
	policyToken.Callback = callbackBase64
	policyToken.FileNamePrefix = conf.BaseURL + conf.Path + uuid.NewV4().String()
	return policyToken
}
