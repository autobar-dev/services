// @generated automatically by Diesel CLI.

pub mod sql_types {
    #[derive(diesel::sql_types::SqlType)]
    #[diesel(postgres_type(name = "client_type"))]
    pub struct ClientType;
}

diesel::table! {
    modules (id) {
        id -> Int4,
        serial_number -> Varchar,
        private_key -> Text,
        created_at -> Timestamptz,
    }
}

diesel::table! {
    use diesel::sql_types::*;
    use super::sql_types::ClientType;

    sessions (id) {
        id -> Uuid,
        client_identifier -> Int4,
        user_agent -> Nullable<Text>,
        valid_until -> Timestamptz,
        last_used -> Timestamptz,
        created_at -> Timestamptz,
        internal_id -> Int4,
        client_type -> ClientType,
    }
}

diesel::table! {
    users (id) {
        id -> Int4,
        email -> Text,
        password -> Text,
        created_at -> Timestamptz,
    }
}

diesel::allow_tables_to_appear_in_same_query!(
    modules,
    sessions,
    users,
);
