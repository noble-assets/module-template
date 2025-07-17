#!/bin/bash
SIMD="./simapp/build/simd"
HOME_DIR=".template"
TEMP=$HOME_DIR/genesis.json

for arg in "$@"; do
	case $arg in
	-r | --reset)
		rm -rf $HOME_DIR
		shift
		;;
	esac
done

if ! [ -f .orbiter/data/priv_validator_state.json ]; then
	$SIMD init validator --chain-id "orbiter-1" --home $HOME_DIR &>/dev/null

	$SIMD keys add validator --home .orbiter --keyring-backend test &>/dev/null
	$SIMD genesis add-genesis-account validator 2000000ustake,1000000000uusdc --home $HOME_DIR --keyring-backend test

	touch $TEMP && jq '.app_state.bank.denom_metadata += [{ "description": "Circle USD Coin", "denom_units": [{ "denom": "uusdc", "exponent": 0, "aliases": ["microusdc"] }, { "denom": "usdc", "exponent": 6 }], "base": "uusdc", "display": "usdc", "name": "Circle USD Coin", "symbol": "USDC" }]' $HOME_DIR/config/genesis.json >$TEMP && mv $TEMP $HOME_DIR/config/genesis.json
	touch $TEMP && jq '.app_state.staking.params.bond_denom = "ustake"' $HOME_DIR/config/genesis.json >$TEMP && mv $TEMP $HOME_DIR/config/genesis.json

	$SIMD genesis gentx validator 1000000ustake --chain-id "orbiter-1" --home $HOME_DIR --keyring-backend test &>/dev/null
	$SIMD genesis collect-gentxs --home $HOME_DIR &>/dev/null

	sed -i '' 's/timeout_commit = "5s"/timeout_commit = "1s"/g' $HOME_DIR/config/config.toml
fi

$SIMD start --home $HOME_DIR
