package pb

func ModSplit(mod Mod) []Mod {
	slice := []Mod{}
	i32 := int32(mod)

	for key := range Mod_name {
		if i32&key > 0 {
			slice = append(slice, Mod(key))
		}
	}

	return slice
}

func ModMerge(values ...Mod) Mod {
	var merge Mod

	for _, value := range values {
		merge |= value
	}

	return merge
}
