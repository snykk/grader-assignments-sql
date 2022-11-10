-- TODO: answer here
UPDATE students
SET address = 'Bandung'
WHERE address is null AND status = 'active';
