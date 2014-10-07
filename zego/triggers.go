package zego

func (a Auth) ListTriggers() *Resource {

	path := "/triggers.json"
	resource := api(a, "GET", path, "")

	return resource

}

func (a Auth) ListActiveTriggers() *Resource {

	path := "/triggers/active.json"
	resource := api(a, "GET", path, "")

	return resource

}
