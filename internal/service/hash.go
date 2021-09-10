package service

import (
	"context"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"

	"github.com/hatlonely/go-kit/rpcx"
	"github.com/spaolacci/murmur3"
	"google.golang.org/grpc/codes"

	"github.com/hatlonely/rpc-tool/api/gen/go/api"
)

func (s *ToolService) Hash(ctx context.Context, req *api.HashReq) (*api.HashRes, error) {
	var buf []byte
	if len(req.Base64) != 0 {
		var err error
		buf, err = base64.StdEncoding.DecodeString(req.Base64)
		if err != nil {
			return nil, rpcx.NewError(err, codes.InvalidArgument, "InvalidArgument.Base64", "invalid base64 format")
		}
	} else {
		buf = []byte(req.Text)
	}

	if h, ok := hash[req.Hash]; ok {
		num, hex := h(buf)
		return &api.HashRes{
			Num: num,
			Hex: hex,
		}, nil
	}

	return nil, rpcx.NewErrorf(nil, codes.InvalidArgument, "InvalidArgument.Hash", "unsupported hash type [%v]", req.Hash)
}

var hash = map[string]func(buf []byte) (string, string){
	"fnv32": func(buf []byte) (string, string) {
		h := fnv.New32()
		_, _ = h.Write(buf)
		b := h.Sum(nil)
		return fmt.Sprintf("%d", binary.BigEndian.Uint32(b)), hex.EncodeToString(b)
	},
	"fnv64": func(buf []byte) (string, string) {
		h := fnv.New64()
		_, _ = h.Write(buf)
		b := h.Sum(nil)
		return fmt.Sprintf("%d", binary.BigEndian.Uint64(b)), hex.EncodeToString(b)
	},
	"fnv128": func(buf []byte) (string, string) {
		h := fnv.New128()
		_, _ = h.Write(buf)
		b := h.Sum(nil)
		return fmt.Sprintf("%d, %d", binary.BigEndian.Uint64(b[0:8]), binary.BigEndian.Uint64(b[8:16])), hex.EncodeToString(b)
	},
	"fnv32a": func(buf []byte) (string, string) {
		h := fnv.New32a()
		_, _ = h.Write(buf)
		b := h.Sum(nil)
		return fmt.Sprintf("%d", binary.BigEndian.Uint32(b)), hex.EncodeToString(b)
	},
	"fnv64a": func(buf []byte) (string, string) {
		h := fnv.New64a()
		_, _ = h.Write(buf)
		b := h.Sum(nil)
		return fmt.Sprintf("%d", binary.BigEndian.Uint64(b)), hex.EncodeToString(b)
	},
	"fnv128a": func(buf []byte) (string, string) {
		h := fnv.New128a()
		_, _ = h.Write(buf)
		b := h.Sum(nil)
		return fmt.Sprintf("%d, %d", binary.BigEndian.Uint64(b[0:8]), binary.BigEndian.Uint64(b[8:16])), hex.EncodeToString(b)
	},
	"crc32": func(buf []byte) (string, string) {
		h := crc32.NewIEEE()
		_, _ = h.Write(buf)
		b := h.Sum(nil)
		return fmt.Sprintf("%d", binary.BigEndian.Uint32(b)), hex.EncodeToString(b)
	},
	"crc64-ecma": func(buf []byte) (string, string) {
		h := crc64.New(crc64.MakeTable(crc64.ECMA))
		_, _ = h.Write(buf)
		b := h.Sum(nil)
		return fmt.Sprintf("%d", binary.BigEndian.Uint64(b)), hex.EncodeToString(b)
	},
	"crc64-iso": func(buf []byte) (string, string) {
		h := crc64.New(crc64.MakeTable(crc64.ISO))
		_, _ = h.Write(buf)
		b := h.Sum(nil)
		return fmt.Sprintf("%d", binary.BigEndian.Uint64(b)), hex.EncodeToString(b)
	},
	"adler32": func(buf []byte) (string, string) {
		h := adler32.New()
		_, _ = h.Write(buf)
		b := h.Sum(nil)
		return fmt.Sprintf("%d", binary.BigEndian.Uint32(b)), hex.EncodeToString(b)
	},
	"murmur3-32": func(buf []byte) (string, string) {
		h := murmur3.New32()
		_, _ = h.Write(buf)
		b := h.Sum(nil)
		return fmt.Sprintf("%d", binary.BigEndian.Uint32(b)), hex.EncodeToString(b)
	},
	"murmur3-64": func(buf []byte) (string, string) {
		h := murmur3.New64()
		_, _ = h.Write(buf)
		b := h.Sum(nil)
		return fmt.Sprintf("%d", binary.BigEndian.Uint64(b)), hex.EncodeToString(b)
	},
	"murmur3-128": func(buf []byte) (string, string) {
		h := murmur3.New128()
		_, _ = h.Write(buf)
		b := h.Sum(nil)
		return fmt.Sprintf("%d, %d", binary.BigEndian.Uint64(b[0:8]), binary.BigEndian.Uint64(b[8:16])), hex.EncodeToString(b)
	},
}
