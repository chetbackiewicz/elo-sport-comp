BEGIN TRANSACTION;

-- Insert gyms
INSERT INTO gym (gym_name, gym_address, gym_city, gym_state, gym_zip, gym_phone, gym_email, gym_website, gym_description)
VALUES 
('Elite MMA', '123 Main St', 'Austin', 'TX', '78701', '5125551234', 'info@elitemma.com', 'www.elitemma.com', 'Premier mixed martial arts training facility'),
('Fight Club', '456 Oak Ave', 'Dallas', 'TX', '75201', '2145555678', 'contact@fightclub.com', 'www.fightclub.com', 'World-class combat sports training'),
('Champions Gym', '789 Pine Rd', 'Houston', 'TX', '77002', '7135559012', 'train@champions.com', 'www.champions.com', 'Training champions since 2005');

-- Insert styles
INSERT INTO style (style_name)
VALUES 
('Muay Thai'),
('Brazilian Jiu-Jitsu'),
('Boxing');

-- Insert athletes
INSERT INTO athlete (first_name, last_name, username, birth_date, password, email)
VALUES 
('John', 'Smith', 'jsmith', '1990-01-15', 'password123', 'john.smith@email.com'),
('John', 'Doe', 'jdoe', '1992-05-20', 'password123', 'john.doe@email.com'),
('Mike', 'Johnson', 'mjohnson', '1988-11-30', 'password123', 'mike.johnson@email.com'),
('Sarah', 'Williams', 'swilliams', '1995-03-25', 'password123', 'sarah.williams@email.com'),
('David', 'Brown', 'dbrown', '1993-07-12', 'password123', 'david.brown@email.com'),
('Emily', 'Davis', 'edavis', '1991-09-08', 'password123', 'emily.davis@email.com'),
('Robert', 'Wilson', 'rwilson', '1989-04-17', 'password123', 'robert.wilson@email.com'),
('Jessica', 'Taylor', 'jtaylor', '1994-12-03', 'password123', 'jessica.taylor@email.com'),
('Michael', 'Anderson', 'manderson', '1992-08-22', 'password123', 'michael.anderson@email.com'),
('Lisa', 'Thomas', 'lthomas', '1993-06-14', 'password123', 'lisa.thomas@email.com');

-- Insert athlete records
INSERT INTO athlete_record (athlete_id, wins, losses, draws)
SELECT athlete_id, 0, 0, 0
FROM athlete;

-- Insert athlete_gym relationships
INSERT INTO athlete_gym (athlete_id, gym_id)
VALUES 
(1, 1), -- John Smith at Elite MMA
(2, 2), -- John Doe at Fight Club
(3, 1), -- Mike Johnson at Elite MMA
(4, 3), -- Sarah Williams at Champions Gym
(5, 2), -- David Brown at Fight Club
(6, 1), -- Emily Davis at Elite MMA
(7, 3), -- Robert Wilson at Champions Gym
(8, 2), -- Jessica Taylor at Fight Club
(9, 1), -- Michael Anderson at Elite MMA
(10, 3); -- Lisa Thomas at Champions Gym

-- Insert athlete_style relationships
INSERT INTO athlete_style (athlete_id, style_id)
VALUES 
(1, 2), -- John Smith does BJJ
(2, 2), -- John Doe does BJJ
(1, 1), -- John Smith also does Muay Thai
(3, 1), -- Mike Johnson does Muay Thai
(4, 3), -- Sarah Williams does Boxing
(5, 2), -- David Brown does BJJ
(6, 1), -- Emily Davis does Muay Thai
(7, 3), -- Robert Wilson does Boxing
(8, 2), -- Jessica Taylor does BJJ
(9, 1), -- Michael Anderson does Muay Thai
(10, 3); -- Lisa Thomas does Boxing

