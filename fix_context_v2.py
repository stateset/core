import os
import re

files = [
    "tests/integration/ecommerce_flow_test.go",
    "tests/integration/ai_agent_transaction_test.go",
    "tests/integration/circuit_breaker_test.go",
    "tests/integration/cross_module_compliance_test.go",
    "tests/integration/stablecoin_lifecycle_test.go"
]

state_store_block = """	// Create multistore
	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	for _, key := range storeKeys {
		stateStore.MountStoreWithDB(key, storetypes.StoreTypeIAVL, db)
	}
	for _, tKey := range tKeys {
		stateStore.MountStoreWithDB(tKey, storetypes.StoreTypeTransient, db)
	}
	require.NoError(s.T(), stateStore.LoadLatestVersion())"""

# Note: previous script removed tKeys loops, so I should be careful matching.
# Let's match simpler patterns or just lines.

for file_path in files:
    if not os.path.exists(file_path):
        continue
        
    with open(file_path, "r") as f:
        content = f.read()

    # Extract keys from NewKVStoreKeys
    match = re.search(r'storeKeys\s*:=\s*storetypes.NewKVStoreKeys\((.*?)\)', content, re.DOTALL)
    if match:
        keys_str = match.group(1)
        # Clean up keys (remove comments, whitespace)
        keys = [k.strip() for k in keys_str.split(',') if k.strip() and not k.strip().startswith('//')]
        
        # Construct arguments string
        args = ",\n\t\t".join([f"storeKeys[{k}]" for k in keys])
        args += ',\n\t\tstoretypes.NewTransientStoreKey("transient_test")'
        
        # Replace DefaultContextWithDB call
        # We look for s.ctx = testutil.DefaultContextWithDB(...)
        # The args inside might vary due to previous replacements.
        
        # Regex to match the call
        call_pattern = r's.ctx\s*=\s*testutil.DefaultContextWithDB\([^)]+\)'
        
        new_call = f's.ctx = testutil.DefaultContextWithDB(s.T(),\n\t\t{args},\n\t)'
        
        content = re.sub(call_pattern, new_call, content)
        
    # Remove stateStore block
    # It might have been modified by previous scripts (removing tKeys loop)
    # So I'll remove lines containing 'stateStore' and 'MountStoreWithDB' and 'LoadLatestVersion' and 'dbm.NewMemDB'
    
    lines = content.split('\n')
    new_lines = []
    skip = False
    for line in lines:
        if 'dbm.NewMemDB()' in line:
            # Check if it's the state store block
            if 'stateStore' in content: # heuristic
                # skip this line and subsequent ones related to manual store creation
                continue
        if 'stateStore :=' in line or 'stateStore.MountStoreWithDB' in line or 'stateStore.LoadLatestVersion' in line:
            continue
        new_lines.append(line)
        
    content = '\n'.join(new_lines)

    with open(file_path, "w") as f:
        f.write(content)
    print(f"Updated {file_path}")
