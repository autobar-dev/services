// @generated automatically by Diesel CLI.

diesel::table! {
    currencies (id) {
        id -> Int4,
        code -> Varchar,
        name -> Text,
        minor_unit_divisor -> Int4,
        enabled -> Bool,
        updated_at -> Timestamptz,
        created_at -> Timestamptz,
        symbol -> Nullable<Varchar>,
    }
}

diesel::table! {
    rates (id) {
        id -> Int4,
        from_currency -> Varchar,
        to_currency -> Varchar,
        rate -> Float8,
        updated_at -> Timestamptz,
    }
}

diesel::allow_tables_to_appear_in_same_query!(
    currencies,
    rates,
);
