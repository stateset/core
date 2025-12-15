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

    # Fix duplicated args
    # Pattern: ), storeKeys[authtypes.StoreKey]... .Ctx
    # We want to replace it with ).Ctx
    
    # Simple replace since the messed up string is predictable based on my previous script
    broken = ', storeKeys[authtypes.StoreKey], storetypes.NewTransientStoreKey("transient_test")).Ctx'
    if broken in content:
        content = content.replace(broken, ').Ctx')
        
    # Also check if there is an extra )
    # The output looked like `),), ...`
    # My previous output: `...("transient_test"),\n\t), storeKeys...`
    # If I replace the broken part with `.Ctx`, I get:
    # `...("transient_test"),\n\t).Ctx`
    # This looks correct:
    # s.ctx = testutil.DefaultContextWithDB(...
    #    storetypes.NewTransientStoreKey("transient_test"),
    # ).Ctx
    
    # Let's try replacing the specific string I see in the file read output.
    # `), storeKeys[authtypes.StoreKey], storetypes.NewTransientStoreKey("transient_test")).Ctx.`
    # Wait, the output has `).Ctx.` ending with dot?
    # No, `.Ctx.\n\t\tWithBlockHeight`
    
    content = content.replace(broken, ').Ctx')

    with open(file_path, "w") as f:
        f.write(content)
    print(f"Updated {file_path}")
