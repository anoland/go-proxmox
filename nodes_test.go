package proxmox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodes(t *testing.T) {
	client := ClientFromLogins()
	nodes, err := client.Nodes()
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(nodes), 1)
	for _, n := range nodes {
		assert.NotEmpty(t, n.Node)
		var node *Node
		t.Run("get status for node "+n.Node, func(t *testing.T) {
			var err error
			node, err = client.Node(n.Node)
			assert.Nil(t, err)
			assert.Equal(t, n.MaxMem, node.Memory.Total)
			assert.Equal(t, n.Disk, node.RootFS.Used)
		})

		t.Run("get VMs for node "+n.Node, func(t *testing.T) {
			_, err := node.VirtualMachines()
			assert.Nil(t, err)
		})

		break // only pull status from one node
	}

	_, err = client.Node("doesnt-exist")
	assert.Contains(t, err.Error(), "500 hostname lookup 'doesnt-exist' failed - failed to get address info for: doesnt-exist:")
}
