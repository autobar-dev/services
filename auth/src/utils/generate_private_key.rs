use rand::{distributions::Alphanumeric, Rng};

pub fn generate_private_key(length: u32) -> String {
    let private_key: String = rand::thread_rng()
        .sample_iter(&Alphanumeric)
        .take(length as usize)
        .map(char::from)
        .collect();

    private_key
}
