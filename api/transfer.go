package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/vpaklatzis/go-simple-bank/db/sqlc"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required, min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required, min=1"`
	Amount        int64  `json:"amount" binding:"required, gt=0"`
	Currency      string `json:"currency" binding:"required,oneof=USD EUR"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	transfer, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, transfer)
}
