SELECT A.ID, A.UserName, B.UserName AS ParentUserName
FROM USER as A
  LEFT JOIN USER as B
    ON A.Parent = B.ID
ORDER BY ID ASC