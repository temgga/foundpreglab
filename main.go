func (h *Heap) bubbleDown(index int) {
	data := h.data
	cmp := h.comparator
	n := len(data)

	for {
		left := 2*index + 1
		if left >= n {
			break
		}

		smallestChild := left
		if right := left + 1; right < n && cmp(data[right], data[left]) < 0 {
			smallestChild = right
		}

		if cmp(data[smallestChild], data[index]) >= 0 {
			break
		}

		data[index], data[smallestChild] = data[smallestChild], data[index]
		index = smallestChild
	}
}