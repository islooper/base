package sharedconfig

// SharedConfig 基础配置结构体
// 推荐以内联形式嵌入到具体配置结构体中
type SharedConfig struct {
	DB    map[string]DBConf
}
