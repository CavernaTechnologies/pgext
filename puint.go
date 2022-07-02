package pgext

import (
	"encoding/binary"
	"fmt"
	"strconv"

	"github.com/jackc/pgio"
	"github.com/jackc/pgtype"
)

// In order to send uint64 over binary wire protocol, it must implement special encoding and decoding methods
// The encoding and decoding process is non-trivial and is taken from
// https://github.com/jackc/pgtype/blob/master/numeric.go

type Puint uint64

func (p *Puint) DecodeText(ci *pgtype.ConnInfo, src []byte) error {
	if src == nil {
		return fmt.Errorf("invalid value of puint")
	} else if string(src) == "NaN" {
		return fmt.Errorf("invalid value of puint")
	} else if string(src) == "Infinity" {
		return fmt.Errorf("invalid value of puint")
	} else if string(src) == "-Infinity" {
		return fmt.Errorf("invalid value of puint")
	}

	v, err := strconv.ParseUint(string(src), 10, 64)
	if err != nil {
		return err
	}
	*p = Puint(v)

	return nil
}

func (p *Puint) DecodeBinary(ci *pgtype.ConnInfo, src []byte) error {
	if src == nil {
		*p = 0
		return nil
	}

	if len(src) < 8 {
		return fmt.Errorf("numeric incomplete %v", src)
	}

	rp := 0
	ndigits := binary.BigEndian.Uint16(src[rp:])
	rp += 2
	weight := int16(binary.BigEndian.Uint16(src[rp:]))
	rp += 2
	sign := binary.BigEndian.Uint16(src[rp:])
	rp += 2
	dscale := int16(binary.BigEndian.Uint16(src[rp:]))
	rp += 2

	if sign != 0 {
		return fmt.Errorf("invalid uint")
	}

	if dscale != 0 {
		return fmt.Errorf("invalid uint")
	}

	if int(weight) != int(ndigits)-1 {
		return fmt.Errorf("invalid uint")
	}

	if ndigits == 0 {
		*p = 0
		return nil
	}

	if len(src[rp:]) < int(ndigits)*2 {
		return fmt.Errorf("numeric incomplete %v", src)
	}

	digits := int(ndigits)
	var accum uint64 = 0

	for i := 0; i < digits; i++ {
		if i > 0 {
			accum *= 10000
		}
		accum += uint64(binary.BigEndian.Uint16(src[rp:]))
		rp += 2
	}

	*p = Puint(accum)
	return nil
}

func (p Puint) EncodeText(ci *pgtype.ConnInfo, buf []byte) (newBuf []byte, err error) {
	s := strconv.FormatUint(uint64(p), 10)
	buf = append(buf, s...)
	return buf, nil
}

func (p Puint) EncodeBinary(ci *pgtype.ConnInfo, buf []byte) (newBuf []byte, err error) {
	v := uint64(p)

	var wholeDigits []int16

	for v != 0 {
		r := v % 10000
		v = v / 10000
		wholeDigits = append(wholeDigits, int16(r))
	}

	// number of digits
	buf = pgio.AppendInt16(buf, int16(len(wholeDigits)))
	// weight of whole number
	weight := int16(len(wholeDigits) - 1)
	buf = pgio.AppendInt16(buf, weight)
	// sign
	buf = pgio.AppendInt16(buf, 0)
	// decimal scale
	buf = pgio.AppendInt16(buf, 0)
	// add whole number
	for i := len(wholeDigits) - 1; i >= 0; i-- {
		buf = pgio.AppendInt16(buf, wholeDigits[i])
	}

	return buf, nil
}
