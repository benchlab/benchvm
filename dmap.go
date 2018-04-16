package benchvm

type detMap struct {
	data  map[interface{}]interface{}
	array []interface{}
}

func (d *detMap) Size() int {
	return len(d.array)
}

func (d *detMap) Assign(key, value interface{}) {
	d.data[key] = value
}

func (d *detMap) Retrieve(key interface{}) interface{} {
	return d.data[key]
}
