use chrono::Utc;
use std::fs;

use crate::types;

pub fn get_meta_factors() -> types::MetaFactors {
    let hash = fs::read_to_string(".meta/HASH")
        .unwrap_or("".to_string())
        .trim_end()
        .to_string();

    let version = fs::read_to_string(".meta/VERSION")
        .unwrap_or("".to_string())
        .trim_end()
        .to_string();

    let start_time = Utc::now();

    types::MetaFactors {
        hash,
        version,
        start_time,
    }
}

pub fn get_meta_from_factors(factors: types::MetaFactors) -> types::Meta {
    let uptime = Utc::now() - factors.start_time;

    types::Meta {
        hash: factors.hash,
        version: factors.version,
        uptime: uptime.num_seconds(),
    }
}
