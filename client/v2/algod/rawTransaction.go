package algod

import (
	"context"
	"strings"

	"github.com/algorand/go-algorand-sdk/client/v2/common"
	"github.com/algorand/go-algorand-sdk/client/v2/common/models"
)

type SendRawTransaction struct {
	c *Client

	rawtxn []byte
}

func (s *SendRawTransaction) Do(ctx context.Context, headers ...*common.Header) (txid string, err error) {
	var response models.PostTransactionsResponse
	// Set default Content-Type, if the user didn't specify it.
	addContentType := true
	for _, header := range headers {
		if strings.ToLower(header.Key) == "content-type" {
			addContentType = false
			break
		}
	}
	if addContentType {
		headers = append(headers, &common.Header{"Content-Type", "application/x-binary"})
	}
	err = s.c.post(ctx, &response, "/v2/transactions", s.rawtxn, headers)
	txid = response.Txid
	return
}
