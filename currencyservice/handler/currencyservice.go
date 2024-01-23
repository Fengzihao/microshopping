package handler

import (
	"context"
	pb "currencyservice/proto"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math"
	"os"
)

type CurrencyService struct {
}

// GetSupportedCurrencies 获取货币
func (c *CurrencyService) GetSupportedCurrencies(ctx context.Context, in *pb.Empty) (out *pb.GetSupportedCurrenciesResponse, e error) {
	data, err := os.ReadFile("data/currency_conversion.json")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "加载货币数据失败：%+v", err)
	}
	currencies := make(map[string]float32)
	if err := json.Unmarshal(data, &currencies); err != nil {
		return nil, status.Errorf(codes.Internal, "解析货币数据失败：%+v", err)
	}
	fmt.Printf("货币：%v\n", currencies)

	out = new(pb.GetSupportedCurrenciesResponse)
	out.CurrencyCodes = make([]string, 0, len(currencies))
	for s := range currencies {
		out.CurrencyCodes = append(out.CurrencyCodes, s)
	}
	return out, nil
}

func (c *CurrencyService) Convert(ctx context.Context, in *pb.CurrencyConversionRequest) (out *pb.Money, e error) {
	data, err := os.ReadFile("data/currency_conversion.json")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "加载货币失败: %+v", err)
	}
	currencies := make(map[string]float64)
	if err := json.Unmarshal(data, &currencies); err != nil {
		return nil, status.Errorf(codes.Internal, "解析货币失败: %+v", err)
	}
	fromCurrency, found := currencies[in.From.CurrencyCode]
	if !found {
		return nil, status.Errorf(codes.InvalidArgument, "不支持的币种: %s", in.From.CurrencyCode)
	}
	toCurrency, found := currencies[in.ToCode]
	if !found {
		return nil, status.Errorf(codes.InvalidArgument, "不支持的币种: %s", in.ToCode)
	}
	out = new(pb.Money)
	out.CurrencyCode = in.ToCode
	total := int64(math.Floor(float64(in.From.Units*10^9+int64(in.From.Nanos)) / fromCurrency * toCurrency))
	out.Units = total / 1e9
	out.Nanos = int32(total % 1e9)
	return out, nil
}
