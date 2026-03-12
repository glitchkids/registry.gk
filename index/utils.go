package index

import (
	"fmt"
	"slices"
)

func FindRegistryItemByName(registryIndex RegistryIndex, name string) (RegistryIndexItem, error) {
	index := slices.IndexFunc(registryIndex, func(e RegistryIndexItem) bool {
		return e.Name == name
	})

	if index == -1 {
		return RegistryIndexItem{}, fmt.Errorf("Registry index item not found : %v", name)
	}

	return registryIndex[index], nil
}
