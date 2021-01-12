package cosmostxdecoder_test

import (
	"encoding/base64"
	"testing"

	"github.com/calvinlauco/cosmostxdecoder"
	"github.com/stretchr/testify/assert"
)

func TestConsensusNodeAddressFromTmPubKey(t *testing.T) {
	prefix := "tcrocnclcons"
	tmpubkey, err := base64.StdEncoding.DecodeString("Og8ZfQTHFgTBGD5qoyo5NpyJCJRddC+WuSPtyZtlE7E=")
	assert.NoError(t, err)
	consensusNodeAddress, err := cosmostxdecoder.ConsensusNodeAddressFromTmPubKey(prefix, tmpubkey)
	assert.NoError(t, err)
	expected := "tcrocnclcons1wqajdt4qseasx4e8r8fz7juwdkfu4quvt9e87u"
	assert.Equal(t, expected, consensusNodeAddress)
}

func TestConsensusNodeAddressFromPubKey(t *testing.T) {
	prefix := "tcrocnclcons"
	consensusNodeAddress, err := cosmostxdecoder.ConsensusNodeAddressFromPubKey(prefix, "tcrocnclconspub1zcjduepq8g83jlgycutqfsgc8e42x23ex6wgjzy5t46zl94ey0kunxm9zwcsuzkxpr")
	assert.NoError(t, err)
	expected := "tcrocnclcons1wqajdt4qseasx4e8r8fz7juwdkfu4quvt9e87u"
	assert.Equal(t, expected, consensusNodeAddress)
}
