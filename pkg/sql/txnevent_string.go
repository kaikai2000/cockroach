// Code generated by "stringer -type=txnEvent"; DO NOT EDIT.

package sql

import "strconv"

const _txnEvent_name = "noEventtxnStarttxnCommittxnAbortedtxnRestart"

var _txnEvent_index = [...]uint8{0, 7, 15, 24, 34, 44}

func (i txnEvent) String() string {
	if i < 0 || i >= txnEvent(len(_txnEvent_index)-1) {
		return "txnEvent(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _txnEvent_name[_txnEvent_index[i]:_txnEvent_index[i+1]]
}
