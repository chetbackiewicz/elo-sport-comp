BEGIN TRANSACTION;

-- Insert completed bouts
INSERT INTO bout (challenger_id, acceptor_id, referee_id, style_id, accepted, completed, cancelled, points)
VALUES 
(1, 3, 9, 1, true, true, false, 20), -- John Smith vs Mike Johnson in Muay Thai (completed)
(8, 5, 10, 2, true, true, false, 15), -- Jessica Taylor vs David Brown in BJJ (completed)
(4, 7, 2, 3, true, true, false, 25), -- Sarah Williams vs Robert Wilson in Boxing (completed)
(6, 9, 1, 1, true, true, false, 18), -- Emily Davis vs Michael Anderson in Muay Thai (completed)
(10, 4, 3, 3, true, true, false, 22); -- Lisa Thomas vs Sarah Williams in Boxing (completed)

-- Insert outcomes for completed bouts
INSERT INTO outcome (bout_id, winner_id, loser_id, style_id, is_draw)
VALUES 
(4, 1, 3, 1, false), -- John Smith beat Mike Johnson in Muay Thai
(5, 5, 8, 2, false), -- David Brown beat Jessica Taylor in BJJ
(6, 7, 4, 3, false), -- Robert Wilson beat Sarah Williams in Boxing
(7, 9, 6, 1, false), -- Michael Anderson beat Emily Davis in Muay Thai
(8, 4, 10, 3, false); -- Sarah Williams beat Lisa Thomas in Boxing

-- Update athlete records for winners and losers
-- John Smith wins +1
UPDATE athlete_record SET wins = wins + 1 WHERE athlete_id = 1;
-- Mike Johnson losses +1
UPDATE athlete_record SET losses = losses + 1 WHERE athlete_id = 3;
-- David Brown wins +1
UPDATE athlete_record SET wins = wins + 1 WHERE athlete_id = 5;
-- Jessica Taylor losses +1
UPDATE athlete_record SET losses = losses + 1 WHERE athlete_id = 8;
-- Robert Wilson wins +1
UPDATE athlete_record SET wins = wins + 1 WHERE athlete_id = 7;
-- Sarah Williams losses +1 (first match)
UPDATE athlete_record SET losses = losses + 1 WHERE athlete_id = 4;
-- Michael Anderson wins +1
UPDATE athlete_record SET wins = wins + 1 WHERE athlete_id = 9;
-- Emily Davis losses +1
UPDATE athlete_record SET losses = losses + 1 WHERE athlete_id = 6;
-- Sarah Williams wins +1 (second match)
UPDATE athlete_record SET wins = wins + 1 WHERE athlete_id = 4;
-- Lisa Thomas losses +1
UPDATE athlete_record SET losses = losses + 1 WHERE athlete_id = 10;

-- Update athlete scores for winners (gained points)
-- John Smith + 20 in Muay Thai
UPDATE athlete_score SET score = score + 20 WHERE athlete_id = 1 AND style_id = 1;
-- David Brown + 15 in BJJ
UPDATE athlete_score SET score = score + 15 WHERE athlete_id = 5 AND style_id = 2;
-- Robert Wilson + 25 in Boxing
UPDATE athlete_score SET score = score + 25 WHERE athlete_id = 7 AND style_id = 3;
-- Michael Anderson + 18 in Muay Thai
UPDATE athlete_score SET score = score + 18 WHERE athlete_id = 9 AND style_id = 1;
-- Sarah Williams + 22 in Boxing
UPDATE athlete_score SET score = score + 22 WHERE athlete_id = 4 AND style_id = 3;

-- Update athlete scores for losers (lost points)
-- Mike Johnson - a portion of points in Muay Thai
UPDATE athlete_score SET score = score - 10 WHERE athlete_id = 3 AND style_id = 1;
-- Jessica Taylor - a portion of points in BJJ
UPDATE athlete_score SET score = score - 8 WHERE athlete_id = 8 AND style_id = 2;
-- Sarah Williams - a portion of points in Boxing (first match)
UPDATE athlete_score SET score = score - 12 WHERE athlete_id = 4 AND style_id = 3;
-- Emily Davis - a portion of points in Muay Thai
UPDATE athlete_score SET score = score - 9 WHERE athlete_id = 6 AND style_id = 1;
-- Lisa Thomas - a portion of points in Boxing
UPDATE athlete_score SET score = score - 11 WHERE athlete_id = 10 AND style_id = 3;

COMMIT;
