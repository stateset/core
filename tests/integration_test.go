package tests

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/stateset/core/x/security/keeper"
	securitytypes "github.com/stateset/core/x/security/types"
	"github.com/stateset/core/x/analytics/keeper"
	analyticstypes "github.com/stateset/core/x/analytics/types"
	"github.com/stateset/core/x/invoice/types"
)

type IntegrationTestSuite struct {
	suite.Suite
	
	securityKeeper  keeper.Keeper
	analyticsKeeper keeper.Keeper
	ctx             types.Context
}

func (suite *IntegrationTestSuite) SetupTest() {
	// Setup test environment
	// This would typically initialize the test app and context
}

func (suite *IntegrationTestSuite) TestSecurityModuleFraudDetection() {
	// Test Case 1: Velocity-based fraud detection
	suite.Run("VelocityFraudDetection", func() {
		// Simulate rapid transactions from same address
		addr := "stateset1abc123"
		
		// Create velocity rule
		rule := securitytypes.SecurityRule{
			ID:          "velocity_rule_1",
			Name:        "High Transaction Velocity",
			Type:        securitytypes.RuleTypeVelocity,
			Threshold:   10.0, // Max 10 transactions per minute
			TimeWindow:  60,   // 1 minute window
			Action:      securitytypes.ActionAlert,
			IsActive:    true,
		}
		
		err := suite.securityKeeper.SetSecurityRule(suite.ctx, rule)
		suite.Require().NoError(err)
		
		// Test high velocity triggers alert
		for i := 0; i < 15; i++ {
			tx := securitytypes.Transaction{
				Hash:   fmt.Sprintf("tx_%d", i),
				From:   addr,
				Amount: types.NewInt64Coin("ustate", 1000),
				Time:   time.Now(),
			}
			
			alert, triggered := suite.securityKeeper.CheckTransaction(suite.ctx, tx)
			if i >= 10 {
				suite.Require().True(triggered, "Alert should be triggered after 10 transactions")
				suite.Require().Equal(securitytypes.AlertTypeVelocity, alert.Type)
			}
		}
	})
	
	// Test Case 2: Pattern-based fraud detection
	suite.Run("PatternFraudDetection", func() {
		// Test round number amount pattern (often indicates automation/bot)
		roundAmounts := []int64{1000, 2000, 5000, 10000}
		
		for _, amount := range roundAmounts {
			tx := securitytypes.Transaction{
				Hash:   fmt.Sprintf("tx_round_%d", amount),
				From:   "stateset1suspicious",
				Amount: types.NewInt64Coin("ustate", amount),
				Time:   time.Now(),
			}
			
			alert, triggered := suite.securityKeeper.CheckTransaction(suite.ctx, tx)
			if amount >= 5000 {
				suite.Require().True(triggered, "Large round amounts should trigger pattern alert")
				suite.Require().Equal(securitytypes.AlertTypePattern, alert.Type)
			}
		}
	})
}

func (suite *IntegrationTestSuite) TestAnalyticsModuleMetrics() {
	suite.Run("BlockMetricsRecording", func() {
		// Test recording block metrics
		blockMetric := analyticstypes.PerformanceMetric{
			Name:        "block_time",
			Value:       2.5, // 2.5 seconds
			Timestamp:   time.Now(),
			BlockHeight: 12345,
			Tags: map[string]string{
				"validator": "validator1",
				"network":   "testnet",
			},
		}
		
		err := suite.analyticsKeeper.RecordMetric(suite.ctx, blockMetric)
		suite.Require().NoError(err)
		
		// Retrieve metrics and verify
		endTime := time.Now().Add(time.Minute)
		startTime := time.Now().Add(-time.Minute)
		
		metrics, err := suite.analyticsKeeper.GetMetrics(suite.ctx, startTime, endTime)
		suite.Require().NoError(err)
		suite.Require().Len(metrics, 1)
		suite.Require().Equal("block_time", metrics[0].Name)
		suite.Require().Equal(2.5, metrics[0].Value)
	})
	
	suite.Run("BusinessMetricsAggregation", func() {
		// Test business metrics calculation
		businessMetrics := analyticstypes.BusinessMetrics{
			TotalInvoices:        100,
			TotalAgreements:      25,
			TotalLoanValue:       1500000.50,
			ActivePurchaseOrders: 15,
			SecurityAlerts:       3,
		}
		
		// Convert to performance metric for storage
		metric := analyticstypes.PerformanceMetric{
			Name:        "business_summary",
			Value:       float64(businessMetrics.TotalInvoices),
			Timestamp:   time.Now(),
			BlockHeight: suite.ctx.BlockHeight(),
			Tags: map[string]string{
				"type":         "business",
				"agreements":   fmt.Sprintf("%d", businessMetrics.TotalAgreements),
				"loan_value":   fmt.Sprintf("%.2f", businessMetrics.TotalLoanValue),
				"purchase_orders": fmt.Sprintf("%d", businessMetrics.ActivePurchaseOrders),
			},
		}
		
		err := suite.analyticsKeeper.RecordMetric(suite.ctx, metric)
		suite.Require().NoError(err)
	})
}

