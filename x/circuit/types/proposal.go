package types

import (
	"fmt"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	// ProposalTypePauseSystem defines the type for PauseSystemProposal
	ProposalTypePauseSystem = "PauseSystem"
	// ProposalTypeResumeSystem defines the type for ResumeSystemProposal
	ProposalTypeResumeSystem = "ResumeSystem"
	// ProposalTypeTripCircuit defines the type for TripCircuitProposal
	ProposalTypeTripCircuit = "TripCircuit"
	// ProposalTypeResetCircuit defines the type for ResetCircuitProposal
	ProposalTypeResetCircuit = "ResetCircuit"
	// ProposalTypeUpdateCircuitParams defines the type for UpdateCircuitParamsProposal
	ProposalTypeUpdateCircuitParams = "UpdateCircuitParams"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypePauseSystem)
	govtypes.RegisterProposalType(ProposalTypeResumeSystem)
	govtypes.RegisterProposalType(ProposalTypeTripCircuit)
	govtypes.RegisterProposalType(ProposalTypeResetCircuit)
	govtypes.RegisterProposalType(ProposalTypeUpdateCircuitParams)
}

// PauseSystemProposal proposes to pause the entire system
type PauseSystemProposal struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	Reason          string `json:"reason"`
	DurationSeconds int64  `json:"duration_seconds"`
}

func (p *PauseSystemProposal) GetTitle() string       { return p.Title }
func (p *PauseSystemProposal) GetDescription() string { return p.Description }
func (p *PauseSystemProposal) ProposalRoute() string  { return RouterKey }
func (p *PauseSystemProposal) ProposalType() string   { return ProposalTypePauseSystem }

func (p *PauseSystemProposal) ValidateBasic() error {
	if p.Title == "" {
		return fmt.Errorf("title cannot be empty")
	}
	if p.Description == "" {
		return fmt.Errorf("description cannot be empty")
	}
	if p.Reason == "" {
		return fmt.Errorf("reason cannot be empty")
	}
	if p.DurationSeconds < 0 {
		return fmt.Errorf("duration cannot be negative")
	}
	return nil
}

func (p *PauseSystemProposal) String() string {
	return fmt.Sprintf(`Pause System Proposal:
  Title:       %s
  Description: %s
  Reason:      %s
  Duration:    %d seconds
`, p.Title, p.Description, p.Reason, p.DurationSeconds)
}

// ResumeSystemProposal proposes to resume the system
type ResumeSystemProposal struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (p *ResumeSystemProposal) GetTitle() string       { return p.Title }
func (p *ResumeSystemProposal) GetDescription() string { return p.Description }
func (p *ResumeSystemProposal) ProposalRoute() string  { return RouterKey }
func (p *ResumeSystemProposal) ProposalType() string   { return ProposalTypeResumeSystem }

func (p *ResumeSystemProposal) ValidateBasic() error {
	if p.Title == "" {
		return fmt.Errorf("title cannot be empty")
	}
	if p.Description == "" {
		return fmt.Errorf("description cannot be empty")
	}
	return nil
}

func (p *ResumeSystemProposal) String() string {
	return fmt.Sprintf(`Resume System Proposal:
  Title:       %s
  Description: %s
`, p.Title, p.Description)
}

// TripCircuitProposal proposes to trip a module's circuit breaker
type TripCircuitProposal struct {
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	ModuleName      string   `json:"module_name"`
	Reason          string   `json:"reason"`
	DisableMessages []string `json:"disable_messages,omitempty"`
}

func (p *TripCircuitProposal) GetTitle() string       { return p.Title }
func (p *TripCircuitProposal) GetDescription() string { return p.Description }
func (p *TripCircuitProposal) ProposalRoute() string  { return RouterKey }
func (p *TripCircuitProposal) ProposalType() string   { return ProposalTypeTripCircuit }

func (p *TripCircuitProposal) ValidateBasic() error {
	if p.Title == "" {
		return fmt.Errorf("title cannot be empty")
	}
	if p.Description == "" {
		return fmt.Errorf("description cannot be empty")
	}
	if p.ModuleName == "" {
		return fmt.Errorf("module name cannot be empty")
	}
	if p.Reason == "" {
		return fmt.Errorf("reason cannot be empty")
	}
	return nil
}

func (p *TripCircuitProposal) String() string {
	return fmt.Sprintf(`Trip Circuit Proposal:
  Title:       %s
  Description: %s
  Module:      %s
  Reason:      %s
`, p.Title, p.Description, p.ModuleName, p.Reason)
}

// ResetCircuitProposal proposes to reset a module's circuit breaker
type ResetCircuitProposal struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ModuleName  string `json:"module_name"`
}

func (p *ResetCircuitProposal) GetTitle() string       { return p.Title }
func (p *ResetCircuitProposal) GetDescription() string { return p.Description }
func (p *ResetCircuitProposal) ProposalRoute() string  { return RouterKey }
func (p *ResetCircuitProposal) ProposalType() string   { return ProposalTypeResetCircuit }

func (p *ResetCircuitProposal) ValidateBasic() error {
	if p.Title == "" {
		return fmt.Errorf("title cannot be empty")
	}
	if p.Description == "" {
		return fmt.Errorf("description cannot be empty")
	}
	if p.ModuleName == "" {
		return fmt.Errorf("module name cannot be empty")
	}
	return nil
}

func (p *ResetCircuitProposal) String() string {
	return fmt.Sprintf(`Reset Circuit Proposal:
  Title:       %s
  Description: %s
  Module:      %s
`, p.Title, p.Description, p.ModuleName)
}

// UpdateCircuitParamsProposal proposes to update circuit breaker parameters
type UpdateCircuitParamsProposal struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Params      Params `json:"params"`
}

func (p *UpdateCircuitParamsProposal) GetTitle() string       { return p.Title }
func (p *UpdateCircuitParamsProposal) GetDescription() string { return p.Description }
func (p *UpdateCircuitParamsProposal) ProposalRoute() string  { return RouterKey }
func (p *UpdateCircuitParamsProposal) ProposalType() string   { return ProposalTypeUpdateCircuitParams }

func (p *UpdateCircuitParamsProposal) ValidateBasic() error {
	if p.Title == "" {
		return fmt.Errorf("title cannot be empty")
	}
	if p.Description == "" {
		return fmt.Errorf("description cannot be empty")
	}
	return p.Params.Validate()
}

func (p *UpdateCircuitParamsProposal) String() string {
	return fmt.Sprintf(`Update Circuit Params Proposal:
  Title:       %s
  Description: %s
`, p.Title, p.Description)
}

// Ensure proposal types implement Content interface
var (
	_ govtypes.Content = &PauseSystemProposal{}
	_ govtypes.Content = &ResumeSystemProposal{}
	_ govtypes.Content = &TripCircuitProposal{}
	_ govtypes.Content = &ResetCircuitProposal{}
	_ govtypes.Content = &UpdateCircuitParamsProposal{}
)