-- Insert initial athlete scores (will be updated through bout history)
INSERT INTO athlete_score (athlete_id, style_id, score)
VALUES 
(1, 2, 400), -- John Smith BJJ
(2, 2, 400), -- John Doe BJJ
(1, 1, 400), -- John Smith Muay Thai
(3, 1, 400), -- Mike Johnson Muay Thai
(4, 3, 400), -- Sarah Williams Boxing
(5, 2, 400), -- David Brown BJJ
(6, 1, 400), -- Emily Davis Muay Thai
(7, 3, 400), -- Robert Wilson Boxing
(8, 2, 400), -- Jessica Taylor BJJ
(9, 1, 400), -- Michael Anderson Muay Thai
(10, 3, 400); -- Lisa Thomas Boxing

-- Insert historical bouts for John Smith (ID: 1) with dates spread from Jan 1 to Apr 10
INSERT INTO bout (challenger_id, acceptor_id, referee_id, style_id, accepted, completed, cancelled, points, created_dt)
VALUES 
-- BJJ matches (showing progression)
(1, 2, 3, 2, true, true, false, 15, '2024-01-05 19:30:00'),  -- vs John Doe
(5, 1, 3, 2, true, true, false, 18, '2024-01-15 20:15:00'),  -- vs David Brown
(1, 8, 3, 2, true, true, false, 20, '2024-01-28 18:45:00'),  -- vs Jessica Taylor
(2, 1, 5, 2, true, true, false, 22, '2024-02-10 19:00:00'),  -- vs John Doe rematch
(1, 5, 8, 2, true, true, false, 25, '2024-02-22 20:30:00'),  -- vs David Brown rematch
(8, 1, 2, 2, true, true, false, 18, '2024-03-05 19:15:00'),  -- vs Jessica Taylor rematch
(1, 2, 5, 2, true, true, false, 20, '2024-03-18 20:00:00'),  -- vs John Doe final

-- Muay Thai matches
(3, 1, 9, 1, true, true, false, 15, '2024-01-08 18:30:00'),  -- vs Mike Johnson
(1, 6, 3, 1, true, true, false, 18, '2024-01-20 19:45:00'),  -- vs Emily Davis
(9, 1, 6, 1, true, true, false, 20, '2024-02-03 20:15:00'),  -- vs Michael Anderson
(1, 3, 9, 1, true, true, false, 22, '2024-02-15 18:30:00'),  -- vs Mike Johnson rematch
(6, 1, 3, 1, true, true, false, 25, '2024-02-28 19:00:00'),  -- vs Emily Davis rematch
(1, 9, 6, 1, true, true, false, 18, '2024-03-12 20:45:00'),  -- vs Michael Anderson rematch
(3, 1, 9, 1, true, true, false, 20, '2024-03-25 19:30:00'),  -- vs Mike Johnson final
(1, 6, 3, 1, true, true, false, 22, '2024-04-05 18:45:00');  -- vs Emily Davis final

-- Insert outcomes for historical bouts with corresponding dates
INSERT INTO outcome (bout_id, winner_id, loser_id, style_id, is_draw, created_dt)
VALUES 
-- BJJ Outcomes (John wins 4, loses 3)
(1, 1, 2, 2, false, '2024-01-05 21:30:00'),    -- Win vs John Doe
(2, 5, 1, 2, false, '2024-01-15 22:00:00'),    -- Loss vs David Brown
(3, 1, 8, 2, false, '2024-01-28 20:45:00'),    -- Win vs Jessica Taylor
(4, 1, 2, 2, false, '2024-02-10 21:15:00'),    -- Win vs John Doe rematch
(5, 5, 1, 2, false, '2024-02-22 22:30:00'),    -- Loss vs David Brown rematch
(6, 8, 1, 2, false, '2024-03-05 21:00:00'),    -- Loss vs Jessica Taylor rematch
(7, 1, 2, 2, false, '2024-03-18 22:15:00'),    -- Win vs John Doe final

