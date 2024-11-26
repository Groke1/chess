CREATE VIEW "puzzles_view"
    (user_id, username, puz_id, start_pos, is_correct,
        user_solution, solution, time_taken) AS
SELECT up.user_id, u.name, puz.id, pos.fen,
       up.is_correct, up.user_solution,
       puz.solution, up.time_taken
FROM
    puzzle puz
    JOIN user_puzzle up ON puz.id = up.puzzle_id
    JOIN position pos ON pos.id = puz.position_id
    JOIN "user" u ON u.id = up.user_id;

CREATE VIEW "puzzle_stat_view"
    (puz_id, author_id, solutions, attempts) AS
SELECT DISTINCT
    puz.id,
    puz.author_id,
    COUNT(up.is_correct) FILTER ( WHERE up.is_correct ) OVER(PARTITION BY puz.id),
    COUNT(up.is_correct) OVER(PARTITION BY puz.id)
FROM
    puzzle puz
    LEFT JOIN user_puzzle up
    ON puz.id = up.puzzle_id;

CREATE FUNCTION get_result(color TEXT, res TEXT)
RETURNS VARCHAR(5)
LANGUAGE plpgsql
AS $$
BEGIN
    RETURN
        CASE
            WHEN (color = 'white' AND res = 'white win') OR
                 (color = 'black' AND res = 'black win')
                THEN 'win'
            WHEN (color = 'white' AND res = 'black win') OR
                 (color = 'black' AND res = 'white win')
                THEN 'lose'
            ELSE 'draw' END;
END;
$$;

CREATE VIEW "multi_games_view"
    (user_id, username, game_id, color, time_control,
        result, opponent_id, date, start_position, moves) AS
SELECT u.id, u.name, m.id,
       CASE WHEN u.id = m.white_player_id THEN 'white' ELSE 'black' END,
       CONCAT(tc.time_minutes, '+', tc.increment),
       get_result(CASE WHEN u.id = m.white_player_id THEN 'white' ELSE 'black' END, r.name),
       CASE WHEN u.id = m.white_player_id THEN m.black_player_id ELSE m.white_player_id END,
       m.date, p.fen, m.moves
FROM
    "user" u
    JOIN multiplayer m ON
        m.white_player_id = u.id OR m.black_player_id = u.id
    JOIN time_control tc ON tc.id = m.time_control_id
    JOIN result r ON r.id = m.result_id
    JOIN position p ON p.id = m.position_id;


CREATE VIEW "single_games_view"
    (user_id, username, game_id, color, time_control,
     result, engine_name, level, date, start_position, moves) AS
SELECT s.player_id, u.name, s.id,
       CASE WHEN s.is_player_white THEN 'white' ELSE 'black' END,
       CONCAT(tc.time_minutes, '+', tc.increment),
       get_result(CASE WHEN s.is_player_white THEN 'white' ELSE 'black' END, r.name),
       eng.name, s.engine_level,
       s.date, p.fen, s.moves
FROM
    singleplayer s
    JOIN time_control tc ON tc.id = s.time_control_id
    JOIN result r ON r.id = s.result_id
    JOIN engine eng ON eng.id = s.engine_id
    JOIN position p ON p.id = s.position_id
    JOIN "user" u ON u.id = s.player_id;


CREATE VIEW "top_view"(user_id, username, rating, top_place) AS
SELECT
    u.id, u.name, u.rating,
    ROW_NUMBER() OVER(ORDER BY u.rating DESC)
FROM "user" u;

CREATE VIEW "stat_view" (user_id, username, rating, top_place,
                        amount_games, amount_wins,
                        amount_loses, amount_draws) AS
SELECT DISTINCT
       tv.user_id,
        u.name,
       tv.rating,
       tv.top_place,
       COUNT(mgv.result) OVER(PARTITION BY mgv.user_id),
       COUNT(mgv.result) FILTER (WHERE mgv.result = 'win') OVER(PARTITION BY mgv.user_id),
       COUNT(mgv.result) FILTER (WHERE mgv.result = 'lose') OVER(PARTITION BY mgv.user_id),
       COUNT(mgv.result) FILTER (WHERE mgv.result = 'draw') OVER(PARTITION BY mgv.user_id)
FROM
    top_view tv
    LEFT JOIN multi_games_view mgv ON tv.user_id = mgv.user_id
    JOIN "user" u ON u.id = tv.user_id;

CREATE VIEW "time_controls_stat_view" (user_id, username, time_control,
                                  amount_games, amount_wins,
                                  amount_loses, amount_draws) AS
SELECT sv.user_id, u.name, mgv.time_control,
       COUNT(*),
       COUNT(*) FILTER (WHERE mgv.result = 'win'),
       COUNT(*) FILTER (WHERE mgv.result = 'lose'),
       COUNT(*) FILTER (WHERE mgv.result = 'draw')

FROM
    stat_view sv
    JOIN multi_games_view mgv ON sv.user_id = mgv.user_id
    JOIN "user" u ON u.id = sv.user_id
    GROUP BY sv.user_id, u.name, mgv.time_control ;