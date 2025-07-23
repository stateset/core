package types

import "fmt"

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		SecurityRules:      []SecurityRule{},
		SecurityAlerts:     []SecurityAlert{},
		RiskProfiles:       []RiskProfile{},
		ComplianceRules:    []ComplianceRule{},
		TransactionMonitors: []TransactionMonitor{},
		Params:             DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicate security rule IDs
	securityRuleIndexMap := make(map[string]struct{})
	for _, elem := range gs.SecurityRules {
		index := elem.ID
		if _, ok := securityRuleIndexMap[index]; ok {
			return fmt.Errorf("duplicated id for securityRule")
		}
		securityRuleIndexMap[index] = struct{}{}
		
		if err := elem.Validate(); err != nil {
			return fmt.Errorf("invalid security rule %s: %w", elem.ID, err)
		}
	}

	// Check for duplicate security alert IDs
	securityAlertIndexMap := make(map[string]struct{})
	for _, elem := range gs.SecurityAlerts {
		index := elem.ID
		if _, ok := securityAlertIndexMap[index]; ok {
			return fmt.Errorf("duplicated id for securityAlert")
		}
		securityAlertIndexMap[index] = struct{}{}
		
		if err := elem.Validate(); err != nil {
			return fmt.Errorf("invalid security alert %s: %w", elem.ID, err)
		}
	}

	// Check for duplicate risk profile addresses
	riskProfileIndexMap := make(map[string]struct{})
	for _, elem := range gs.RiskProfiles {
		index := elem.Address
		if _, ok := riskProfileIndexMap[index]; ok {
			return fmt.Errorf("duplicated address for riskProfile")
		}
		riskProfileIndexMap[index] = struct{}{}
		
		if err := elem.Validate(); err != nil {
			return fmt.Errorf("invalid risk profile %s: %w", elem.Address, err)
		}
	}

	// Validate params
	if err := gs.Params.Validate(); err != nil {
		return fmt.Errorf("invalid params: %w", err)
	}

	return nil
}

// GenesisState defines the security module's genesis state.
type GenesisState struct {
	SecurityRules       []SecurityRule       `json:"security_rules"`
	SecurityAlerts      []SecurityAlert      `json:"security_alerts"`
	RiskProfiles        []RiskProfile        `json:"risk_profiles"`
	ComplianceRules     []ComplianceRule     `json:"compliance_rules"`
	TransactionMonitors []TransactionMonitor `json:"transaction_monitors"`
	Params              Params               `json:"params"`
}

// Params defines the parameters for the security module.
type Params struct {
	EnableFraudDetection     bool   `json:"enable_fraud_detection"`
	EnableVelocityMonitoring bool   `json:"enable_velocity_monitoring"`
	EnableComplianceCheck    bool   `json:"enable_compliance_check"`
	MaxRiskScore             int32  `json:"max_risk_score"`
	AlertThreshold           int32  `json:"alert_threshold"`
	BlockThreshold           int32  `json:"block_threshold"`
	MonitoringWindow         string `json:"monitoring_window"` // Duration format
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return Params{
		EnableFraudDetection:     true,
		EnableVelocityMonitoring: true,
		EnableComplianceCheck:    true,
		MaxRiskScore:             100,
		AlertThreshold:           70,
		BlockThreshold:           90,
		MonitoringWindow:         "24h",
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if p.MaxRiskScore < 0 || p.MaxRiskScore > 100 {
		return fmt.Errorf("max risk score must be between 0 and 100")
	}
	if p.AlertThreshold < 0 || p.AlertThreshold > p.MaxRiskScore {
		return fmt.Errorf("alert threshold must be between 0 and max risk score")
	}
	if p.BlockThreshold < 0 || p.BlockThreshold > p.MaxRiskScore {
		return fmt.Errorf("block threshold must be between 0 and max risk score")
	}
	if p.AlertThreshold > p.BlockThreshold {
		return fmt.Errorf("alert threshold cannot be higher than block threshold")
	}
	return nil
}