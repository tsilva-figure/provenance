package antewrapper

import (
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

type ValidateStateful = func(ctx types.Context, m types.Msg) error

type StatefulValidatorCreator interface {
	StatefulValidator() ValidateStateful
}

type StatefulValidatorAnteDecorator struct {
	manager *module.Manager
}

func NewStatefulValidatorAnteDecorator(manager *module.Manager) types.AnteDecorator {
	return &StatefulValidatorAnteDecorator{manager}
}

func (h *StatefulValidatorAnteDecorator) AnteHandle(ctx types.Context, tx types.Tx, simulate bool, next types.AnteHandler) (types.Context, error) {
	for _, m := range tx.GetMsgs() {
		module := h.manager.Modules[m.Route()]
		if module == nil {
			continue
		}

		// Implements the stateful validator type?
		sv, ok := module.(StatefulValidatorCreator)
		if !ok {
			continue
		}

		validator := sv.StatefulValidator()
		if err := validator(ctx, m); err != nil {
			return types.Context{}, err
		}
	}
	return next(ctx, tx, simulate)
}
