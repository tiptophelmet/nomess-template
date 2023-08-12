package intl

var lz *Localizer

func Init(defaultLocale string) *Localizer {
	if lz != nil {
		return lz
	}

	lz = &Localizer{locale: defaultLocale}

	return lz
}

type Localizer struct {
	locale string
}

func SetLocale(locale string) {
	lz.locale = locale
}

func GetLocale() string {
	return lz.locale
}