-- Muay Thai Outcomes (John wins 5, loses 3)
(8, 3, 1, 1, false, '2024-01-08 20:30:00'),    -- Loss vs Mike Johnson
(9, 1, 6, 1, false, '2024-01-20 21:45:00'),    -- Win vs Emily Davis
(10, 9, 1, 1, false, '2024-02-03 22:15:00'),   -- Loss vs Michael Anderson
(11, 1, 3, 1, false, '2024-02-15 20:30:00'),   -- Win vs Mike Johnson rematch
(12, 1, 6, 1, false, '2024-02-28 21:00:00'),   -- Win vs Emily Davis rematch
(13, 1, 9, 1, false, '2024-03-12 22:45:00'),   -- Win vs Michael Anderson rematch
(14, 3, 1, 1, false, '2024-03-25 21:30:00'),   -- Loss vs Mike Johnson final
(15, 1, 6, 1, false, '2024-04-05 20:45:00');   -- Win vs Emily Davis final

-- Update athlete records for historical bouts
UPDATE athlete_record 
SET wins = wins + 9, losses = losses + 6 
WHERE athlete_id = 1;

-- Update other athletes' records from historical bouts
UPDATE athlete_record SET wins = wins + 2, losses = losses + 2 WHERE athlete_id = 2;  -- John Doe
UPDATE athlete_record SET wins = wins + 2, losses = losses + 1 WHERE athlete_id = 3;  -- Mike Johnson
UPDATE athlete_record SET wins = wins + 2, losses = losses + 1 WHERE athlete_id = 5;  -- David Brown
UPDATE athlete_record SET wins = wins + 0, losses = losses + 3 WHERE athlete_id = 6;  -- Emily Davis
UPDATE athlete_record SET wins = wins + 1, losses = losses + 2 WHERE athlete_id = 8;  -- Jessica Taylor
UPDATE athlete_record SET wins = wins + 1, losses = losses + 1 WHERE athlete_id = 9;  -- Michael Anderson

-- Insert historical score progression with corresponding dates
INSERT INTO athlete_score_history (athlete_id, style_id, outcome_id, previous_score, new_score, created_dt)
VALUES
-- BJJ progression
(1, 2, 1, 400, 415, '2024-01-05 21:30:00'),    -- First win vs John Doe
(1, 2, 2, 415, 405, '2024-01-15 22:00:00'),    -- Loss vs David Brown
(1, 2, 3, 405, 425, '2024-01-28 20:45:00'),    -- Win vs Jessica Taylor
(1, 2, 4, 425, 447, '2024-02-10 21:15:00'),    -- Win vs John Doe rematch
(1, 2, 5, 447, 435, '2024-02-22 22:30:00'),    -- Loss vs David Brown rematch
(1, 2, 6, 435, 425, '2024-03-05 21:00:00'),    -- Loss vs Jessica Taylor rematch
(1, 2, 7, 425, 445, '2024-03-18 22:15:00'),    -- Win vs John Doe final

-- Muay Thai progression
(1, 1, 8, 400, 390, '2024-01-08 20:30:00'),    -- Loss vs Mike Johnson
(1, 1, 9, 390, 408, '2024-01-20 21:45:00'),    -- Win vs Emily Davis
(1, 1, 10, 408, 398, '2024-02-03 22:15:00'),   -- Loss vs Michael Anderson
(1, 1, 11, 398, 420, '2024-02-15 20:30:00'),   -- Win vs Mike Johnson rematch
(1, 1, 12, 420, 445, '2024-02-28 21:00:00'),   -- Win vs Emily Davis rematch
(1, 1, 13, 445, 463, '2024-03-12 22:45:00'),   -- Win vs Michael Anderson rematch
(1, 1, 14, 463, 453, '2024-03-25 21:30:00'),   -- Loss vs Mike Johnson final
(1, 1, 15, 453, 475, '2024-04-05 20:45:00');   -- Win vs Emily Davis final

