package images

import "golang.org/x/net/context"

func (rule *PhotoRule) Enforce(ctx context.Context, data interface{}) (success bool, message string, err error) {

	if rule.Big != nil && rule.Big.Value() {
		photo := data.(*Photo)
		if photo.Width.Value() < 1200 {
			return false, "Photo must be big!", nil
		}
	}

	return true, "", nil
}
