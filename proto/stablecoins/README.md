# Legacy `stablecoins` Protobufs

This directory contains legacy/future multi‑asset stablecoins protobuf definitions (module name `stablecoins`).  
The current, production ssUSD implementation is the **singular** `x/stablecoin` module, which supports vault‑based minting and US‑Treasury reserve‑backed mint/redeem.

Notes:
- The chain is wired to `stablecoin` (CLI exposes an alias `stablecoins` for backward compatibility).
- Do not treat these protos as canonical unless you are explicitly re‑introducing a multi‑stablecoin module with a migration plan.

