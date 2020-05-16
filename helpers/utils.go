package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"

	"github.com/sony/sonyflake"
)

var flake = sonyflake.NewSonyflake(sonyflake.Settings{})

func GetUuidString() string {
	return strconv.FormatUint(GetUuidUInt64(), 10)
}

func GetUuidUInt64() uint64 {
	uid, _ := flake.NextID()
	return uid
}

func GenStrMd5sum(OriginStr string) string {
	var byteData = []byte(OriginStr)
	ctx := md5.New()
	ctx.Write(byteData)
	md5Str := hex.EncodeToString(ctx.Sum(nil))

	return md5Str
}
