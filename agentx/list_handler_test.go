// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package agentx_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go-snmp-agentx/agentx"
	"go-snmp-agentx/agentx/pdu"
	"go-snmp-agentx/agentx/value"
)

func TestListHandler(t *testing.T) {
	e := setUpTestEnvironment(t)
	defer e.tearDown()

	session, err := e.client.Session()
	require.NoError(t, err)
	defer session.Close()

	lh := &agentx.ListHandler{}
	i := lh.Add("1.3.6.1.4.1.45995.3.1")
	i.Type = pdu.VariableTypeOctetString
	i.Value = func() interface{} {
		return "test"
	}
	session.Handler = lh

	baseOID := value.MustParseOID("1.3.6.1.4.1.45995")

	require.NoError(t, session.Register(127, baseOID))
	defer session.Unregister(127, baseOID)

	t.Run("Get", func(t *testing.T) {
		assert.Equal(t,
			".1.3.6.1.4.1.45995.3.1 = STRING: \"test\"",
			SNMPGet(t, "1.3.6.1.4.1.45995.3.1"))

		assert.Equal(t,
			".1.3.6.1.4.1.45995.3.2 = No Such Object available on this agent at this OID",
			SNMPGet(t, "1.3.6.1.4.1.45995.3.2"))
	})

	t.Run("GetNext", func(t *testing.T) {
		assert.Equal(t,
			".1.3.6.1.4.1.45995.3.1 = STRING: \"test\"",
			SNMPGetNext(t, "1.3.6.1.4.1.45995.3.0"))

		assert.Equal(t,
			".1.3.6.1.4.1.45995.3.1 = STRING: \"test\"",
			SNMPGetNext(t, "1.3.6.1.4.1.45995.3"))

	})

	t.Run("GetBulk", func(t *testing.T) {
		assert.Equal(t,
			".1.3.6.1.4.1.45995.3.1 = STRING: \"test\"",
			SNMPGetBulk(t, "1.3.6.1.4.1.45995.3.0", 0, 1))
	})
}
