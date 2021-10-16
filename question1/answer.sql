SELECT A.ID, A.UserName, B.UserName AS ParentUserName
FROM USER as A
  LEFT JOIN USER as B
    ON a.Parent = B.ID