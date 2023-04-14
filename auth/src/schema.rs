// @generated automatically by Diesel CLI.

diesel::table! {
    sessions (id) {
        id -> Uuid,
        user_id -> Int4,
        user_agent -> Nullable<Text>,
        valid_until -> Timestamptz,
        last_used -> Timestamptz,
        created_at -> Timestamptz,
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
    sessions,
    users,
);
