package aggregate

import (
	"testing"

	"github.com/apache/incubator-pegasus/go-client/idl/base"
	"github.com/stretchr/testify/assert"
)

func TestDecodePartitionPerfCounter(t *testing.T) {
	tests := []struct {
		name string

		isNil          bool
		counterName    string
		appID          int32
		partitionIndex int32
	}{
		{name: "replica*app.pegasus*get_latency@2.5.p999", isNil: true},

		// server-level counter, does not contain gpid.
		{name: "replica*eon.replica*table.level.RPC_RRDB_RRDB_CHECK_AND_MUTATE.latency(ns)@temp", isNil: true},
		{
			name:  "replica*eon.replica*table.level.RPC_RRDB_RRDB_MULTI_PUT.latency(ns)@temp.p999",
			isNil: true,
		},

		{
			name:           "replica*app.pegasus*recent.abnormal.count@1.2",
			counterName:    "replica*app.pegasus*recent.abnormal.count",
			appID:          1,
			partitionIndex: 2,
		},
	}

	for _, tt := range tests {
		pc := decodePartitionPerfCounter(tt.name, 1.0)
		assert.Equal(t, pc == nil, tt.isNil, tt.name)
		if pc != nil {
			assert.Equal(t, pc.name, tt.counterName)
			assert.Equal(t, pc.gpid, base.Gpid{Appid: tt.appID, PartitionIndex: tt.partitionIndex})
		}
	}
}
