package models

import (
	"context"

	"github.com/davelondon/kerr"
	"kego.io/process/validate"
	"kego.io/system"
	"kego.io/system/node"
)

type NodeModel struct {
	Node               *node.Node
	Invalid            bool
	Errors             []validate.ValidationError
	ShowSystemControls bool
}

func (n *NodeModel) Validate(ctx context.Context, rules []system.RuleInterface) (changed bool, err error) {

	invalid := false
	errors := []validate.ValidationError{}

	in := func(n validate.ValidationError, h []validate.ValidationError) bool {
		for _, v := range h {
			if n.Description == v.Description {
				return true
			}
		}
		return false
	}

	eq := func(a, b []validate.ValidationError) bool {

		if a == nil && b == nil {
			return true
		}

		if a == nil || b == nil {
			return false
		}

		if len(a) != len(b) {
			return false
		}

		for _, v := range a {
			if !in(v, b) {
				return false
			}
		}

		return true
	}

	hasChanged := func() bool {
		if n.Invalid != invalid {
			return true
		}
		if !invalid {
			return false
		}
		return !eq(n.Errors, errors)
	}

	if v, ok := n.Node.Value.(system.Validator); ok {
		failed, messages, err := v.Validate(ctx)
		if err != nil {
			return false, kerr.Wrap("BJEWACGCUJ", err)
		}
		if failed {
			invalid = true
			for _, message := range messages {
				errors = append(errors, validate.ValidationError{Struct: kerr.New("QEMMGBTXBC", message), Source: n.Node})
			}
			// if the object isn't valid we don't want to continue
			changed = hasChanged()
			n.Invalid = invalid
			n.Errors = errors
			return changed, nil
		}
	}

	for _, rule := range rules {
		en, ok := rule.(system.Enforcer)
		if !ok {
			continue
		}
		failed, messages, err := en.Enforce(ctx, n.Node.Value)
		if err != nil {
			return false, kerr.Wrap("TOJMUPODPT", err)
		} else if failed {
			invalid = true
			for _, m := range messages {
				errors = append(errors, validate.ValidationError{
					Struct: kerr.New("JSATVFEBEJ", m),
					Source: n.Node,
				})
			}
		}
	}
	changed = hasChanged()
	n.Invalid = invalid
	n.Errors = errors
	return changed, nil
}
