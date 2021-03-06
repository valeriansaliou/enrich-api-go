// Copyright 2017 Valerian Saliou. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package enrich


import (
  "fmt"
  "net/url"
)


// VerifyService service
type VerifyService service


// ValidateEmailData mapping
type ValidateEmailData struct {
  Valid     *bool                  `json:"valid,omitempty"`
  Accuracy  *float32               `json:"accuracy,omitempty"`
  Results   *ValidateEmailResults  `json:"results,omitempty"`
}

// ValidateEmailResults mapping
type ValidateEmailResults struct {
  Gravatar     *bool  `json:"gravatar,omitempty"`
  Gibberish    *bool  `json:"gibberish,omitempty"`
  Disposable   *bool  `json:"disposable,omitempty"`
  Webmail      *bool  `json:"webmail,omitempty"`
  MXRecords    *bool  `json:"mx_records,omitempty"`
  SMTPServer   *bool  `json:"smtp_server,omitempty"`
  SMTPCheck    *bool  `json:"smtp_check,omitempty"`
  SPFPolicy    *bool  `json:"spf_policy,omitempty"`
  DMARCPolicy  *bool  `json:"dmarc_policy,omitempty"`
  CatchAll     *bool  `json:"catch_all,omitempty"`
  HighVolume   *bool  `json:"high_volume,omitempty"`
}


// String returns the string representation of ValidateEmailData
func (instance ValidateEmailData) String() string {
  return Stringify(instance)
}

// String returns the string representation of ValidateEmailResults
func (instance ValidateEmailResults) String() string {
  return Stringify(instance)
}


// ValidateEmail verifies if an email is valid and if it exists.
func (service *VerifyService) ValidateEmail(email string) (*ValidateEmailData, *Response, error) {
  url := fmt.Sprintf("verify/validate/email?email=%s", url.QueryEscape(email))
  req, _ := service.client.NewRequest("GET", url, nil)

  data := new(ValidateEmailData)
  resp, err := service.client.Do(req, data)
  if err != nil {
    return nil, resp, err
  }

  return data, resp, err
}
