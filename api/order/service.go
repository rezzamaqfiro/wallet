package order

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/rezzamaqfiro/wallet/constant"
	"github.com/rezzamaqfiro/wallet/util"

	repo "github.com/rezzamaqfiro/wallet/repo/generated"
)

func (h *Handler) OrderDisbursement(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp := util.NewResponse(http.StatusInternalServerError, http.StatusInternalServerError, "Internal Server Error", map[string]interface{}{})
	defer resp.WriteResponse(w, r)

	//proper to add authorization here

	payload := DisbursementRequest{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		errMsg := "Payload error to parse"
		log.Printf("%s with err: %v", errMsg, err)
		resp.Message = errMsg
		return
	}

	invoiceID, err := uuid.Parse(payload.InvoiceID)
	if err != nil {
		errMsg := "Invoice ID got bad request"
		log.Printf("%s with err: %v", errMsg, err)
		resp.Message = errMsg
		resp.Status = http.StatusBadRequest
		resp.Code = http.StatusBadRequest
		return
	}

	if strings.ToUpper(payload.Status) != constant.OrderDisburseApproval {
		errMsg := "Wrong Status Request"
		log.Println(errMsg)
		resp.Message = errMsg
		resp.Status = http.StatusBadRequest
		resp.Code = http.StatusBadRequest
		return
	}

	order, err := h.db.GetOrderByInvoiceID(ctx, invoiceID)
	if err != nil {
		errMsg := "Order not found"
		log.Printf("%s with err: %v", errMsg, err)
		resp.Message = errMsg
		resp.Status = http.StatusBadRequest
		resp.Code = http.StatusBadRequest
		return
	}

	user, err := h.db.GetBalanceByUserID(ctx, order.UserID.UUID)
	if err != nil {
		errMsg := "User not found"
		log.Printf("%s with err: %v", errMsg, err)
		resp.Message = errMsg
		resp.Status = http.StatusBadRequest
		resp.Code = http.StatusBadRequest
		return
	}

	if order.Status.String != "" && order.Status.String != constant.OrderStatusPending {
		errMsg := "Order can't be disbursed"
		log.Printf("%s, status: %s", errMsg, order.Status.String)
		resp.Message = fmt.Sprintf("%s, status: %s", errMsg, order.Status.String)
		resp.Status = http.StatusBadRequest
		resp.Code = http.StatusBadRequest
		return
	}

	orderStatusParam := repo.UpdateOrderStatusByInvoiceIDParams{
		Status:  util.SqlString(constant.OrderStatusComplete),
		Invoice: invoiceID,
	}
	_, err = h.db.UpdateOrderStatusByInvoiceID(ctx, orderStatusParam)
	if err != nil {
		errMsg := "Failed to update status order"
		log.Printf("%s, err: %v", errMsg, err)
		resp.Message = errMsg
		return
	}

	userBalanceParam := repo.UpdateBalanceByUserIDParams{
		Balance: order.TotalPayment,
		UserID:  user.UserID,
	}
	_, err = h.db.UpdateBalanceByUserID(ctx, userBalanceParam)
	if err != nil {
		errMsg := "Failed to update balance user"
		log.Printf("%s, err: %v", errMsg, err)
		resp.Message = errMsg
		return
	}

	resp.Status = http.StatusOK
	resp.Code = http.StatusOK
	resp.Message = "Request has been disbursed."
}
