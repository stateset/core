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

    # Extract keys
    match = re.search(r'storeKeys\s*:=\s*storetypes.NewKVStoreKeys\((.*?)\)', content, re.DOTALL)
    if match:
        keys_str = match.group(1)
        keys = [k.strip() for k in keys_str.split(',') if k.strip() and not k.strip().startswith('//')]
        
        args = ",\n\t\t".join([f"storeKeys[{k}]" for k in keys])
        args += ',\n\t\tstoretypes.NewTransientStoreKey("transient_test")'
        
        new_block = f's.ctx = testutil.DefaultContextWithDB(\n\t\t{args},\n\t).Ctx'
        
        # Replace the broken block
        # The block starts with s.ctx = testutil.DefaultContextWithDB
        # and ends with .Ctx
        # But my previous attempts might have left it weird.
        # I'll look for the start, and replace until `.Ctx`
        
        pattern = r's.ctx\s*=\s*testutil.DefaultContextWithDB.*?\.Ctx'
        
        # Check if we can match the broken state
        # The broken state has `)).Ctx` or `).Ctx` or `), ... .Ctx`
        # Using DOTALL to match newlines
        
        # To be safe, I'll match until `.Ctx.` (with dot) because it's followed by `.WithBlockHeight`
        # But wait, `.Ctx` is a field, not method. So `.Ctx.` is correct (field access then method call on context?).
        # No, `Ctx` is a field of type `sdk.Context`. `WithBlockHeight` is a method of `sdk.Context`.
        
        content = re.sub(pattern, new_block, content, flags=re.DOTALL)
        
        # Fix any potential double dots ..
        content = content.replace('.Ctx..', '.Ctx.')

    with open(file_path, "w") as f:
        f.write(content)
    print(f"Updated {file_path}")
