/*
 * MIT License
 *
 * Copyright (c) [2020] [Semchenko Aleksandr]
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package kaspersky

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
)

//DpeKeyService Class Reference
//
//Interface for working with encrypted devices..
//
//List of all members.
type DpeKeyService service

//GetDeviceKeys3
//Returns information about host and key for chosen encrypted device.
//
//	Parameters:
//	- wstrDeviceId (string) Device id
//	Return:
//	- pKeyInfos (array) Array of params with key info. See Srvview encrypted devices on hosts.
//	The difference is that the key is decrypted.
func (di *DpeKeyService) GetDeviceKeys3(ctx context.Context, wstrDeviceId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrDeviceId": "%s"}`, wstrDeviceId))

	request, err := http.NewRequest("POST", di.client.Server+"/api/v1.0/DpeKeyService.GetDeviceKeys3", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := di.client.Do(ctx, request, nil)
	return raw, err
}
