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

    # Replace user account creation with NewAccountWithAddress
    # Pattern: var := authtypes.NewBaseAccountWithAddress(addr)
    content = re.sub(
        r'(\w+)\s*:=\s*authtypes.NewBaseAccountWithAddress\(([^)]+)\)',
        r'\1 := s.accountKeeper.NewAccountWithAddress(s.ctx, \2)',
        content
    )

    # Fix Module Account creation
    # We need to capture the creation and the SetModuleAccount call
    # This is trickier with regex. I'll search for specific patterns I saw.
    
    # Generic replacement for SetModuleAccount call to ensure it uses an account with ID
    # But I need to call NewAccount first.
    
    # Let's replace the whole setupTestAccounts function body if possible or specific lines.
    
    # Strategy: Find lines like `acc := authtypes.NewEmptyModuleAccount(...)`
    # and append `acc = s.accountKeeper.NewAccount(s.ctx, acc).(*authtypes.ModuleAccount)`
    
    lines = content.split('\n')
    new_lines = []
    for line in lines:
        new_lines.append(line)
        if 'authtypes.NewEmptyModuleAccount(' in line:
            # Extract variable name
            match = re.search(r'(\w+)\s*:=', line)
            if match:
                var_name = match.group(1)
                # Add line to set account number
                # We cast to *authtypes.ModuleAccount because NewEmptyModuleAccount returns *ModuleAccount
                # and NewAccount returns AccountI.
                new_lines.append(f'\t{var_name} = s.accountKeeper.NewAccount(s.ctx, {var_name}).(*authtypes.ModuleAccount)')

    content = '\n'.join(new_lines)

    with open(file_path, "w") as f:
        f.write(content)
    print(f"Updated {file_path}")
