package registry

import (
	"log"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"
)

type MyRegistry struct {
	timeout time.Duration
	mu      sync.Mutex
	severs  map[string]*ServerItem
}

type ServerItem struct {
	Addr  string
	start time.Time
}

const (
	defaultPath    = "myrpc/registry"
	defaultTimeout = time.Minute * 5
)

func New(timeout time.Duration) *MyRegistry {
	return &MyRegistry{
		severs:  make(map[string]*ServerItem),
		timeout: timeout,
	}
}

var DefaultMyRegister = New(defaultTimeout)

//添加服务实例接口,如果服务存在则更新start
func (r *MyRegistry) putServer(addr string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	s := r.severs[addr]
	if s == nil {
		r.severs[addr] = &ServerItem{Addr: addr, start: time.Now()}
	} else {
		s.start = time.Now()
	}
}

// 返回可用服务列表,如果存在超时的服务,则删除
func (r *MyRegistry) aliveServers() []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	var alive []string
	for addr, s := range r.severs {
		if r.timeout == 0 || s.start.Add(r.timeout).After(time.Now()) {
			alive = append(alive, addr)
		} else {
			delete(r.severs, addr)
		}
	}
	sort.Strings(alive)
	return alive
}

//使用http协议进行服务提供
func (r *MyRegistry) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		w.Header().Set("X-Geerpc-Servers", strings.Join(r.aliveServers(), ",")) // keep it simple, server is in req.Header
	case "POST":
		addr := req.Header.Get("X-Geerpc-Server") // keep it simple, server is in req.Header
		if addr == "" {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		r.putServer(addr)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// HandleHTTP registers an HTTP handler for GeeRegistry messages on registryPath
func (r *MyRegistry) HandleHTTP(registryPath string) {
	http.Handle(registryPath, r)
	log.Println("rpc registry path:", registryPath)
}

func HandleHTTP() {
	DefaultMyRegister.HandleHTTP(defaultPath)
}

//服务存活性检测,通过http心跳实现
func Heartbeat(registry, addr string, duration time.Duration) {
	if duration == 0 {
		duration = defaultTimeout - time.Duration(1)*time.Minute
	}
	var err error
	err = sendHeartbeat(registry, addr)
	go func() {
		t := time.NewTicker(duration)
		for err == nil {
			<-t.C
			err = sendHeartbeat(registry, addr)
		}
	}()
}

func sendHeartbeat(registry, addr string) error {
	log.Println(addr, "send heart beat to registry", registry)
	httpClient := &http.Client{}
	req, _ := http.NewRequest("POST", registry, nil)
	req.Header.Set("X-Myrpc-Server", addr)
	if _, err := httpClient.Do(req); err != nil {
		log.Println("rpc server: heart beat err:", err)
		return err
	}
	return nil
}
