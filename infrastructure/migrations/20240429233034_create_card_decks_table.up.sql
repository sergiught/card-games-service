-- Enable the UUID extension if it's not already enabled.
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE card_decks
(
    deck_id   UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    shuffled  BOOLEAN NOT NULL DEFAULT FALSE,
    remaining INT     NOT NULL,
    cards     JSONB
);
