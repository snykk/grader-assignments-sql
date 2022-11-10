-- TODO: answer here
SELECT rp.id, st.fullname, st.class, st.status, rp.study, rp.score
FROM students st
JOIN reports rp ON st.id = rp.student_id
WHERE rp.score < 70
ORDER BY rp.score