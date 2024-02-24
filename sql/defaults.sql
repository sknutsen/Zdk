INSERT INTO equipment(name)
VALUES  
    ('Barbell'),
    ('Body'),
    ('Dumbbell'),
    ('Kettlebell'),
    ('Machine'),
    ('Trap bar'),
    ('Treadmill')
ON CONFLICT (name) DO NOTHING;

INSERT INTO exercise_defs(name)
VALUES  
    ('Bench press'),
    ('Burpee'),
    ('Deadlift'),
    ('Hip thrust'),
    ('Jog'),
    ('Lat pull-down'),
    ('Leg curl'),
    ('Leg extension'),
    ('Leg press'),
    ('Lunge'),
    ('Man-maker'),
    ('Mountain climber'),
    ('Plank with shoulder taps'),
    ('Pull-up'),
    ('Push-up'),
    ('Row indoor'),
    ('Run'),
    ('Shoulder press'),
    ('Sit-up'),
    ('Sled push'),
    ('Squat'),
    ('Squat thrust'),
    ('Standing calf raise'),
    ('Walk')
ON CONFLICT (name) DO NOTHING;

INSERT INTO unit_types(name)
VALUES
    ('Calories'),
    ('Reps'),
    ('Meters')
ON CONFLICT (name) DO NOTHING;
