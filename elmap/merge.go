package elmap

import "strings"

// MergeStringMap merge src to dst map
func MergeSSMap(src, dst map[string]string, ignoreCase bool) {
	for k, v := range src {
		if ignoreCase {
			k = strings.ToLower(k)
		}
		dst[k] = v
	}
}

// MergeStringMap merge src to dst map
func MergeSIMap(src, dst map[string]int, ignoreCase bool) {
	for k, v := range src {
		if ignoreCase {
			k = strings.ToLower(k)
		}
		dst[k] = v
	}
}

// MergeStringMap merge src to dst map
func MergeSUMap(src, dst map[string]uint, ignoreCase bool) {
	for k, v := range src {
		if ignoreCase {
			k = strings.ToLower(k)
		}
		dst[k] = v
	}
}

// MergeStringMap merge src to dst map
func MergeSFMap(src, dst map[string]float64, ignoreCase bool) {
	for k, v := range src {
		if ignoreCase {
			k = strings.ToLower(k)
		}
		dst[k] = v
	}
}

// MergeStringMap merge src to dst map
func MergeIIMap(src, dst map[int]int) {
	for k, v := range src {
		dst[k] = v
	}
}

// MergeStringMap merge src to dst map
func MergeIUMap(src, dst map[int]uint) {
	for k, v := range src {
		dst[k] = v
	}
}

// MergeStringMap merge src to dst map
func MergeISMap(src, dst map[int]string) {
	for k, v := range src {
		dst[k] = v
	}
}

// MergeStringMap merge src to dst map
func MergeIFMap(src, dst map[int]float64) {
	for k, v := range src {
		dst[k] = v
	}
}

// MergeStringMap merge src to dst map
func MergeUIMap(src, dst map[uint]int) {
	for k, v := range src {
		dst[k] = v
	}
}

// MergeStringMap merge src to dst map
func MergeUUMap(src, dst map[uint]uint) {
	for k, v := range src {
		dst[k] = v
	}
}

// MergeStringMap merge src to dst map
func MergeUSMap(src, dst map[uint]string) {
	for k, v := range src {
		dst[k] = v
	}
}

// MergeStringMap merge src to dst map
func MergeUFMap(src, dst map[uint]float64) {
	for k, v := range src {
		dst[k] = v
	}
}
