import os

files = [
    "tests/integration/ecommerce_flow_test.go",
    "tests/integration/ai_agent_transaction_test.go",
    "tests/integration/circuit_breaker_test.go",
    "tests/integration/cross_module_compliance_test.go",
    "tests/integration/stablecoin_lifecycle_test.go"
]

setup_decl = "func (s *%s) SetupTest() {"
config_setup = """
	// Set SDK config
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("stateset", "statesetpub")
	config.SetBech32PrefixForValidator("statesetvaloper", "statesetvaloperpub")
	config.SetBech32PrefixForConsensusNode("statesetvalcons", "statesetvalconspub")
"""

struct_names = {
    "tests/integration/ecommerce_flow_test.go": "ECommerceFlowTestSuite",
    "tests/integration/ai_agent_transaction_test.go": "AIAgentTransactionTestSuite",
    "tests/integration/circuit_breaker_test.go": "CircuitBreakerTestSuite",
    "tests/integration/cross_module_compliance_test.go": "CrossModuleComplianceTestSuite",
    "tests/integration/stablecoin_lifecycle_test.go": "StablecoinLifecycleTestSuite"
}

for file_path in files:
    if not os.path.exists(file_path):
        continue
        
    with open(file_path, "r") as f:
        content = f.read()

    if 'config.SetBech32PrefixForAccount("stateset", "statesetpub")' in content:
        continue

    struct_name = struct_names.get(file_path)
    if struct_name:
        search_str = f"func (s *{struct_name}) SetupTest() {{"
        if search_str in content:
            content = content.replace(search_str, search_str + config_setup)
            with open(file_path, "w") as f:
                f.write(content)
            print(f"Updated {file_path}")

print("Config setup added.")