-- Update scores after historical matches
UPDATE athlete_score SET score = 445 WHERE athlete_id = 1 AND style_id = 2;  -- John Smith BJJ
UPDATE athlete_score SET score = 475 WHERE athlete_id = 1 AND style_id = 1;  -- John Smith Muay Thai
UPDATE athlete_score SET score = 435 WHERE athlete_id = 2 AND style_id = 2;  -- John Doe BJJ
UPDATE athlete_score SET score = 445 WHERE athlete_id = 3 AND style_id = 1;  -- Mike Johnson Muay Thai
UPDATE athlete_score SET score = 450 WHERE athlete_id = 5 AND style_id = 2;  -- David Brown BJJ
UPDATE athlete_score SET score = 380 WHERE athlete_id = 6 AND style_id = 1;  -- Emily Davis Muay Thai
UPDATE athlete_score SET score = 420 WHERE athlete_id = 8 AND style_id = 2;  -- Jessica Taylor BJJ
UPDATE athlete_score SET score = 425 WHERE athlete_id = 9 AND style_id = 1;  -- Michael Anderson Muay Thai

-- Insert pending bouts (John Smith as referee)
INSERT INTO bout (challenger_id, acceptor_id, referee_id, style_id, accepted, completed, cancelled, points, created_dt)
VALUES 
(2, 5, 1, 2, false, false, false, 25, '2024-04-08 19:00:00'),  -- John Doe vs David Brown in BJJ
(3, 6, 1, 1, false, false, false, 20, '2024-04-09 20:00:00'),  -- Mike Johnson vs Emily Davis in Muay Thai
(8, 5, 1, 2, false, false, false, 30, '2024-04-10 18:30:00');  -- Jessica Taylor vs David Brown in BJJ

