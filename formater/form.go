package formater

import (
	"encoding/xml"
)

func FormatResponse(res SEnvelope) ([]byte, error) {
	return xml.Marshal(res)
}
