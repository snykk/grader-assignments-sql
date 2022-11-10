-- TODO: answer here
SELECT id, first_name  || ' ' || last_name AS fullname, substring(exam_id from 1 for 2) as class, (bahasa_indonesia + bahasa_inggris + matematika + ipa) / 4 as average_score
FROM final_scores
WHERE exam_status = 'pass' AND fee_status != 'not paid'
ORDER BY average_score DESC
LIMIT 5