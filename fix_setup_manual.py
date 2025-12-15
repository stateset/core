import os
import re

files = [
    "tests/integration/ecommerce_flow_test.go",
    "tests/integration/ai_agent_transaction_test.go",
    "tests/integration/circuit_breaker_test.go",
    "tests/integration/cross_module_compliance_test.go",
    "tests/integration/stablecoin_lifecycle_test.go"
]

for file_path in files:
    if not os.path.exists(file_path):
        continue
        
    with open(file_path, "r") as f:
        content = f.read()

    # We need to find where store creation starts
    # It starts with `// Create multistore`
    # And ends with `WithBlockTime(time.Now())`
    
    start_marker = "// Create multistore"
    end_marker = "WithBlockTime(time.Now())"
    
    if start_marker in content and end_marker in content:
        # Extract keys first just in case, though they are defined before this block
        # storeKeys definition is before `// Create multistore`.
        
        # New block to insert
        new_block = """// Create multistore
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	for _, key := range storeKeys {
		cms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, db)
	}
	transientKey := storetypes.NewTransientStoreKey("transient_test")
	cms.MountStoreWithDB(transientKey, storetypes.StoreTypeTransient, db)
	
	err := cms.LoadLatestVersion()
	require.NoError(s.T(), err)

	// Create context
	s.ctx = sdk.NewContext(cms, false, log.NewNopLogger()).
		WithBlockHeight(1).
		WithBlockTime(time.Now())"""

        # Regex to replace from start_marker to end_marker
        pattern = re.escape(start_marker) + r'.*?' + re.escape(end_marker)
        
        # Check if replacement works (handling newlines)
        # re.DOTALL is needed
        
        content = re.sub(pattern, new_block, content, flags=re.DOTALL)
        
        with open(file_path, "w") as f:
            f.write(content)
        print(f"Updated {file_path}")
    else:
        print(f"Markers not found in {file_path}")
