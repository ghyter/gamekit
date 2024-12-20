package mouse

import "time"

// mouseOptions defines configurable settings for the Mouse struct.
type mouseOptions struct {
	doubleClickThreshold time.Duration
	holdThreshold        time.Duration
}

// Option defines a functional option for configuring mouseOptions.
type MouseOption func(*mouseOptions)

// WithDoubleClickThreshold sets the double-click threshold for the Mouse instance.
func WithDoubleClickThreshold(threshold time.Duration) MouseOption {
	return func(o *mouseOptions) {
		o.doubleClickThreshold = threshold
	}
}

// WithHoldThreshold sets the hold threshold for the Mouse instance.
func WithHoldThreshold(threshold time.Duration) MouseOption {
	return func(o *mouseOptions) {
		o.holdThreshold = threshold
	}
}
