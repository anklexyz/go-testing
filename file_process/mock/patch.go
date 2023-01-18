package mock

//// Patch replaces a function with another
//func Patch(traget, replacement interface{}) *PatchGuard {
//	t := reflect.ValueOf(target)
//	r := reflect.ValueOf(replacement)
//	patchValue(t, r)
//
//	return &PatchGuard{t, r}
//}
//
//// Unpatch removes any monkey patches on target
//// return whether target was patched in the first place
//func Unpatch(target interface{}) bool {
//	return unpatchValue(reflect.ValueOf(target))
//}
