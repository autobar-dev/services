pub mod delete;
pub mod index;
pub mod remote;
pub mod set;
pub mod update;

pub use delete::delete_route;
pub use index::index_route;
pub use remote::remote_route;
pub use set::set_route;
pub use update::update_route;
