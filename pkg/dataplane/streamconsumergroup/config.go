/*
Copyright 2019 Iguazio Systems Ltd.

Licensed under the Apache License, Version 2.0 (the "License") with
an addition restriction as set forth herein. You may not use this
file except in compliance with the License. You may obtain a copy of
the License at http://www.apache.org/licenses/LICENSE-2.0.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
implied. See the License for the specific language governing
permissions and limitations under the License.

In addition, you may not use the software for any purposes that are
illegal under applicable law, and the grant of the foregoing license
under the Apache 2.0 license is conditioned upon your compliance with
such restriction.
*/
package streamconsumergroup

import (
	"time"

	"github.com/v3io/v3io-go/pkg/common"
	"github.com/v3io/v3io-go/pkg/dataplane"
)

type Config struct {
	Session struct {
		Timeout           time.Duration `json:"timeout,omitempty"`
		HeartbeatInterval time.Duration
	} `json:"session,omitempty"`
	State struct {
		ModifyRetry struct {
			Attempts int            `json:"attempts,omitempty"`
			Backoff  common.Backoff `json:"backoff,omitempty"`
		} `json:"modifyRetry,omitempty"`
	} `json:"state,omitempty"`
	SequenceNumber struct {
		CommitInterval    time.Duration `json:"commitInterval,omitempty"`
		ShardWaitInterval time.Duration `json:"shardWaitInterval,omitempty"`
	}
	Claim struct {
		RecordBatchChanSize int `json:"recordBatchChanSize,omitempty"`
		RecordBatchFetch    struct {
			Interval          time.Duration           `json:"interval,omitempty"`
			NumRecordsInBatch int                     `json:"numRecordsInBatch,omitempty"`
			InitialLocation   v3io.SeekShardInputType `json:"initialLocation,omitempty"`
		} `json:"recordBatchFetch,omitempty"`
		GetShardLocationRetry struct {
			Attempts int            `json:"attempts,omitempty"`
			Backoff  common.Backoff `json:"backoff,omitempty"`
		} `json:"getShardLocationRetry,omitempty"`
	} `json:"claim,omitempty"`
}

// NewConfig returns a new configuration instance with sane defaults.
func NewConfig() *Config {
	c := &Config{}
	c.Session.Timeout = 10 * time.Second
	c.Session.HeartbeatInterval = 3 * time.Second
	c.State.ModifyRetry.Attempts = 100
	c.State.ModifyRetry.Backoff = common.Backoff{
		Min:    50 * time.Millisecond,
		Max:    1 * time.Second,
		Factor: 4,
	}
	c.SequenceNumber.CommitInterval = 10 * time.Second
	c.SequenceNumber.ShardWaitInterval = 1 * time.Second
	c.Claim.RecordBatchChanSize = 100
	c.Claim.RecordBatchFetch.Interval = 250 * time.Millisecond
	c.Claim.RecordBatchFetch.NumRecordsInBatch = 10
	c.Claim.RecordBatchFetch.InitialLocation = v3io.SeekShardInputTypeEarliest
	c.Claim.GetShardLocationRetry.Attempts = 100
	c.Claim.GetShardLocationRetry.Backoff = common.Backoff{
		Min:    50 * time.Millisecond,
		Max:    1 * time.Second,
		Factor: 2,
	}

	return c
}
