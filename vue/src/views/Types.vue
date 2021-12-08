<template>
  <div>
    <div class="container">
      <!-- this line is used by starport scaffolding # 4 -->
		<SpType modulePath="stateset.core.agreement" moduleType="Agreement"  />
		<SpType modulePath="stateset.core.agreement" moduleType="TimedoutAgreement"  />
		<SpType modulePath="stateset.core.agreement" moduleType="SentAgreement"  />
		<SpType modulePath="stateset.core.invoice" moduleType="TimedoutInvoice"  />
		<SpType modulePath="stateset.core.invoice" moduleType="SentInvoice"  />
		<SpType modulePath="stateset.core.purchaseorder" moduleType="TimedoutPurchaseorder"  />
		<SpType modulePath="stateset.core.purchaseorder" moduleType="SentPurchaseorder"  />
		<SpType modulePath="stateset.core.invoice" moduleType="Invoice"  />
		<SpType modulePath="stateset.core.purchaseorder" moduleType="Purchaseorder"  />
		<SpType modulePath="stateset.core.loan" moduleType="Loan"  />
    </div>
  </div>
</template>

<script>
export default {
  name: 'Types',
}
</script>
