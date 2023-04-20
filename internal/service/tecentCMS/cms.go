package tecentCMS

import (
	"encoding/json"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111" // 引入sms
	"github.com/wujunyi792/flamego-quick-template/config"
	"github.com/wujunyi792/flamego-quick-template/internal/core/logx"
)

func SendCMS(phone string, parameters []string) bool {
	conf := &config.GetConfig().CMS
	credential := common.NewCredential(
		conf.SecretId,
		conf.SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	cpf.SignMethod = "HmacSHA1"
	client, _ := sms.NewClient(credential, "ap-guangzhou", cpf)
	request := sms.NewSendSmsRequest()
	/* 基本类型的设置:
	 * SDK采用的是指针风格指定参数，即使对于基本类型你也需要用指针来对参数赋值。
	 * SDK提供对基本类型的指针引用封装函数
	 * 帮助链接：
	 * 短信控制台: https://console.cloud.tencent.com/smsv2
	 * sms helper: https://cloud.tencent.com/document/product/382/3773 */
	/* 短信应用ID: 短信SdkAppId在 [短信控制台] 添加应用后生成的实际SdkAppId，示例如1400006666 */
	request.SmsSdkAppId = common.StringPtr(conf.AppId)

	/* 短信签名内容: 使用 UTF-8 编码，必须填写已审核通过的签名，签名信息可登录 [短信控制台] 查看 */
	request.SignName = common.StringPtr(conf.Sign)

	/* 国际/港澳台短信 SenderId: 国内短信填空，默认未开通，如需开通请联系 [sms helper] */
	request.SenderId = common.StringPtr("")
	/* 用户的 session 内容: 可以携带用户侧 ID 等上下文信息，server 会原样返回 */
	request.SessionContext = common.StringPtr(phone)
	/* 短信码号扩展号: 默认未开通，如需开通请联系 [sms helper] */
	request.ExtendCode = common.StringPtr("")
	/* 模板参数: 若无模板参数，则设置为空*/
	request.TemplateParamSet = common.StringPtrs(parameters) //这边parameters是字符数组类型，所以可以直接传

	/* 模板 ID: 必须填写已审核通过的模板 ID。模板ID可登录 [短信控制台] 查看 */
	request.TemplateId = common.StringPtr(conf.TemplateId)

	/* 下发手机号码，采用 E.164 标准，+[国家或地区码][手机号]
	 * 示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号，最多不要超过200个手机号*/
	request.PhoneNumberSet = common.StringPtrs([]string{"+86" + phone})
	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.SendSms(request)
	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		logx.NameSpace("cms").Errorln("An API error has returned: %s", err)
		return false
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		logx.NameSpace("cms").Errorln("An SDK error has returned: %s", err)
		return false
	}
	b, _ := json.Marshal(response.Response)
	// 打印返回的json字符串
	logx.NameSpace("cms").Infof("SendSMS: %s -> %s", phone, string(b))

	return *response.Response.SendStatusSet[0].Code == "Ok"
}
