package models

import "fmt"

func CollectionKey(aid string, ct CollectionType) string {
	return fmt.Sprintf("%s:%d", aid, ct)
}
