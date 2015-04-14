/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          uuid.go
 * Description:   UUID Tools
 */

package utils

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
)

// UUID() provides unique identifier strings.
func UUID() string {
	b := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		log.Fatal(err)
	}
	b[6] = (b[6] & 0x0F) | 0x40
	b[8] = (b[8] &^ 0x40) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[:4], b[4:6], b[6:8], b[8:10], b[10:])
}
