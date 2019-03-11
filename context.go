package mid

type Context struct {
	mid *Mid
    index int8
	Keys map[string]interface{}
    Errors []error
}

