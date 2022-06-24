package dh_key

import "github.com/monnand/dhkx"

func NewKeyExchange() (dhKeyExch *dhKey, err error) {
	var (
		g       *dhkx.DHGroup
		privKey *dhkx.DHKey
	)
	if g, err = dhkx.GetGroup(0); err != nil {
		return
	}
	dhKeyExch = new(dhKey)
	dhKeyExch.group = g
	if privKey, err = g.GeneratePrivateKey(nil); err != nil {
		return
	}
	dhKeyExch.group = g
	dhKeyExch.privKey = privKey
	return dhKeyExch, nil
}
