import os

files = [
    "tests/integration/ecommerce_flow_test.go",
    "tests/integration/ai_agent_transaction_test.go",
    "tests/integration/circuit_breaker_test.go",
    "tests/integration/cross_module_compliance_test.go"
]

for file_path in files:
    with open(file_path, "r") as f:
        content = f.read()

    # Add import
    if 'github.com/cosmos/cosmos-sdk/codec/address' not in content:
        content = content.replace(
            '\t"github.com/cosmos/cosmos-sdk/codec"',
            '\t"github.com/cosmos/cosmos-sdk/codec"\n\t"github.com/cosmos/cosmos-sdk/codec/address"'
        )

    # Remove banktypes.TransientKey usage
    content = content.replace(
        '\t// Create transient store keys\n\ttKeys := storetypes.NewTransientStoreKeys(banktypes.TransientKey)\n\n',
        ''
    )
    content = content.replace(
        '\tfor _, tKey := range tKeys {\n\t\tstateStore.MountStoreWithDB(tKey, storetypes.StoreTypeTransient, db)\n\t}\n',
        ''
    )
    # Remove usage in DefaultContextWithDB
    content = content.replace(
        'testutil.DefaultContextWithDB(s.T(), storeKeys[authtypes.StoreKey], tKeys[banktypes.TransientKey])',
        'testutil.DefaultContextWithDB(s.T(), storeKeys[authtypes.StoreKey], storetypes.NewTransientStoreKey("transient_test"))'
    )

    # Update authkeeper.NewAccountKeeper
    # We need to inject the address codec.
    old_call = """s.accountKeeper = authkeeper.NewAccountKeeper(
		s.cdc,
		runtime.NewKVStoreService(storeKeys[authtypes.StoreKey]),
		authtypes.ProtoBaseAccount,
		maccPerms,
		"stateset",
		s.authority.String(),
	)"""
    
    new_call = """s.accountKeeper = authkeeper.NewAccountKeeper(
		s.cdc,
		runtime.NewKVStoreService(storeKeys[authtypes.StoreKey]),
		authtypes.ProtoBaseAccount,
		maccPerms,
		address.NewBech32Codec("stateset"),
		"stateset",
		s.authority.String(),
	)"""

    content = content.replace(old_call, new_call)

    with open(file_path, "w") as f:
        f.write(content)

print("Files updated.")
