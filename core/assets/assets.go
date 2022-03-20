package assets

// Values holds path for the assets' folder.
type Values struct {
	RelativePath string
	Root         string
	HTMLGlob     string
}

func (v Values) GetAssets() (string, string) {
	return v.RelativePath, v.Root
}

func (v Values) GetHTMLGlob() string {
	return v.HTMLGlob
}