func (suite *IntegrationTestSuite) TestEnhancedInvoiceWorkflow() {
	suite.Run("AutomatedPaymentScheduling", func() {
		// Test enhanced invoice with payment scheduling
		invoice := invoicetypes.EnhancedInvoice{
			ID:           "INV-2024-001",
			Amount:       types.NewInt64Coin("ustate", 50000),
			DueDate:      time.Now().Add(30 * 24 * time.Hour), // 30 days
			Status:       invoicetypes.InvoiceStatusPending,
			PaymentTerms: invoicetypes.PaymentTerms{
				DueDays:           30,
				EarlyPayDiscount:  0.02, // 2% early pay discount
				LateFeePercentage: 0.015, // 1.5% late fee
			},
			Currencies: []invoicetypes.CurrencyOption{
				{
					Currency: "USD",
					Amount:   50000,
					Rate:     1.0,
				},
				{
					Currency: "EUR", 
					Amount:   45000,
					Rate:     0.9,
				},
			},
		}
		
		// Test payment schedule generation
		schedule := invoice.GeneratePaymentSchedule()
		suite.Require().NotEmpty(schedule)
		
		// Test early payment discount calculation
		earlyPayment := invoice.CalculateEarlyPaymentAmount()
		expectedDiscount := 50000 * 0.02
		suite.Require().Equal(50000-int64(expectedDiscount), earlyPayment)
	})
}

func (suite *IntegrationTestSuite) TestComplianceAndRiskManagement() {
	suite.Run("ComplianceRuleValidation", func() {
		// Test KYC compliance rule
		complianceRule := securitytypes.ComplianceRule{
			ID:           "kyc_rule_1",
			Name:         "KYC Required for Large Transactions",
			Type:         securitytypes.ComplianceTypeKYC,
			Threshold:    100000, // Amounts over $100k require KYC
			Jurisdictions: []string{"US", "EU", "UK"},
			IsActive:     true,
		}
		
		err := suite.securityKeeper.SetComplianceRule(suite.ctx, complianceRule)
		suite.Require().NoError(err)
		
		// Test transaction requiring KYC
		largeTx := securitytypes.Transaction{
			Hash:   "kyc_test_tx",
			From:   "stateset1largecustomer",
			Amount: types.NewInt64Coin("ustate", 150000),
			Time:   time.Now(),
		}
		
		compliance, required := suite.securityKeeper.CheckCompliance(suite.ctx, largeTx)
		suite.Require().True(required, "KYC should be required for large transaction")
		suite.Require().Equal(securitytypes.ComplianceTypeKYC, compliance.Type)
	})
	
	suite.Run("RiskProfileAssessment", func() {
		// Test dynamic risk profiling
		profile := securitytypes.RiskProfile{
			Address:           "stateset1risktest",
			RiskScore:         0.3, // Medium risk
			TransactionCount:  50,
			TotalVolume:       types.NewInt64Coin("ustate", 500000),
			LastActivity:      time.Now(),
			RiskFactors: []securitytypes.RiskFactor{
				{
					Type:        securitytypes.RiskFactorVelocity,
					Score:       0.2,
					Description: "Moderate transaction velocity",
				},
				{
					Type:        securitytypes.RiskFactorGeographic,
					Score:       0.1,
					Description: "Low geographic risk",
				},
			},
		}
		
		err := suite.securityKeeper.SetRiskProfile(suite.ctx, profile)
		suite.Require().NoError(err)
		
		// Retrieve and validate
		retrievedProfile, found := suite.securityKeeper.GetRiskProfile(suite.ctx, "stateset1risktest")
		suite.Require().True(found)
		suite.Require().Equal(0.3, retrievedProfile.RiskScore)
		suite.Require().Len(retrievedProfile.RiskFactors, 2)
	})
}

func (suite *IntegrationTestSuite) TestPerformanceOptimizations() {
	suite.Run("MemoryUsageOptimization", func() {
		// Test that memory usage stays within acceptable bounds
		initialMemory := runtime.MemStats{}
		runtime.ReadMemStats(&initialMemory)
		
		// Perform intensive operations
		for i := 0; i < 1000; i++ {
			metric := analyticstypes.PerformanceMetric{
				Name:        fmt.Sprintf("test_metric_%d", i),
				Value:       float64(i),
				Timestamp:   time.Now(),
				BlockHeight: int64(i),
			}
			
			err := suite.analyticsKeeper.RecordMetric(suite.ctx, metric)
			suite.Require().NoError(err)
		}
		
		finalMemory := runtime.MemStats{}
		runtime.ReadMemStats(&finalMemory)
		
		// Memory increase should be reasonable (less than 100MB)
		memoryIncrease := finalMemory.Alloc - initialMemory.Alloc
		suite.Require().Less(memoryIncrease, uint64(100*1024*1024), "Memory usage should not increase by more than 100MB")
	})
}

func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}