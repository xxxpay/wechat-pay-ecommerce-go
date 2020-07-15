package ecommerce

import (
	"encoding/json"
	"net/http"
	"github.com/xxxpay/wechat-pay-ecommerce-go/core"
	"github.com/xxxpay/wechat-pay-ecommerce-go/dto"
)

func (c *payClient) Certificate() (*dto.CertificateResp, error) {
	body, err := c.doRequest(nil, core.BuildUrl(nil, nil, core.ApiCertification), http.MethodGet)
	if err != nil {
		return nil, err
	}
	var resp dto.CertificateResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	for _, data := range resp.Data {
		encryptCert := data.EncryptCertificate
		if encryptCert == nil {
			continue
		}
		decryptCert, err := c.Decrypt(encryptCert.Algorithm, encryptCert.Ciphertext, encryptCert.AssociatedData, encryptCert.Nonce)
		if err != nil {
			return nil, err
		}
		data.DecryptCertificate = string(decryptCert)
	}
	return &resp, nil
}
