package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type ConfigClient struct {
	Path string
}

type Client struct {
	Db     *kubernetes.Clientset // 驱动
	Config *ConfigClient         // 配置
}

func NewClient(config *ConfigClient) (*Client, error) {

	var err error
	c := &Client{Config: config}

	cfg, err := clientcmd.BuildConfigFromFlags("", c.Config.Path)
	if err != nil {
		return nil, err
	}
	c.Db, err = kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}

	return c, nil
}
