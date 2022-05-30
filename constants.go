package chardet

var (
	// UTF-8
	BOM_UTF8_PREFIX = []byte("\xef\xbb\xbf")

	// UTF-16 big endian
	BOM_UTF16_BE_PREFIX = []byte("\xfe\xff")

	// UTF-16 little endian
	BOM_UTF16_LE_PREFIX = []byte("\xff\xfe")
)

const (
	// Bit mask
	HighestBitMask = byte(0x80)

	// Encoding
	BOM_UTF8 = "BOM UTF-8"

	BOM_UTF16_BE = "UTF-16 BE"

	BOM_UTF16_LE = "UTF-16 LE"

	UTF8 = "UTF-8"

	ANSI = "ANSI"
)
