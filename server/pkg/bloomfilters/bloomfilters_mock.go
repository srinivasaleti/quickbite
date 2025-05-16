package bloomfilters

import "github.com/stretchr/testify/mock"

type BloomFilterMock struct {
	mock.Mock
}

func (filter *BloomFilterMock) Load(filenames []string) error {
	args := filter.Called(filenames)
	return args.Error(0)
}

func (filter *BloomFilterMock) ElmentExistsInWhichFiles(coupon string) []string {
	args := filter.Mock.Called(coupon)
	result, _ := args.Get(0).([]string)
	return result
}

func (filter *BloomFilterMock) IsLoaded() bool {
	args := filter.Mock.Called()
	result, _ := args.Get(0).(bool)
	return result
}

func (filter *BloomFilterMock) Reset() {
	filter.ExpectedCalls = nil
}
