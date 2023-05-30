package conf

import "github.com/olivere/elastic/v7"

var (
	EsClient *elastic.Client
)

// NewEs 初始化es
func NewEs() error {
	// 创建ES client用于后续操作ES
	client, err := elastic.NewClient(
		// 设置ES服务地址，支持多个地址
		elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		// 设置基于http base auth验证的账号和密码
		// elastic.SetBasicAuth("user", "secret")
	)

	if err != nil {
		return err
	}

	EsClient = client
	return nil
}
