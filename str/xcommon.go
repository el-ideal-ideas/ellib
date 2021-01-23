// Copy from https://github.com/huandu/xstrings

package str

const bufferMaxInitGrowSize = 2048

// Lazy initialize a buffer.
func allocBuffer(orig, cur string) *stringBuilder {
	output := &stringBuilder{}
	maxSize := len(orig) * 4

	// Avoid to reserve too much memory at once.
	if maxSize > bufferMaxInitGrowSize {
		maxSize = bufferMaxInitGrowSize
	}

	output.Grow(maxSize)
	output.WriteString(orig[:len(orig)-len(cur)])
	return output
}
