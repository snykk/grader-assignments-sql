mkdir "1. Exercise DDL"
cd "1. Exercise DDL"
grader-cli assignment pull exercise-db-1
grader-cli assignment pull exercise-db-2
grader-cli assignment pull exercise-db-3
cd ../

mkdir "2. Assignment 1"
cd "2. Assignment 1"
grader-cli assignment pull assignment-db-1
cd ../

mkdir "3. Exercise DML"
cd "3. Exercise DML"
grader-cli assignment pull exercise-db-4-insert
grader-cli assignment pull exercise-db-5-update
grader-cli assignment pull exercise-db-6-delete
cd ../

mkdir "4. Assignment 2"
cd "4. Assignment 2"	
grader-cli assignment pull assignment-db-2
cd ../


mkdir "5. Exercise DQL"
cd "5. Exercise DQL"
grader-cli assignment pull exercise-db-7-query
grader-cli assignment pull exercise-db-8-query
cd ../

mkdir "6. Assignment 3"
cd "6. Assignment 3"
grader-cli assignment pull assignment-db-3
cd ../

mkdir "7. Exercise Join"
cd "7. Exercise Join"
grader-cli assignment pull exercise-db-9-join
cd ../

mkdir "8. Assignment 4"
cd "8. Assignment 4"
grader-cli assignment pull assignment-db-4-join
cd ../
