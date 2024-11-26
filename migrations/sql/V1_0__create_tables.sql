CREATE TABLE "role" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(25) NOT NULL,
    UNIQUE (name)
);

CREATE TABLE "result" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    UNIQUE (name)
);

CREATE TABLE "time_control" (
    id SERIAL PRIMARY KEY,
    time_minutes INT NOT NULL,
    increment INT NOT NULL DEFAULT 0,
    UNIQUE (time_minutes, increment),
    CHECK ( time_minutes >= 0 AND time_minutes <= 120 ),
    CHECK ( increment >= 0 AND increment <= 60 )
);

CREATE TABLE "position" (
    id SERIAL PRIMARY KEY,
    fen TEXT NOT NULL,
    UNIQUE(fen)
);

CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(25) NOT NULL,
    rating INT NOT NULL DEFAULT 1000,
    role_id INT NOT NULL,
    UNIQUE(name),
    FOREIGN KEY(role_id) REFERENCES "role"(id)
        ON DELETE CASCADE
);

CREATE TABLE "engine" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(25) NOT NULL,
    min_level INT NOT NULL DEFAULT 0,
    max_level INT NOT NULL,
    CHECK ( max_level >= min_level ),
    UNIQUE (name)
);

CREATE TABLE "multiplayer" (
    id SERIAL PRIMARY KEY,
    white_player_id INT NOT NULL,
    black_player_id INT NOT NULL,
    date DATE NOT NULL,
    result_id INT NOT NULL,
    position_id INT NOT NULL,
    time_control_id INT NOT NULL,
    moves TEXT NOT NULL,
    FOREIGN KEY (white_player_id) REFERENCES "user"(id)
        ON DELETE CASCADE,
    FOREIGN KEY (black_player_id) REFERENCES "user"(id)
        ON DELETE CASCADE,
    FOREIGN KEY (result_id) REFERENCES "result"(id)
        ON DELETE CASCADE,
    FOREIGN KEY (position_id) REFERENCES "position"(id)
        ON DELETE CASCADE,
    FOREIGN KEY (time_control_id) REFERENCES "time_control"(id)
        ON DELETE CASCADE,
    CHECK (white_player_id != black_player_id)
);

CREATE TABLE "singleplayer" (
    id SERIAL PRIMARY KEY,
    player_id INT NOT NULL,
    engine_id INT NOT NULL,
    engine_level INT NOT NULL,
    is_player_white BOOLEAN NOT NULL,
    date DATE NOT NULL,
    result_id INT NOT NULL,
    position_id INT NOT NULL,
    time_control_id INT NOT NULL,
    moves TEXT NOT NULL,
    FOREIGN KEY (player_id) REFERENCES "user"(id)
        ON DELETE CASCADE,
    FOREIGN KEY (engine_id) REFERENCES "engine"(id)
        ON DELETE CASCADE,
    FOREIGN KEY (result_id) REFERENCES "result"(id)
        ON DELETE CASCADE,
    FOREIGN KEY (position_id) REFERENCES "position"(id)
        ON DELETE CASCADE,
    FOREIGN KEY (time_control_id) REFERENCES "time_control"(id)
        ON DELETE CASCADE
);

CREATE TABLE "puzzle" (
    id SERIAL PRIMARY KEY,
    position_id INT NOT NULL,
    solution TEXT NOT NULL,
    author_id INT,
    UNIQUE (position_id),
    FOREIGN KEY (position_id) REFERENCES "position"(id)
        ON DELETE CASCADE,
    FOREIGN KEY (author_id) REFERENCES "user"(id)
        ON DELETE SET NULL
);

CREATE TABLE "user_puzzle" (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    puzzle_id INT NOT NULL,
    user_solution TEXT NOT NULL,
    time_taken TIME NOT NULL,
    is_correct BOOLEAN NOT NULL,
    FOREIGN KEY (user_id) REFERENCES "user"(id)
        ON DELETE CASCADE,
    FOREIGN KEY (puzzle_id) REFERENCES "puzzle"(id)
        ON DELETE CASCADE,
    UNIQUE (user_id, puzzle_id)
);