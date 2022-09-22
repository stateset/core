// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import StatesetCoreStatesetCoreAgreement from './stateset/core/stateset.core.agreement';
import StatesetCoreStatesetCoreInvoice from './stateset/core/stateset.core.invoice';
import StatesetCoreStatesetCoreLoan from './stateset/core/stateset.core.loan';
import StatesetCoreStatesetCorePurchaseorder from './stateset/core/stateset.core.purchaseorder';
import StatesetCoreStatesetCoreRefund from './stateset/core/stateset.core.refund';
export default {
    StatesetCoreStatesetCoreAgreement: load(StatesetCoreStatesetCoreAgreement, 'stateset.core.agreement'),
    StatesetCoreStatesetCoreInvoice: load(StatesetCoreStatesetCoreInvoice, 'stateset.core.invoice'),
    StatesetCoreStatesetCoreLoan: load(StatesetCoreStatesetCoreLoan, 'stateset.core.loan'),
    StatesetCoreStatesetCorePurchaseorder: load(StatesetCoreStatesetCorePurchaseorder, 'stateset.core.purchaseorder'),
    StatesetCoreStatesetCoreRefund: load(StatesetCoreStatesetCoreRefund, 'stateset.core.refund'),
};
function load(mod, fullns) {
    return function init(store) {
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: ' + fullns);
        }
        else {
            store.registerModule([fullns], mod);
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns + '/init', null, {
                        root: true
                    });
                }
            });
        }
    };
}
