package db

import (
	"context"

	"github.com/scalarorg/xchains-api/internal/types"
	"github.com/scalarorg/xchains-api/internal/utils"
)

func (db *Database) TransitionToWithdrawnState(ctx context.Context, txHashHex string) error {
	err := db.transitionState(
		ctx, txHashHex, types.Withdrawn.ToString(),
		utils.QualifiedStatesToWithdraw(), nil,
	)
	if err != nil {
		return err
	}
	return nil
}
