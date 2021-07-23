package util

import "github.com/seiflotfy/cuckoofilter"

type CuckooFilter struct {
	filter *cuckoo.Filter
}

func NewCuckooFilter() *CuckooFilter {
	return &CuckooFilter{filter: cuckoo.NewFilter(10000)}
}

func (c *CuckooFilter) DeleteItem(item string) bool {
	return c.filter.Delete([]byte(item))
}

func (c *CuckooFilter) InsertUniqueItem(item string) bool {
	return c.filter.InsertUnique([]byte(item))
}

func (c *CuckooFilter) LookupItem(item string) bool {
	return c.filter.Lookup([]byte(item))
}

func (c *CuckooFilter) CountItem() uint {
	return c.filter.Count()
}

func (c *CuckooFilter) ResetItems() {
	c.filter.Reset()
}
