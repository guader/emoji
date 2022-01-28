package provider

type SliceProvider struct{}

func (p *SliceProvider) CodePoints() ([][]rune, error) {
	return sliceProviderCodePoints, nil
}

func NewSliceProvider() *SliceProvider {
	return &SliceProvider{}
}
