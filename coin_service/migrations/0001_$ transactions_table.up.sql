CREATE TABLE transactions (
                                                   id SERIAL PRIMARY KEY,
                                                   user_id VARCHAR NOT NULL,
                                                   category VARCHAR NOT NULL,
                                                   amount DOUBLE PRECISION NOT NULL,
                                                   currency VARCHAR NOT NULL,
                                                   timestamp TIMESTAMP NOT NULL,
                                                   description TEXT

                          )