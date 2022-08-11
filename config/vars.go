package config

type GlobalConfig struct {
	MODE        string `yaml:"Mode"`
	ProgramName string `yaml:"ProgramName"`
	AUTHOR      string `yaml:"Author"`
	VERSION     string `yaml:"Version"`
	Host        string `yaml:"Host"`
	Port        string `yaml:"Port"`
	LogPath     string `yaml:"LogPath"`
	Auth        struct {
		Secret string `yaml:"Secret"`
		Issuer string `yaml:"Issuer"`
	} `yaml:"Auth"`
	Databases []Datasource `yaml:"Databases"`
	REDIS     struct {
		Use    bool `yaml:"Use"`
		Config struct {
			IP       string `yaml:"Ip"`
			PORT     string `yaml:"Port"`
			PASSWORD string `yaml:"Password"`
			DB       int    `yaml:"Db"`
		} `yaml:"Config"`
	} `yaml:"Redis"`
	OSS struct {
		Use    bool `yaml:"Use"`
		Config struct {
			AccessKeySecret string `yaml:"AccessKeySecret"`
			AccessKeyId     string `yaml:"AccessKeyId"`
			EndPoint        string `yaml:"EndPoint"`
			BucketName      string `yaml:"BucketName"`
			BaseURL         string `yaml:"BaseURL"`
			Path            string `yaml:"Path"`
			CallbackUrl     string `yaml:"CallbackUrl"`
			ExpireTime      int64  `yaml:"ExpireTime"`
		} `yaml:"Config"`
	} `yaml:"Oss"`
	Mail struct {
		Use    bool `yaml:"Use"`
		Config struct {
			SMTP     string `yaml:"Smtp"`
			PORT     int    `yaml:"Port"`
			ACCOUNT  string `yaml:"Account"`
			PASSWORD string `yaml:"Password"`
		} `yaml:"Config"`
	} `yaml:"Mail"`
	CMS struct {
		Use    bool `yaml:"Use"`
		Config struct {
			SecretId   string `yaml:"SecretId"`
			SecretKey  string `yaml:"SecretKey"`
			AppId      string `yaml:"AppId"`
			TemplateId string `yaml:"TemplateId"`
			Sign       string `yaml:"Sign"`
		} `yaml:"Config"`
	} `yaml:"Cms"`
}

type Datasource struct {
	Key      string `yaml:"Key"`
	Type     string `yaml:"Type"`
	IP       string `yaml:"Ip"`
	PORT     string `yaml:"Port"`
	USER     string `yaml:"User"`
	PASSWORD string `yaml:"Password"`
	DATABASE string `yaml:"Database"`
}
