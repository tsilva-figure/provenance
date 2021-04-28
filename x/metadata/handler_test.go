package metadata_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/provenance-io/provenance/app"
	"github.com/provenance-io/provenance/x/metadata"
	"github.com/provenance-io/provenance/x/metadata/types"
	"github.com/provenance-io/provenance/x/metadata/types/p8e"
)

type HandlerTestSuite struct {
	suite.Suite

	app     *app.App
	ctx     sdk.Context
	handler sdk.Handler

	pubkey1   cryptotypes.PubKey
	user1     string
	user1Addr sdk.AccAddress

	pubkey2   cryptotypes.PubKey
	user2     string
	user2Addr sdk.AccAddress
}

func (s *HandlerTestSuite) SetupTest() {
	s.app = app.Setup(false)
	s.ctx = s.app.BaseApp.NewContext(false, tmproto.Header{})
	s.handler = metadata.NewHandler(s.app.MetadataKeeper)

	s.pubkey1 = secp256k1.GenPrivKey().PubKey()
	s.user1Addr = sdk.AccAddress(s.pubkey1.Address())
	s.user1 = s.user1Addr.String()

	s.pubkey2 = secp256k1.GenPrivKey().PubKey()
	s.user2Addr = sdk.AccAddress(s.pubkey2.Address())
	s.user2 = s.user2Addr.String()

	s.app.AccountKeeper.SetAccount(s.ctx, s.app.AccountKeeper.NewAccountWithAddress(s.ctx, s.user1Addr))
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func createContractSpec(inputSpecs []*p8e.DefinitionSpec, outputSpec p8e.OutputSpec, definitionSpec p8e.DefinitionSpec) p8e.ContractSpec {
	return p8e.ContractSpec{ConsiderationSpecs: []*p8e.ConsiderationSpec{
		{FuncName: "additionalParties",
			InputSpecs:       inputSpecs,
			OutputSpec:       &outputSpec,
			ResponsibleParty: 1,
		},
	},
		Definition:      &definitionSpec,
		InputSpecs:      inputSpecs,
		PartiesInvolved: []p8e.PartyType{p8e.PartyType_PARTY_TYPE_AFFILIATE},
	}
}

func createDefinitionSpec(name string, classname string, reference p8e.ProvenanceReference, defType int) p8e.DefinitionSpec {
	return p8e.DefinitionSpec{
		Name: name,
		ResourceLocation: &p8e.Location{Classname: classname,
			Ref: &reference,
		},
		Type: 1,
	}
}

// TODO: AddScope tests
// TODO: DeleteScope tests
// TODO: AddSession tests
// TODO: AddRecord tests
// TODO: DeleteRecord tests
// TODO: AddScopeSpecification tests
// TODO: DeleteScopeSpecification tests
// TODO: AddContractSpecification tests
// TODO: DeleteContractSpecification tests
// TODO: AddRecordSpecification tests
// TODO: DeleteRecordSpecification tests

func (s HandlerTestSuite) TestAddP8EContractSpec() {
	validDefSpec := createDefinitionSpec("perform_input_checks", "io.provenance.loan.LoanProtos$PartiesList", p8e.ProvenanceReference{Hash: "Adv+huolGTKofYCR0dw5GHm/R7sUWOwF32XR8r8r9kDy4il5U/LApxOWYHb05jhK4+eY4YzRMRiWcxU3Lx0+Mw=="}, 1)
	invalidDefSpec := createDefinitionSpec("perform_action", "", p8e.ProvenanceReference{Hash: "Adv+huolGTKofYCR0dw5GHm/R7sUWOwF32XR8r8r9kDy4il5U/LApxOWYHb05jhK4+eY4YzRMRiWcxU3Lx0+Mw=="}, 1)

	cases := []struct {
		name     string
		v39CSpec p8e.ContractSpec
		signers  []string
		errorMsg string
	}{
		{
			"should successfully ADD contract spec in from v38 to v40",
			createContractSpec([]*p8e.DefinitionSpec{&validDefSpec}, p8e.OutputSpec{Spec: &validDefSpec}, validDefSpec),
			[]string{s.user1},
			"",
		},
		{
			"should successfully UPDATE contract spec in from v38 to v40",
			createContractSpec([]*p8e.DefinitionSpec{&validDefSpec}, p8e.OutputSpec{Spec: &validDefSpec}, validDefSpec),
			[]string{s.user1},
			"",
		},
		{
			"should fail to add due to invalid signers",
			createContractSpec([]*p8e.DefinitionSpec{&validDefSpec}, p8e.OutputSpec{Spec: &validDefSpec}, validDefSpec),
			[]string{s.user2},
			fmt.Sprintf("missing signature from existing owner %s; required for update", s.user1),
		},
		{
			"should fail on converting contract validate basic",
			createContractSpec([]*p8e.DefinitionSpec{&invalidDefSpec}, p8e.OutputSpec{Spec: &validDefSpec}, validDefSpec),
			[]string{s.user1},
			"input specification type name cannot be empty",
		},
	}

	for _, tc := range cases {
		s.T().Run(tc.name, func(t *testing.T) {
			_, err := s.handler(s.ctx, &types.MsgWriteP8EContractSpecRequest{Contractspec: tc.v39CSpec, Signers: tc.signers})
			if len(tc.errorMsg) > 0 {
				assert.EqualError(t, err, tc.errorMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TODO: P8EMemorializeContract tests
// TODO: BindOSLocatorRequest tests
// TODO: DeleteOSLocatorRequest tests
// TODO: ModifyOSLocatorRequest tests

func ownerPartyList(addresses ...string) []types.Party {
	retval := make([]types.Party, len(addresses))
	for i, addr := range addresses {
		retval[i] = types.Party{Address: addr, Role: types.PartyType_PARTY_TYPE_OWNER}
	}
	return retval
}

func (s HandlerTestSuite) TestAddAndDeleteScopeDataAccess() {

	scopeID := types.ScopeMetadataAddress(uuid.New())
	scopeSpecID := types.ScopeSpecMetadataAddress(uuid.New())
	scope := types.NewScope(scopeID, scopeSpecID, ownerPartyList(s.user1), []string{s.user1}, "")
	dneScopeID := types.ScopeMetadataAddress(uuid.New())

	user3 := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address()).String()

	cases := []struct {
		name     string
		msg      sdk.Msg
		signers  []string
		errorMsg string
	}{
		{
			"setup test with new scope",
			types.NewMsgWriteScopeRequest(*scope, []string{s.user1}),
			[]string{s.user1},
			"",
		},
		{
			"should fail to ADD address to data access, msg validate basic failure",
			types.NewMsgAddScopeDataAccessRequest(scopeID, []string{}, []string{s.user1}),
			[]string{s.user1},
			"data access list cannot be empty",
		},
		{
			"should fail to ADD address to data access, validate add failure",
			types.NewMsgAddScopeDataAccessRequest(dneScopeID, []string{s.user1}, []string{s.user1}),
			[]string{s.user1},
			fmt.Sprintf("scope not found with id %s", dneScopeID),
		},
		{
			"should fail to ADD address to data access, validate add failure",
			types.NewMsgAddScopeDataAccessRequest(scopeID, []string{s.user1}, []string{s.user1}),
			[]string{s.user1},
			fmt.Sprintf("address already exists for data access %s", s.user1),
		},
		{
			"should successfully ADD address to data access",
			types.NewMsgAddScopeDataAccessRequest(scopeID, []string{s.user2}, []string{s.user1}),
			[]string{s.user1},
			"",
		},
		{
			"should fail to DELETE address from data access, msg validate basic failure",
			types.NewMsgDeleteScopeDataAccessRequest(scopeID, []string{}, []string{s.user1}),
			[]string{s.user1},
			"data access list cannot be empty",
		},
		{
			"should fail to DELETE address from data access, validate add failure",
			types.NewMsgDeleteScopeDataAccessRequest(dneScopeID, []string{s.user1}, []string{s.user1}),
			[]string{s.user1},
			fmt.Sprintf("scope not found with id %s", dneScopeID),
		},
		{
			"should fail to DELETE address from data access, validate add failure",
			types.NewMsgDeleteScopeDataAccessRequest(scopeID, []string{user3}, []string{s.user1}),
			[]string{s.user1},
			fmt.Sprintf("address does not exist in scope data access: %s", user3),
		},
		{
			"should successfully DELETE address from data access",
			types.NewMsgDeleteScopeDataAccessRequest(scopeID, []string{s.user2}, []string{s.user1}),
			[]string{s.user1},
			"",
		},
	}

	for _, tc := range cases {
		s.T().Run(tc.name, func(t *testing.T) {
			_, err := s.handler(s.ctx, tc.msg)
			if len(tc.errorMsg) > 0 {
				assert.EqualError(t, err, tc.errorMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