-- Insert incomplete bouts (John Smith as referee)
INSERT INTO bout (challenger_id, acceptor_id, referee_id, style_id, accepted, completed, cancelled, points, created_dt)
VALUES 
(2, 1, 3, 2, true, false, false, 25, '2024-04-08 19:30:00'),   -- John Doe challenging John Smith in BJJ
(6, 1, 9, 1, true, false, false, 20, '2024-04-09 20:30:00'),   -- Emily Davis challenging John Smith in Muay Thai
(9, 1, 3, 1, false, false, false, 30, '2024-04-10 19:00:00'),  -- Michael Anderson challenging John Smith in Muay Thai
(3, 8, 1, 2, true, false, false, 25, '2024-04-11 19:00:00');   -- Mike Johnson vs Jessica Taylor in BJJ (accepted, awaiting jsmith's decision)

-- Insert completed bouts with recent dates
INSERT INTO bout (challenger_id, acceptor_id, referee_id, style_id, accepted, completed, cancelled, points, created_dt)
VALUES 
(1, 3, 9, 1, true, true, false, 20, '2024-04-06 20:00:00'), -- John Smith vs Mike Johnson in Muay Thai
(8, 5, 10, 2, true, true, false, 15, '2024-04-07 19:15:00'), -- Jessica Taylor vs David Brown in BJJ
(4, 7, 2, 3, true, true, false, 25, '2024-04-08 20:30:00'), -- Sarah Williams vs Robert Wilson in Boxing
(6, 9, 1, 1, true, true, false, 18, '2024-04-09 18:45:00'), -- Emily Davis vs Michael Anderson in Muay Thai
(10, 4, 3, 3, true, true, false, 22, '2024-04-10 19:45:00'); -- Lisa Thomas vs Sarah Williams in Boxing

-- Insert outcomes for completed bouts
INSERT INTO outcome (bout_id, winner_id, loser_id, style_id, is_draw, created_dt)
VALUES 
(16, 1, 3, 1, false, '2024-04-06 22:00:00'), -- John Smith beat Mike Johnson in Muay Thai
(17, 5, 8, 2, false, '2024-04-07 21:15:00'), -- David Brown beat Jessica Taylor in BJJ
(18, 7, 4, 3, false, '2024-04-08 22:30:00'), -- Robert Wilson beat Sarah Williams in Boxing
(19, 9, 6, 1, false, '2024-04-09 20:45:00'), -- Michael Anderson beat Emily Davis in Muay Thai
(20, 4, 10, 3, false, '2024-04-10 21:45:00'); -- Sarah Williams beat Lisa Thomas in Boxing

-- Update bout completion status
UPDATE bout SET completed = true WHERE bout_id IN (16, 17, 18, 19, 20);

-- Update athlete records for winners and losers
UPDATE athlete_record SET wins = wins + 1 WHERE athlete_id = 1;  -- John Smith wins +1
UPDATE athlete_record SET losses = losses + 1 WHERE athlete_id = 3;  -- Mike Johnson losses +1
UPDATE athlete_record SET wins = wins + 1 WHERE athlete_id = 5;  -- David Brown wins +1
UPDATE athlete_record SET losses = losses + 1 WHERE athlete_id = 8;  -- Jessica Taylor losses +1
UPDATE athlete_record SET wins = wins + 1 WHERE athlete_id = 7;  -- Robert Wilson wins +1
UPDATE athlete_record SET losses = losses + 1 WHERE athlete_id = 4;  -- Sarah Williams losses +1 (first match)
UPDATE athlete_record SET wins = wins + 1 WHERE athlete_id = 9;  -- Michael Anderson wins +1
UPDATE athlete_record SET losses = losses + 1 WHERE athlete_id = 6;  -- Emily Davis losses +1
UPDATE athlete_record SET wins = wins + 1 WHERE athlete_id = 4;  -- Sarah Williams wins +1 (second match)
UPDATE athlete_record SET losses = losses + 1 WHERE athlete_id = 10;  -- Lisa Thomas losses +1

-- Update athlete scores for winners (gained points)
UPDATE athlete_score SET score = score + 20 WHERE athlete_id = 1 AND style_id = 1;  -- John Smith + 20 in Muay Thai
UPDATE athlete_score SET score = score + 15 WHERE athlete_id = 5 AND style_id = 2;  -- David Brown + 15 in BJJ
UPDATE athlete_score SET score = score + 25 WHERE athlete_id = 7 AND style_id = 3;  -- Robert Wilson + 25 in Boxing
UPDATE athlete_score SET score = score + 18 WHERE athlete_id = 9 AND style_id = 1;  -- Michael Anderson + 18 in Muay Thai
UPDATE athlete_score SET score = score + 22 WHERE athlete_id = 4 AND style_id = 3;  -- Sarah Williams + 22 in Boxing

-- Update athlete scores for losers (lost points)
UPDATE athlete_score SET score = score - 10 WHERE athlete_id = 3 AND style_id = 1;  -- Mike Johnson - points in Muay Thai
UPDATE athlete_score SET score = score - 8 WHERE athlete_id = 8 AND style_id = 2;   -- Jessica Taylor - points in BJJ
UPDATE athlete_score SET score = score - 12 WHERE athlete_id = 4 AND style_id = 3;  -- Sarah Williams - points in Boxing
UPDATE athlete_score SET score = score - 9 WHERE athlete_id = 6 AND style_id = 1;   -- Emily Davis - points in Muay Thai
UPDATE athlete_score SET score = score - 11 WHERE athlete_id = 10 AND style_id = 3; -- Lisa Thomas - points in Boxing

-- Insert score history records for recent matches
INSERT INTO athlete_score_history (athlete_id, style_id, outcome_id, previous_score, new_score, created_dt)
VALUES
(1, 1, 16, 475, 495, '2024-04-06 22:00:00'),   -- John Smith's latest Muay Thai win
(5, 2, 17, 450, 465, '2024-04-07 21:15:00'),   -- David Brown's latest BJJ win
(7, 3, 18, 400, 425, '2024-04-08 22:30:00'),   -- Robert Wilson's Boxing win
(9, 1, 19, 425, 443, '2024-04-09 20:45:00'),   -- Michael Anderson's Muay Thai win
(4, 3, 20, 388, 410, '2024-04-10 21:45:00');   -- Sarah Williams' Boxing win

COMMIT; 