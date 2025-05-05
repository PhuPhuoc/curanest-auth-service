package accountnotirpc

import (
	"context"
	"fmt"

	"github.com/PhuPhuoc/curanest-auth-service/common"
	accountqueries "github.com/PhuPhuoc/curanest-auth-service/module/account/usecase/queries"
)

func (rpc *externalNotiService) RegisterPushTokenRPC(ctx context.Context, reqBody *accountqueries.RegisPushToken) error {
	fmt.Println("rpcUrl: ", rpc.apiURL)
	response, err := common.CallExternalAPI(ctx, common.RequestOptions{
		Method:  "POST",
		URL:     rpc.apiURL + "/external/rpc/notifications/push-token",
		Payload: reqBody,
	})
	if err != nil {
		resp := common.NewInternalServerError().WithReason("cannot call external api: " + err.Error())
		return resp
	}

	success, ok := response["success"].(bool)
	if !ok || !success {
		return common.ExtractErrorFromResponse(response)
	}

	return nil
}
