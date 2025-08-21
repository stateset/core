use cosmwasm_std::{Addr, Api, Deps, StdResult, Storage, Uint128};
use cw_storage_plus::Map;

/// Gas-optimized batch operations for reading multiple state items
pub struct BatchReader<'a> {
    storage: &'a dyn Storage,
}

impl<'a> BatchReader<'a> {
    pub fn new(storage: &'a dyn Storage) -> Self {
        Self { storage }
    }
    
    /// Read multiple items from a map in a single operation
    pub fn read_batch<K, V>(&self, map: &Map<K, V>, keys: &[K]) -> Vec<Option<V>>
    where
        K: cw_storage_plus::PrimaryKey<'a>,
        V: serde::de::DeserializeOwned,
    {
        keys.iter()
            .map(|key| map.may_load(self.storage, key.clone()).ok().flatten())
            .collect()
    }
}

/// Cache frequently accessed values to reduce storage reads
pub struct CachedValue<T> {
    value: Option<T>,
    loaded: bool,
}

impl<T: Clone> CachedValue<T> {
    pub fn new() -> Self {
        Self {
            value: None,
            loaded: false,
        }
    }
    
    pub fn get_or_load<F>(&mut self, loader: F) -> StdResult<T>
    where
        F: FnOnce() -> StdResult<T>,
    {
        if !self.loaded {
            self.value = Some(loader()?);
            self.loaded = true;
        }
        Ok(self.value.as_ref().unwrap().clone())
    }
    
    pub fn invalidate(&mut self) {
        self.loaded = false;
        self.value = None;
    }
}

/// Optimized validation functions that short-circuit on first error
pub fn validate_amount_nonzero(amount: &Uint128, asset: &str) -> Result<(), String> {
    if amount.is_zero() {
        return Err(format!("Amount must be greater than 0 for {}", asset));
    }
    Ok(())
}

pub fn validate_address_batch(api: &dyn Api, addresses: &[String]) -> StdResult<Vec<Addr>> {
    addresses
        .iter()
        .map(|addr| api.addr_validate(addr))
        .collect()
}

/// Gas-efficient comparison operations
#[inline(always)]
pub fn is_within_range(value: Uint128, min: Uint128, max: Uint128) -> bool {
    value >= min && value <= max
}

#[inline(always)]
pub fn min_u128(a: Uint128, b: Uint128) -> Uint128 {
    if a < b { a } else { b }
}

#[inline(always)]
pub fn max_u128(a: Uint128, b: Uint128) -> Uint128 {
    if a > b { a } else { b }
}