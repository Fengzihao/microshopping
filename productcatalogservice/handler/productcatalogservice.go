package handler

import (
	"bytes"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"os"
	"os/signal"
	pb "productcatalogservice/proto"
	"strings"
	"sync"
	"syscall"
)

var reloadCatalog bool

// 日志
var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger:", log.Lshortfile)
)

// ProductCatalogService 商品分类结构体
type ProductCatalogService struct {
	sync.Mutex
	products []*pb.Product
}

// 初始化
func init() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		for {
			sig := <-sigs
			logger.Printf("接收信号: %s", sig)
			if sig == syscall.SIGUSR1 {
				reloadCatalog = true
				logger.Printf("可以加载商品信息")
			} else {
				reloadCatalog = false
				logger.Println("不能加载商品信息")
			}
		}
	}()
}

// readCatalogFile 读取配置的数据json文件
func (p *ProductCatalogService) readCatalogFile() (*pb.ListProductsResponse, error) {
	p.Lock()
	defer p.Unlock()
	catalogJson, err := os.ReadFile("data/products.json")
	if err != nil {
		logger.Printf("打开商品json文件失败:%v", err)
		return nil, err
	}
	catalog := &pb.ListProductsResponse{}
	if err := protojson.Unmarshal(catalogJson, catalog); err != nil {
		logger.Printf("序列化json文件失败:%v", err)
		return nil, err
	}
	logger.Printf("序列号json文件成功")
	return catalog, nil
}

// 解析配置数据json文件
func (p *ProductCatalogService) parseCatalog() []*pb.Product {
	if reloadCatalog || len(p.products) == 0 {
		catalog, err := p.readCatalogFile()
		if err != nil {
			return []*pb.Product{}
		}
		p.products = catalog.Products
	}
	return p.products
}

// ListProducts 获取商品列表
func (p *ProductCatalogService) ListProducts(ctx context.Context, in *pb.Empty) (out *pb.ListProductsResponse, e error) {
	out = new(pb.ListProductsResponse)
	out.Products = p.parseCatalog() // todo 这里需要获取解析json文件的方法
	return out, nil
}

// GetProduct 获取单个商品
func (p *ProductCatalogService) GetProduct(ctx context.Context, in *pb.GetProductRequest) (out *pb.Product, e error) {
	var found *pb.Product
	out = new(pb.Product)
	products := p.parseCatalog()
	for _, product := range products {
		if in.Id == product.Id {
			found = product
		}
	}
	if found == nil {
		return out, status.Errorf(codes.NotFound, "no product with ID %s", in.Id)
	}
	out.Id = found.Id
	out.Name = found.Name
	out.Categories = found.Categories
	out.Description = found.Description
	out.PriceUsd = found.PriceUsd
	return out, nil
}

func (p *ProductCatalogService) SearchProducts(ctx context.Context, in *pb.SearchProductsRequest) (out *pb.SearchProductsResponse, e error) {
	var ps []*pb.Product
	products := p.parseCatalog()
	for _, product := range products {
		if strings.Contains(strings.ToLower(product.Name), strings.ToLower(in.Query)) ||
			strings.Contains(strings.ToLower(product.Description), strings.ToLower(in.Query)) {
			ps = append(ps, product)
		}
	}
	out.Results = ps
	return out, nil
}
