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

-- Insert bouts
INSERT INTO bout (challenger_id, acceptor_id, referee_id, style_id, accepted, completed, cancelled, points)
VALUES 
(1, 2, 3, 2, true, false, false, 0), -- John Smith vs John Doe in BJJ
(4, 5, 6, 3, true, false, false, 0), -- Sarah Williams vs David Brown in Boxing
(7, 8, 9, 2, true, false, false, 0); -- Robert Wilson vs Jessica Taylor in BJJ

-- Insert athlete scores
INSERT INTO athlete_score (athlete_id, style_id, score)
SELECT athlete_id, style_id, 400
FROM athlete_style;

-- Insert following relationships
INSERT INTO following (follower_id, followed_id)
VALUES 
(3, 1), -- Mike Johnson follows John Smith
(4, 1), -- Sarah Williams follows John Smith
(5, 1), -- David Brown follows John Smith
(1, 3), -- John Smith follows Mike Johnson
(1, 4), -- John Smith follows Sarah Williams
(1, 5), -- John Smith follows David Brown
(6, 1), -- Emily Davis follows John Smith
(7, 1), -- Robert Wilson follows John Smith
(8, 1), -- Jessica Taylor follows John Smith
(1, 6), -- John Smith follows Emily Davis
(1, 7), -- John Smith follows Robert Wilson
(1, 8); -- John Smith follows Jessica Taylor

COMMIT; 