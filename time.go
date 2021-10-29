package msgpack

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"time"

	"git.gostudent.de/pkg/log/errors"
	"github.com/vmihailenco/msgpack/v5/msgpcode"
)

const millisec = 1000000

var timeExtID int8 = -1

func init() {
	RegisterExtEncoder(timeExtID, time.Time{}, timeEncoder)
	RegisterExtDecoder(timeExtID, time.Time{}, timeDecoder)

	RegisterExtEncoder(13, (*time.Time)(nil), func(e *Encoder, v reflect.Value) ([]byte, error) {
		t := v.Interface().(*time.Time)
		if t == nil {
			return nil, nil
		}
		return marshalTime(*t)
	})
	RegisterExtDecoder(13, (*time.Time)(nil), func(d *Decoder, v reflect.Value, extLen int) error {
		b := make([]byte, extLen)
		_, err := d.Buffered().Read(b)
		if err != nil {
			return err
		}

		return unmarshalTime(b, v.Interface().(*time.Time))
	})
}

func marshalTime(d time.Time) ([]byte, error) {
	return Marshal(d.UnixNano() / millisec)
}

func unmarshalTime(data []byte, d *time.Time) error {
	if len(data) == 0 {
		return nil
	}

	var v interface{}
	if err := Unmarshal(data, &v); err != nil {
		return err
	}

	switch val := v.(type) {
	case float64:
		*d = time.Unix(0, int64(val*millisec))
	case int64:
		*d = time.Unix(val/1000, (val%1000)*millisec)
	case string:
		t, err := time.Parse(time.RFC3339Nano, val)
		if err != nil {
			return errors.Wrap(err, errors.Msg("unmarshalTime: string layout not implemented"),
				errors.Fields{"value": val})
		}
		*d = t
	case *time.Time:
		*d = *val
	default:
		return errors.New("unmarshalTime: unimplemented type", errors.Fields{
			"type":      reflect.TypeOf(v).String(),
			"value":     string(data),
			"value_hex": fmt.Sprintf("%x", data),
		})
	}

	return nil
}

func timeEncoder(e *Encoder, v reflect.Value) ([]byte, error) {
	return e.encodeTime(v.Interface().(time.Time)), nil
}

func timeDecoder(d *Decoder, v reflect.Value, extLen int) error {
	tm, err := d.decodeTime(extLen)
	if err != nil {
		return err
	}

	ptr := v.Addr().Interface().(*time.Time)
	*ptr = tm

	return nil
}

func (e *Encoder) EncodeTime(tm time.Time) error {
	b := e.encodeTime(tm)
	if err := e.encodeExtLen(len(b)); err != nil {
		return err
	}
	if err := e.w.WriteByte(byte(timeExtID)); err != nil {
		return err
	}
	return e.write(b)
}

func (e *Encoder) encodeTime(tm time.Time) []byte {
	if e.timeBuf == nil {
		e.timeBuf = make([]byte, 12)
	}

	secs := uint64(tm.Unix())
	if secs>>34 == 0 {
		data := uint64(tm.Nanosecond())<<34 | secs

		if data&0xffffffff00000000 == 0 {
			b := e.timeBuf[:4]
			binary.BigEndian.PutUint32(b, uint32(data))
			return b
		}

		b := e.timeBuf[:8]
		binary.BigEndian.PutUint64(b, data)
		return b
	}

	b := e.timeBuf[:12]
	binary.BigEndian.PutUint32(b, uint32(tm.Nanosecond()))
	binary.BigEndian.PutUint64(b[4:], secs)
	return b
}

func (d *Decoder) DecodeTime() (time.Time, error) {
	c, err := d.readCode()
	if err != nil {
		return time.Time{}, err
	}

	// Legacy format.
	if c == msgpcode.FixedArrayLow|2 {
		sec, err := d.DecodeInt64()
		if err != nil {
			return time.Time{}, err
		}

		nsec, err := d.DecodeInt64()
		if err != nil {
			return time.Time{}, err
		}

		return time.Unix(sec, nsec), nil
	}

	if msgpcode.IsString(c) {
		s, err := d.string(c)
		if err != nil {
			return time.Time{}, err
		}
		return time.Parse(time.RFC3339Nano, s)
	}

	extID, extLen, err := d.extHeader(c)
	if err != nil {
		return time.Time{}, err
	}

	// NodeJS seems to use extID 13.
	if extID != timeExtID && extID != 13 {
		return time.Time{}, errors.Errorf("msgpack: invalid time ext id=%d", extID)
	}

	tm, err := d.decodeTime(extLen)
	if err != nil {
		return tm, err
	}

	if tm.IsZero() {
		// Zero time does not have timezone information.
		return tm.UTC(), nil
	}
	return tm, nil
}

func (d *Decoder) decodeTime(extLen int) (time.Time, error) {
	b, err := d.readN(extLen)
	if err != nil {
		return time.Time{}, err
	}

	switch len(b) {
	case 4:
		sec := binary.BigEndian.Uint32(b)
		return time.Unix(int64(sec), 0), nil
	case 8:
		sec := binary.BigEndian.Uint64(b)
		nsec := int64(sec >> 34)
		sec &= 0x00000003ffffffff
		return time.Unix(int64(sec), nsec), nil
	case 12:
		nsec := binary.BigEndian.Uint32(b)
		sec := binary.BigEndian.Uint64(b[4:])
		return time.Unix(int64(sec), int64(nsec)), nil
	default:
		err = errors.Errorf("msgpack: invalid ext len=%d decoding time", extLen)
		return time.Time{}, err
	}
}
