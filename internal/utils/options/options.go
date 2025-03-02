package utils

var defaultRestaurantOptions = restaurantOptions{
	count:  10,
	offset: 0,
}

type restaurantOptions struct {
	count  int
	offset int
}

type applyRestaurantOption interface {
	apply(*restaurantOptions)
}

type funcRestaurantOption struct {
	f func(option *restaurantOptions)
}

func (fdo *funcRestaurantOption) apply(opt *restaurantOptions) {
	fdo.f(opt)
}

func newFuncRestaurantOption(f func(option *restaurantOptions)) *funcRestaurantOption {
	return &funcRestaurantOption{
		f: f,
	}
}

func WithCustomCount(count, total int) applyRestaurantOption {
	return newFuncRestaurantOption(func(o *restaurantOptions) {
		if count >= 0 && count <= total {
			o.count = count
		}
		if count > total {
			o.count = total
		}
		if count < 0 {
			o.count = 0
		}
	})
}

func WithCustomOffset(offset, total int) applyRestaurantOption {
	return newFuncRestaurantOption(func(o *restaurantOptions) {
		if offset >= 0 && offset < total {
			o.offset = offset
		} else {
			o.count = 0
		}
	})
}

type Options struct {
	opts restaurantOptions
}

func NewOptions(opts ...applyRestaurantOption) *Options {
	options := defaultRestaurantOptions
	for _, option := range opts {
		option.apply(&options)
	}
	return &Options{opts: options}
}

func (o *Options) GetCount() int {
	return o.opts.count
}

func (o *Options) GetOffset() int {
	return o.opts.offset
}
