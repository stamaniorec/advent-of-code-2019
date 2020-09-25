package permutation

type Permutation []int

func GeneratePermutations(items []int) (perms []Permutation) {
	permute(items, &perms, 0, len(items))
	return
}

func permute(items []int, perms *[]Permutation, low int, high int) {
	if low == high {
		p := make([]int, len(items))
		copy(p, items)

		*perms = append(*perms, p)
		return
	}

	for i := range items {
		items[i], items[low] = items[low], items[i]
		permute(items, perms, low+1, high)
		items[i], items[low] = items[low], items[i]
	}
}
