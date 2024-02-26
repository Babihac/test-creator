package utils

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

func ParseUUID(uuid pgtype.UUID) (string, error) {
	if !uuid.Valid {
		return "", errors.New("invalid UUID")
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid.Bytes[:4], uuid.Bytes[4:6], uuid.Bytes[6:8], uuid.Bytes[8:10], uuid.Bytes[10:]), nil
}

func StringToUUID(s string) pgtype.UUID {
	data, err := parseUUID(s)
	if err != nil {
		return pgtype.UUID{
			Bytes: [16]byte{},
			Valid: false,
		}
	}
	return pgtype.UUID{
		Bytes: data,
		Valid: true,
	}
}

func parseUUID(src string) (dst [16]byte, err error) {
	switch len(src) {
	case 36:
		src = src[0:8] + src[9:13] + src[14:18] + src[19:23] + src[24:]
	case 32:
	default:
		return dst, fmt.Errorf("cannot parse UUID %v", src)
	}

	buf, err := hex.DecodeString(src)
	if err != nil {
		return dst, err
	}

	copy(dst[:], buf)
	return dst, err
}
