use crate::types::consts::BCRYPT_COST;

pub fn hash_string(value: String) -> String {
    bcrypt::hash(value, BCRYPT_COST).unwrap() // Probably should handle the result here
}
