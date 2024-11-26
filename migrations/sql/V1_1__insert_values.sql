INSERT INTO "role" (name)
VALUES
    ('user'),
    ('admin'),
    ('creator');

INSERT INTO "result" (name)
VALUES
    ('white win'),
    ('black win'),
    ('draw');

INSERT INTO "time_control" (time_minutes)
VALUES (0);

INSERT INTO "engine" (name, max_level)
VALUES ('stockfish', 20);

INSERT INTO "position" (fen)
VALUES ('start_position');
