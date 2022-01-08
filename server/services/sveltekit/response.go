/**
 * File: response.go
 * Project: sveltekit
 * File Created: 08 Jan 2022 15:44:43
 * Author: und3fined (me@und3fined.com)
 * -----
 * Last Modified: 08 Jan 2022 16:47:07
 * Modified By: und3fined (me@und3fined.com)
 * -----
 * Copyright (c) 2022 und3fined.com
 */
package sveltekit

import "encoding/json"

type Response struct {
	Status  uint16            `json:"status"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

func parseResponse(respStr string) (*Response, error) {
	var resp Response
	err := json.Unmarshal([]byte(respStr), &resp)
	return &resp, err
}
