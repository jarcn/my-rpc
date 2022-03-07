package xclient

import (
	"errors"
	"math"
	"math/rand"
	"sync"
	"time"
)

// Discovery 是一个接口类型,包含了服务发现所需的最基本的接口
// Refresh() 从注册中心更新服务列表
// Update(servers []string) 手动更细服务列表
// Get(mode SelectMode) 根据负载均衡策略,选择一个服务实例
// GetAll() 返回所有的服务实例

type SelectMode int

const (
	RandomSelect SelectMode = iota //自动加1
	RoundRobinSelect
)

type Discovery interface {
	Refresh() error                      //从注册中心更新服务列表
	Update(servers []string) error       //手动更细服务列表
	Get(mode SelectMode) (string, error) //根据负载均衡策略,选择一个服务实例
	GetAll() ([]string, error)           //返回所有的服务实例
}

type MultiServersDiscovery struct {
	r       *rand.Rand   //随机数
	mu      sync.RWMutex //保证请求的时序性
	servers []string     //服务列表
	index   int          //被命中的服务序号
}

//创建服务发现实例
func NewMultiServerDiscovery(servers []string) *MultiServersDiscovery {
	d := &MultiServersDiscovery{
		servers: servers,
		r:       rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	d.index = d.r.Intn(math.MaxInt32 - 1)
	return d
}

var _ Discovery = (*MultiServersDiscovery)(nil)

func (d *MultiServersDiscovery) Refresh() error {
	return nil
}

func (d *MultiServersDiscovery) Update(servers []string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.servers = servers
	return nil
}

func (d *MultiServersDiscovery) Get(mode SelectMode) (string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	n := len(d.servers)
	if n == 0 {
		return "", errors.New("rps discovery: no available servers")
	}
	switch mode {
	case RandomSelect:
		return d.servers[d.r.Intn(n)], nil
	case RoundRobinSelect:
		s := d.servers[d.index%n]
		d.index = (d.index + 1) % n
		return s, nil
	default:
		return "", errors.New("rpc discovery: not supported select mode")
	}
}

func (d *MultiServersDiscovery) GetAll() ([]string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	servers := make([]string, len(d.servers), len(d.servers))
	copy(servers, d.servers)
	return servers, nil
}
