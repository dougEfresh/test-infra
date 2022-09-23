package nodekit

import (
	"fmt"
	"net"

	"github.com/celestiaorg/celestia-node/logs"
	"github.com/celestiaorg/celestia-node/nodebuilder"
	"github.com/celestiaorg/celestia-node/nodebuilder/node"
	"github.com/celestiaorg/celestia-node/params"
	logging "github.com/ipfs/go-log/v2"

	"go.uber.org/fx"
)

func NewNode(path string, tp node.Type, IP net.IP, trustedHash string, options ...fx.Option) (*nodebuilder.Node, error) {
	cfg := nodebuilder.DefaultConfig(tp)
	cfg.P2P.ListenAddresses = []string{fmt.Sprintf("/ip4/%s/tcp/2121", IP)}
	cfg.Header.TrustedHash = trustedHash

	err := nodebuilder.Init(*cfg, path, tp)
	if err != nil {
		return nil, err
	}
	store, err := nodebuilder.OpenStore(path)
	if err != nil {
		return nil, err
	}

	options = append([]fx.Option{nodebuilder.WithNetwork(params.Private)}, options...)
	return nodebuilder.New(tp, store, options...)
}

func SetLoggersLevel(lvl string) error {
	level, err := logging.LevelFromString(lvl)
	if err != nil {
		return err
	}
	logs.SetAllLoggers(level)

	return nil
}
