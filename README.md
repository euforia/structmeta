structmeta
==========
This library allows to parse and work with struct tags.

Example:
	
	import(
		"fmt"
		"github.com/euforia/structmeta"
	)
	
	type sample struct {
		Id int `tag:"id,primary_key"`
		Name string
	}

	var s sample
	sm := structmeta.ParseStructMetadata(&s, "tag", true)
	if fm := sm.FieldByName("Id");fm !=nil {
		fmt.Println(fm)
	}