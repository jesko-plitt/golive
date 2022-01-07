package cloud

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/ao-concepts/logging"
)

type ImageProxy struct {
	cfg *ImageProxyConfig
	log logging.Logger
}

func ProvideImageProxy(log logging.Logger) *ImageProxy {
	return &ImageProxy{
		cfg: ProvideImageproxyConfig(),
		log: log,
	}
}

func (ip *ImageProxy) GetFileUrl(hash string, width int, height int) string {
	var keyBin, saltBin []byte
	var err error

	if keyBin, err = hex.DecodeString(ip.cfg.Key); err != nil {
		ip.log.Fatal("Key expected to be hex-encoded string")
	}

	if saltBin, err = hex.DecodeString(ip.cfg.Salt); err != nil {
		ip.log.Fatal("Salt expected to be hex-encoded string")
	}

	resize := "fit"
	gravity := "no"
	enlarge := 1
	extension := "jpg"

	url := "s3://" + ip.cfg.BucketName + "/" + hash
	encodedURL := base64.RawURLEncoding.EncodeToString([]byte(url))

	path := fmt.Sprintf("/%s/%d/%d/%s/%d/%s.%s", resize, width, height, gravity, enlarge, encodedURL, extension)

	mac := hmac.New(sha256.New, keyBin)
	mac.Write(saltBin)
	mac.Write([]byte(path))
	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	return fmt.Sprintf("%s%s/%s%s", ip.cfg.URL, ip.cfg.Prefix, signature, path)
}
