package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	// ProposalTypeCreateAttribute defines the type for a CreateRootAttribute
	ProposalTypeCreateAttribute = "CreateAttribute"
)

// Assert CreateAttributeProposal implements govtypes.Content at compile-time
var _ govtypes.Content = &CreateAttributeProposal{}

func init() {
	govtypes.RegisterProposalType(ProposalTypeCreateAttribute)
	govtypes.RegisterProposalTypeCodec(&CreateAttributeProposal{}, "provenance/attribute/CreateAttributeProposal")
}

// NewCreateAttributeProposal create a new governance proposal request to create a root name
//nolint:interfacer
func NewCreateAttributeProposal(title, description, name string, owner sdk.AccAddress, value string) *CreateAttributeProposal {
	return &CreateAttributeProposal{
		Title:       title,
		Description: description,
		Name:        name,
		Value:       value,
		Owner:       owner.String(),
	}
}

// GetTitle returns the title of a community pool spend proposal.
func (crnp *CreateAttributeProposal) GetTitle() string { return crnp.Title }

// GetDescription returns the description of a community pool spend proposal.
func (crnp *CreateAttributeProposal) GetDescription() string { return crnp.Description }

// ProposalRoute returns the routing key of a community pool spend proposal.
func (crnp *CreateAttributeProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a community pool spend proposal.
func (crnp *CreateAttributeProposal) ProposalType() string { return ProposalTypeCreateAttribute }

// ValidateBasic runs basic stateless validity checks
func (crnp *CreateAttributeProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(crnp)
	if err != nil {
		return err
	}
	if strings.TrimSpace(crnp.Owner) != "" {
		if _, err := sdk.AccAddressFromBech32(crnp.Owner); err != nil {
			return ErrInvalidAddress
		}
	}
	if strings.TrimSpace(crnp.Name) == "" {
		return ErrInvalidLengthAttribute
	}
	if strings.Contains(crnp.Name, ".") {
		return ErrAttributeContainsSegments
	}

	return nil
}

// String implements the Stringer interface.
func (crnp CreateAttributeProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`Create Attribute Proposal:
  Title:       %s
  Description: %s
  Owner:       %s
  Name:        %s
  Value:       %s
`, crnp.Title, crnp.Description, crnp.Owner, crnp.Name, crnp.Value))
	return b.String()
}
