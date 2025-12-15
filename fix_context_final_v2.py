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

    # Update imports
    # Remove testutil, add tmproto
    if '"github.com/cosmos/cosmos-sdk/testutil"' in content:
        content = content.replace('"github.com/cosmos/cosmos-sdk/testutil"', '')
        
    if '"github.com/cometbft/cometbft/proto/tendermint/types"' not in content:
        # Add it after "testing"
        content = content.replace(
            '\t"testing"',
            '\t"testing"\n\ttmproto "github.com/cometbft/cometbft/proto/tendermint/types"'
        )

    # Update NewContext call
    old_call = 'sdk.NewContext(cms, false, log.NewNopLogger())'
    new_call = 'sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())'
    
    if old_call in content:
        content = content.replace(old_call, new_call)

    with open(file_path, "w") as f:
        f.write(content)
    print(f"Updated {file_path}")
