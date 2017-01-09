package save

type Jsoner interface {
	Json(expand bool) (string, map[string]interface{})
}
