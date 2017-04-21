package images

// notest

import "context"

func (rule *PhotoRule) Enforce(ctx context.Context, data interface{}) (fail bool, messages []string, err error) {

	if i, ok := data.(PhotoInterface); ok {
		data = i.GetPhoto(ctx)
	}

	if rule.Big != nil && rule.Big.Value() {
		photo := data.(*Photo)
		if photo.Width.Value() < 1200 {
			fail = true
			messages = append(messages, "Photo must be big!")
		}
	}
	return
}
