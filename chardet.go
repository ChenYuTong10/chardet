package chardet

// HasPrefix tests whether different encoding byte slice begins with prefix.
// The byte slice has different encoding. So we can not use HasPrefix
// in package bytes implemented by string compare.
func HasPrefix(s, prefix []byte) bool {
	if len(s) < len(prefix) {
		return false
	}

	s = s[:len(prefix)]

	for i := 0; i < len(s); i++ {
		if s[i] != prefix[i] {
			return false
		}
	}
	return true
}

type Detector struct {
	Encoding string
}

// Feed feeds a chunk of a byte slice bs from a document through Detector.
// It can only distinguish encodings on Windows OS, such as ANSI, UTF-8,
// BOM UTF-8, UTF-16 BE, UTF-16 LE.
func (d *Detector) Feed(bs []byte) {

	if len(bs) == 0 {
		return
	}

	if HasPrefix(bs, BOM_UTF8_PREFIX) {
		// `EF BB BF` => `UTF-8 with BOM`
		d.Encoding = BOM_UTF8
	} else if HasPrefix(bs, BOM_UTF16_BE_PREFIX) {
		// `FE FF` => `UTF-16 BE`
		d.Encoding = BOM_UTF16_BE
	} else if HasPrefix(bs, BOM_UTF16_LE_PREFIX) {
		// `FF FE` => `UTF-16 LE`
		d.Encoding = BOM_UTF16_LE
	}

	if len(d.Encoding) != 0 {
		return
	}

	utf8Detector := new(UTF8Detector)
	utf8Detector.Feed(bs)

	if utf8Detector.State {
		d.Encoding = UTF8
	} else {
		// Obviously, a direct conclusion to ANSI is not correct.
		// But it is enough on Windows OS.
		d.Encoding = ANSI
	}
}

type UTF8Detector struct {
	State bool
}

// Feed feeds the byte slice through the UTF8 specific Detector.
// The Detector will check every bytes whether it accords with UTF8 features.
func (u *UTF8Detector) Feed(bs []byte) {

	wordState := false
	oneBitCount := 0
	continueByteCount := 0

	for i := 0; i < len(bs); i++ {
		// UTF8 stream has the following features:
		// · With the highest bit set to 0, it is a single byte value.
		// · With the highest two bit set to 10, it is a continuation byte value.
		// · If the byte sequence prefix is not either a single byte value or a continuation byte value, it will be a multipart byte value.
		//   The first byte of a multipart byte value will indicate how many continuation byte value in total for this word.
		//
		// In conclusion, we can classify a UTF8 stream only through the number of leading 1 bits in first byte.
		// · '\x65' is '0110 0101' written in binary. There is no bit leads to 1 before the high bit first leads to 0, so it is a single bit value.
		// · '\x88' is '1000 1000'. There is only one bit leads to 1 before the bit leads to 0 comes.
		// · '\xef' is '1110 0110'. There are 3 bits leads to 1.
		// So that is how to distinguish between UTF8 and other encodings.
		b := bs[i]
		oneBitCount = 0
		for b&HighestBitMask == HighestBitMask {
			oneBitCount++
			b <<= 1
		}

		switch oneBitCount {
		case 0:
			// a single byte value
			if wordState {
				u.State = false
				return
			}
		case 1:
			// a continuation byte value
			if !wordState {
				u.State = false
				return
			}
			continueByteCount--
			if continueByteCount == 0 {
				wordState = false
			}
		default:
			// a multipart byte value
			continueByteCount = oneBitCount
			continueByteCount--
			wordState = true
		}
	}

	u.State = true
}
